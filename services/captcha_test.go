package services

import "testing"
import _ "github.com/viadroid/xpanel-go/init"

var captcha Captcha

func init() {
	captcha = NewCaptcha()
}

func TestGenerate(t *testing.T) {
	r := captcha.Generate()
	println(r)
}

func TestVerify(t *testing.T) {
	params := map[string]any{
		"turnstile": "",
	}
	r := captcha.Verify(params)
	println(r)
}
