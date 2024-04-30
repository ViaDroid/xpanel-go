package utils

import (
	"encoding/json"
	"time"
)

type AuthClaims struct {
	Uid      int    `json:"uid"`
	Email    string `json:"email"`
	Key      string `json:"key"`
	Ip       string `json:"ip"`
	Device   string `json:"device"`
	ExpireIn int64  `json:"expire_in"`
}

func NewAuthClaims(uid int, email string, key string, ip string, device string, expireIn int64) *AuthClaims {
	return &AuthClaims{
		Uid:      uid,
		Email:    email,
		Key:      key,
		Ip:       ip,
		Device:   device,
		ExpireIn: expireIn,
	}
}

// To json
func (c *AuthClaims) ToJSON() string {
	bs, _ := json.Marshal(c)
	return string(bs)
}


func (c *AuthClaims) FromJSON(jsonStr string) error {
	bs := []byte(jsonStr)
	return json.Unmarshal(bs, c)
}

func (c *AuthClaims) IsValid() bool {
	return c.Uid > 0 && c.Email != "" && c.Key != "" && c.Ip != "" && c.Device != "" && c.ExpireIn > 0
}

func (c *AuthClaims) IsExpired() bool {
	return c.ExpireIn < time.Now().UnixMilli()
}
