package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/golang-jwt/jwt/v5"
	"github.com/slack-go/slack"
	"github.com/viadroid/xpanel-go/global"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/utils"
)

type OAuthController struct {
	BaseController
}

func (c *OAuthController) Index() {
	typ := c.Ctx.Input.Param(":type")

	switch typ {
	case "slack":
		c.Slack()
	case "discord":
		c.Discord()
	case "telegram":
		c.Telegram()
	default:
		c.Error(&c.Controller, "xxxx")
	}

	c.Error(&c.Controller, "xxxx")

}

func (c *OAuthController) Slack() {

	conf := models.NewConfig()

	slack_client_id := conf.Obtain("slack_client_id").Value
	slack_team_id := conf.Obtain("slack_team_id").Value
	slack_client_secret := conf.Obtain("slack_client_secret").Value

	baseUrl := global.ConfMap["baseUrl"].(string)
	redirctUrl := fmt.Sprintf("%s/oauth/slack", baseUrl)

	code := c.GetString("code")
	state := c.GetString("state")

	stateKey := fmt.Sprintf("slack_state:%d", c.User.Id)
	if code == "" {
		state = utils.GenRandomString(16)
		err := global.Redis.SetEx(c.Ctx.Request.Context(), stateKey, state, 300*time.Second).Err()
		if err != nil {
			c.Error(&c.Controller, "xxxx")
		}

		redirUrl := fmt.Sprintf(`https://slack.com/openid/connect/authorize?response_type=code&scope=openid profile&client_id=%s&state=%s&team=%s&nonce=%s&redirect_uri=%s`,
			slack_client_id, state, slack_team_id, state, redirctUrl)

		c.JSONResp(map[string]interface{}{
			"ret":   1,
			"redir": redirUrl,
		})
		return
	}

	value, err := global.Redis.Get(c.Ctx.Request.Context(), stateKey).Result()

	if err != nil || state != value {
		c.Error(&c.Controller, "OAuth 请求失败")
		return
	}

	connectResp, err := slack.GetOpenIDConnectToken(http.DefaultClient, slack_client_id, slack_client_secret, code, redirctUrl)
	if err != nil {
		c.Error(&c.Controller, "OAuth 请求失败")
		return
	}
	// fmt.Println(connectResp)

	id_token := connectResp.IdToken
	jwtClaims := jwt.MapClaims{}
	// Parse id_token without key
	token, _, err := new(jwt.Parser).ParseUnverified(id_token, &jwtClaims)
	fmt.Println(token)

	if err != nil {
		c.Error(&c.Controller, "Token parse 请求失败")
		return
	}
	slack_user_id := jwtClaims["https://slack.com/user_id"].(string)

	// check Slack bind
	if c.User.ImType == 1 && c.User.ImValue == slack_user_id || models.NewUser().IsBindIm(1, slack_user_id) {
		c.Error(&c.Controller, "Slack 账户已绑定")
		return
	}

	c.User.ImType = 1
	c.User.ImValue = slack_user_id
	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "Slack 账户绑定失败")
		return
	}

	// redirect to user edit page
	redirctUrl = fmt.Sprintf("%s/user/edit", global.ConfMap["baseUrl"])
	c.Redirect(redirctUrl, 302)
}

