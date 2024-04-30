package services

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/viadroid/xpanel-go/models"
)

type Captcha struct{}

func NewCaptcha() Captcha {
	return Captcha{}
}

func (c Captcha) Generate() map[string]string {
	conf := models.NewConfig()
	captchaProvider := conf.Obtain("captcha_provider")
	switch captchaProvider.Value {
	case "turnstile":
		return map[string]string{"turnstile_sitekey": conf.ObtainValue("turnstile_sitekey")}
	case "geetest":
		return map[string]string{"geetest_id": conf.ObtainValue("geetest_id")}
	case "hcaptcha":
		return map[string]string{"hcaptcha_sitekey": conf.ObtainValue("hcaptcha_sitekey")}
	default:
		return map[string]string{}
	}
}

func (c Captcha) Verify(param map[string]any) bool {

	conf := models.NewConfig()
	switch conf.ObtainValue("captcha_provider") {
	case "turnstile":
		if param["turnstile"] != nil {
			turnstile_secret := conf.ObtainValue("turnstile_secret")
			m := map[string]any{
				"secret":   turnstile_secret,
				"response": param["turnstile"],
			}
			data, _ := json.Marshal(m)

			req, _ := http.NewRequest("POST", "https://challenges.cloudflare.com/turnstile/v0/siteverify", bytes.NewBuffer(data))
			req.Header.Add("Content-Type", "application/json")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return false
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return false
			}

			result := make(map[string]any, 0)
			err = json.Unmarshal(body, &result)
			if err != nil {
				return false
			}

			return result["success"] == true
		}
		return false
	case "geetest":
		if param["geetest"] != nil {
			geetest := param["geetest"]
			captchaId := conf.ObtainValue("geetest_id")
			captchaKey := conf.ObtainValue("geetest_key")

			geetestMap, ok := geetest.(map[string]string)
			if !ok {
				return false
			}

			lot_number := geetestMap["lot_number"]
			captcha_output := geetestMap["captcha_output"]
			pass_token := geetestMap["pass_token"]
			gen_time := geetestMap["gen_time"]

			h := hmac.New(sha256.New, []byte(captchaKey))
			h.Write([]byte(lot_number))
			sign_token := hex.EncodeToString(h.Sum(nil))

			m := map[string]any{
				"lot_number":     lot_number,
				"captcha_output": captcha_output,
				"pass_token":     pass_token,
				"gen_time":       gen_time,
				"sign_token":     sign_token,
			}
			data, _ := json.Marshal(m)

			geetestUrl := fmt.Sprintf("https://gcaptcha4.geetest.com/validate?captcha_id=%s", captchaId)

			req, _ := http.NewRequest("POST", geetestUrl, bytes.NewBuffer(data))
			req.Header.Add("Content-Type", "application/json")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return false
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return false
			}

			result := make(map[string]any, 0)
			err = json.Unmarshal(body, &result)

			// success := result["result"] == "success"

			return err == nil
		}

		return false
	case "hcaptcha":
		if param["hcaptcha"] != nil {
			hcaptchaUrl := "https://hcaptcha.com/siteverify"

			m := map[string]any{
				"secret":   conf.ObtainValue("hcaptcha_secret"),
				"response": param["hcaptcha"],
			}

			data, _ := json.Marshal(m)

			req, _ := http.NewRequest("POST", hcaptchaUrl, bytes.NewBuffer(data))
			req.Header.Add("Content-Type", "application/json")

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				return false
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return false
			}

			result := make(map[string]any, 0)
			err = json.Unmarshal(body, &result)

			// success := result["success"]

			return err == nil

		}

		return false
	default:
		return false
	}
}
