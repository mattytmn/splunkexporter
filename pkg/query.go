package pkg

import (
	"fmt"
	"log"

	"github.com/mattytmn/splunkextractor/internal"
	"github.com/spf13/viper"
)

var token string = GetTokenValue()

func RunSplunkQuery(month, year string) {
	directoryName := fmt.Sprintf("%s_%s", month, year)
	filepath := internal.CreateLogsDir(directoryName)
	internal.CreateLogFile(filepath)
	monthAndYear := fmt.Sprintf("%s %s", month, year)
	days, first, last := internal.GetDaysInMonth(monthAndYear)

	for i := 1; i <= days; i++ {
		fmt.Println(i)
	}
	fmt.Println(first)
	fmt.Println(last)
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
