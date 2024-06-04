package config

import (
	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

var Get = map[string]Data{
	"app":         getApp(),
	"flutterwave": getFlutterwave(),
}
