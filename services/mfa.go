package services

import (
	"fmt"
	"time"

	"github.com/viadroid/xpanel-go/global"
	"github.com/xlzd/gotp"
)

type MFA struct{}

func NewMFA() *MFA {
	return &MFA{}
}

func (m *MFA) GenerateGaToken() string {
	return gotp.RandomSecret(10)
}

func (m *MFA) ValidateGaToken(token string, code string) bool {
	totp := gotp.NewDefaultTOTP(token)
	return totp.Verify(code, time.Now().Unix())
}

func (m *MFA) GetGaUrl(token, email string) string {
	totp := gotp.NewDefaultTOTP(token)

	appName := global.ConfMap["appName"].(string)
	return totp.ProvisioningUri(fmt.Sprintf("%s:%s", appName, email), appName)
}
