package command

import (
	"encoding/json"
	"log"
	"os"
	"slices"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/viadroid/xpanel-go/conf"
	"github.com/viadroid/xpanel-go/models"
	"github.com/viadroid/xpanel-go/utils"
)

type Tool struct{}

func SetTelegram() {
	webhookUrl := "baseUrl" + "/callback/telegram?token=" + models.Config{}.Obtain("telegram_request_token").Value
	bot, err := tgbotapi.NewBotAPI(models.Config{}.Obtain("telegram_token").Value)
	if err != nil {
		log.Fatal(err)
	}

	webHookConfig, _ := tgbotapi.NewWebhook(webhookUrl)

	_, err = bot.Request(webHookConfig)
	if err != nil {
		log.Fatalln(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	user, err := bot.GetMe()
	if err != nil {
		log.Println("设置失败！", err)
	} else {
		log.Println("Bot @", user.UserName, "设置成功！")
	}
}

func ResetSetting() {
	orm.NewOrm().Raw("UPDATE config SET value=`default`").Exec()
	log.Println("已使用默认值覆盖所有数据库设置")
}

func ExportSetting() {
	ormer := orm.NewOrm()
	var settings []*models.Config
	ormer.QueryTable(&models.Config{}).All(&settings)

	// var list []any
	// for i, v := range settings {
	// 	m := make(map[string]any)
	// 	// 因为主键自增所以即便设置为 null 也会在导入时自动分配 id
	// 	// 同时避免多位开发者 pull request 时 settings.json 文件 id 重复所可能导致的冲突
	// 	// 设置golang int默认值0
	// 	settings[i].Id = 0
	// 	// 避免开发者调试配置泄露
	// 	settings[i].Value = settings[i].Default
	// 	m["id"] = nil
	// 	m["item"] = v.Item
	// 	m["value"] = v.Value
	// 	m["class"] = v.Class
	// 	m["is_public"] = v.IsPublic
	// 	m["default"] = v.Default
	// 	m["mark"] = v.Mark
	// 	list = append(list, m)
	// }

	jsonBytes, _ := json.MarshalIndent(settings, "", "    ")

	err := os.WriteFile(conf.ConfDir("settings.export.json"), jsonBytes, 0644)
	if err != nil {
		log.Println(err)
	}
}

func ImportSetting() {
	jsonBytes, err := os.ReadFile(conf.ConfDir("settings.json"))
	if err != nil {
		panic("File settings.json not found")
	}
	var configs []models.Config
	err = json.Unmarshal(jsonBytes, &configs)
	if err != nil {
		panic("Json format error")
	}

	addCounter := 0
	updateCounter := 0
	delCounter := 0
	ormer := orm.NewOrm()

	for _, v := range configs {
		config := models.Config{Item: v.Item}
		
		if !config.IsExist(v.Item) {
			ormer.Insert(&v)

			log.Println("添加新数据库设置：", v.Item)
			addCounter++
			continue
		}
		if err := ormer.Read(&config, "item"); err != nil {
			log.Println(err)
		}
		if config.Class != v.Class {
			config.Class = v.Class
			ormer.Update(&config)

			log.Println("更新数据库设置：", v.Item)
			updateCounter++
		}

		// ormer.ReadOrCreate(&v, "item")
	}

	// successNums, err := orm.NewOrm().InsertMulti(len(configs), configs)
	// fmt.Println(successNums, err)

	var settings []*models.Config
	ormer.QueryTable(&models.Config{}).All(&settings)

	for _, v := range settings {
		if !slices.ContainsFunc(configs, func(c models.Config) bool { return c.Item == v.Item }) {
			ormer.Delete(v)
			delCounter++
		}
	}

	if addCounter > 0 {
		log.Println("添加了", addCounter, "项新数据库设置")
	}
	if updateCounter > 0 {
		log.Println("更新了", updateCounter, "项新数据库设置")
	}
	if delCounter > 0 {
		log.Println("移除了", delCounter, "项新数据库设置")
	}

}

func CreateAdmin() {
	y := ""
	email := ""
	passwd := ""

	log.Println(y, email, passwd)

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)

	now := time.Now()
	// do reg user
	user := models.User{}
	user.UserName = "Admin"
	user.Email = email
	user.Remark = ""
	user.Pass = string(hashPassword)
	//  user.Passwd = Tools::genRandomChar(16)
	user.Uuid = uuid.NewString()
	user.ApiToken = uuid.NewString()
	user.Port = utils.GetSsPort()
	user.U = 0
	user.D = 0
	user.TransferEnable = 0
	user.RefBy = 0
	user.IsAdmin = true
	user.RegDate = now.Format(time.DateTime)
	user.Money = 0
	user.ImType = 0
	user.ImValue = ""
	user.Class = 0
	user.NodeIplimit = 0
	user.NodeSpeedlimit = 0
	//  user.Theme = $_ENV['theme']
	//  user.Locale = $_ENV['locale']

	//  user.GaToken = MFA.generateGaToken()
	user.GaEnable = false

}
