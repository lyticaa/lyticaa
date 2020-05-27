package types

import "os"

type Config struct {
	IntercomId   string
	SupportEmail string
}

func NewConfig() *Config {
	return &Config{
		IntercomId:   os.Getenv("INTERCOM_ID"),
		SupportEmail: os.Getenv("SUPPORT_EMAIL"),
	}
}
