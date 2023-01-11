package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

type Config struct {
	Key    string
	Secret string
}

func New(key, secret string) *Config {
	config := &Config{
		Key:    key,
		Secret: secret,
	}

	return config
}

func (p *Config) Signature(message string) string {
	mac := hmac.New(sha256.New, []byte(p.Secret))
	mac.Write([]byte(message))
	return hex.EncodeToString(mac.Sum(nil))
}
