package services

import (
	"fmt"
	"testing"
	"time"
	"github.com/xlzd/gotp"
)



func TestGenerateOTP(t *testing.T) {
	token := NewMFA().GenerateGaToken()
	fmt.Printf("NewMFA().GenerateGaToken(): %s\n", token)

	totp := gotp.NewDefaultTOTP(token)

	code := totp.Now()
	fmt.Printf("totp.Now(): %s\n", code)

	v := totp.Verify(code, time.Now().Unix())
	fmt.Printf("totp.Validate(code): %v\n", v)

}
