package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/viadroid/xpanel-go/global"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/sha3"
)

func CookieHash(pass string, expire_in int64) string {
	captchaKey := global.ConfMap["key"].(string)
	return Hash(fmt.Sprintf("%s%s%d", pass, captchaKey, expire_in))
}

func IpHash(ip string, userId int, expire_in int64) string {
	captchaKey := global.ConfMap["key"].(string)
	return Hash(fmt.Sprintf("%s%s%d%d", ip, captchaKey, userId, expire_in))
}

func DeviceHash(ua string, userId int, expire_in int64) string {
	captchaKey := global.ConfMap["key"].(string)
	return Hash(fmt.Sprintf("%s%s%d%d", ua, captchaKey, userId, expire_in))

}

func Hash(content string) string {
	h := make([]byte, 64)
	sha3.ShakeSum256(h, []byte(content))
	str := fmt.Sprintf("%x\n", h)
	return str[5:45]
}

func CheckPassword(hashPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)) == nil
}

func PasswordHash(password string) string {
	bts, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bts)
}

// SHA256 加密算法
func SHA256(content string) string {
	h := sha3.New256()
	h.Write([]byte(content))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// HEX HMAC-SHA-256
func HmacSha256(key, content string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(content))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func VaildTelegramAuthorization(token string, mp map[string]any) bool {
	hash := mp["hash"].(string)
	delete(mp, "hash")
	var data_check_arr []string
	for k, v := range mp {
		s := fmt.Sprintf("%s=%v", k, v)
		if _, ok := v.(float64); ok {
			s = fmt.Sprintf("%s=%.f", k, v)
		}
		data_check_arr = append(data_check_arr, s)
	}
	sort.Strings(data_check_arr)
	data_check_string := strings.Join(data_check_arr, "\n")

	sha256hash := sha256.New()
	io.WriteString(sha256hash, token)
	hmachash := hmac.New(sha256.New, sha256hash.Sum(nil))
	io.WriteString(hmachash, data_check_string)
	return hex.EncodeToString(hmachash.Sum(nil)) == hash
}
