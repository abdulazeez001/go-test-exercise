package config

import (
	"os"
	"strconv"
)

func getApp() Data {
	var http, _ = strconv.Atoi(os.Getenv("HTTP_PORT"))

	return Data{
		ServiceName: os.Getenv("SERVICE_NAME"),
		HTTPPort:    http,
		BodyLimit:   os.Getenv("BODY_LIMIT"),
		APIVersion:  os.Getenv("API_VERSION"),
		Env:         os.Getenv("APP_ENV"),
	}
}
