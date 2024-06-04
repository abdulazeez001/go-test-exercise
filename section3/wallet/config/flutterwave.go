package config

import (
	"os"
)

func getFlutterwave() Data {

	return Data{
		FlutterwaveApikey: os.Getenv("FLUTTERWAVE_API_KEY"),
	}
}