func (c *OAuthController) Discord() {
	fmt.Println("Discord")
	conf := models.NewConfig()

	discord_client_id := conf.Obtain("discord_client_id").Value
	discord_client_secret := conf.Obtain("discord_client_secret").Value

	baseUrl := global.ConfMap["baseUrl"].(string)
	redirctUrl := fmt.Sprintf("%s/oauth/discord", baseUrl)

	code := c.GetString("code")
	state := c.GetString("state")
	stateKey := fmt.Sprintf("discord_state:%d", c.User.Id)
	if code == "" {
		state = utils.GenRandomString(16)
		err := global.Redis.SetEx(c.Ctx.Request.Context(), stateKey, state, 300*time.Second).Err()
		if err != nil {
			c.Error(&c.Controller, "xxxx")
		}

		redirUrl := fmt.Sprintf(`https://discord.com/api/oauth2/authorize?client_id=%s&response_type=code&scope=guilds.join identify&state=%s&redirect_uri=%s`,
			discord_client_id, state, redirctUrl)

		c.JSONResp(map[string]interface{}{
			"ret":   1,
			"redir": redirUrl,
		})
		return
	}

	value, err := global.Redis.Get(c.Ctx.Request.Context(), stateKey).Result()

	if err != nil || state != value {
		c.Error(&c.Controller, "OAuth 请求失败")
		return
	}

	discordUrl := "https://discord.com/api/oauth2/token"

	formData := url.Values{}
	formData.Set("client_id", discord_client_id)
	formData.Set("client_secret", discord_client_secret)
	formData.Set("grant_type", "authorization_code")
	formData.Set("code", code)
	formData.Set("redirect_uri", redirctUrl)

	resp, err := http.PostForm(discordUrl, formData)
	if err != nil {
		c.Error(&c.Controller, "OAuth 请求失败")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != 200 {
		c.Error(&c.Controller, "OAuth 请求失败")
		return
	}

	var result map[string]any
	json.Unmarshal(body, &result)
	access_token := result["access_token"]

	discord, _ := discordgo.New(fmt.Sprintf("Bearer %s", access_token))
	user, err := discord.User("@me")
	if err != nil {
		c.Error(&c.Controller, "OAuth 请求失败")
		return
	}

	discord_user_id := user.ID
	// check Discord bind
	if c.User.ImType == 2 && c.User.ImValue == discord_user_id || models.NewUser().IsBindIm(1, discord_user_id) {
		c.Error(&c.Controller, "Discord 账户已绑定")
		return
	}

	c.User.ImType = 2
	c.User.ImValue = discord_user_id
	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "Discord 账户绑定失败")
		return
	}

	discord_bot_token := conf.Obtain("discord_bot_token").Value

	discord_guild_id := conf.Obtain("discord_guild_id")
	if discord_guild_id.ValueToInt() != 0 {
		guild_body := &discordgo.GuildMemberAddParams{
			AccessToken: fmt.Sprintf("Bot %s", access_token),
		}
		discord,_ = discordgo.New(fmt.Sprintf("Bot %s", discord_bot_token))
		err = discord.GuildMemberAdd(discord_guild_id.Value, c.User.ImValue, guild_body)
		if err != nil {
			log.Printf("GuildMemberAdd Err: %v\n", err)
		}

	}

	// redirect to user edit page
	redirctUrl = fmt.Sprintf("%s/user/edit", global.ConfMap["baseUrl"])
	c.Redirect(redirctUrl, 302)

}

type TelegramUser struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	AuthDate  int64  `json:"auth_date"`
	Hash      string `json:"hash"`
}

func (c *OAuthController) Telegram() {
	var tUser TelegramUser
	if err := json.Unmarshal([]byte(c.GetString("user")), &tUser); err != nil {
		c.Error(&c.Controller, "Telegram 账户绑定失败")
		return
	}
	var userMap map[string]any
	json.Unmarshal([]byte(c.GetString("user")), &userMap)

	telegram_token := models.NewConfig().Obtain("telegram_token").Value

	isValid := utils.VaildTelegramAuthorization(telegram_token, userMap)
	if !isValid || time.Now().Unix()-tUser.AuthDate > 86400 {
		c.Error(&c.Controller, "OAuth 请求失败")
		return
	}

	telegram_id := strconv.FormatInt(tUser.Id, 10)
	if c.User.ImType == 2 && c.User.ImValue == telegram_id || models.NewUser().IsBindIm(2, telegram_id) {
		c.Error(&c.Controller, "Telegram 账户已绑定")
		return
	}

	c.User.ImType = 4
	c.User.ImValue = telegram_id
	if _, err := c.User.Update(); err != nil {
		c.Error(&c.Controller, "Telegram 账户绑定失败")
		return
	}
	// redirctUrl := fmt.Sprintf("%s/user/edit", global.ConfMap["baseUrl"])
	// c.Redirect(redirctUrl, 302)

	c.Success(&c.Controller, "Telegram 账户绑定成功")
}
