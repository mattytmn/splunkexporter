package pkg

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

func RunSplunkQuery() {
	fmt.Println(time.February)
}
func GetTokenValue() string {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/splunkextractor")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not get config file: %s", err)
	}

	return viper.Get("api_token").(string)
}
