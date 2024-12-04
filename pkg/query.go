package pkg

import (
	"fmt"
	"log"
	"time"

	"github.com/mattytmn/splunkextractor/internal"
	"github.com/spf13/viper"
)

var url, token string = getConfigValues()

func RunSplunkQuery(month, year string) {
	// To iterate through all days use addDate to increment base date and add one 1 base date to get end of range
	// Save log export to file
	// splunkUrl, apiToken := getConfigValues()
	dir, first, days := setupSplunkExport(month, year)
	fmt.Println(days)
	fmt.Println(dir)

	fmt.Println(first)
	for i := 0; i < days; i++ {
		fmt.Println(i)
		earliest, latest := internal.QueryDates(first.AddDate(0, 0, i))

		fmt.Printf("%s to %s\n", earliest, latest)
	}
}

func setupSplunkExport(month, year string) (filepath string, firstDay time.Time, days int) {
	directoryName := fmt.Sprintf("%s_%s", month, year)
	filepath = internal.CreateLogsDir(directoryName)
	internal.CreateLogFile(filepath)
	monthAndYear := fmt.Sprintf("%s %s", month, year)
	days, firstDay = internal.GetDaysInMonth(monthAndYear)

	return filepath, firstDay, days
}

func getConfigValues() (url, token string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/splunk")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not get config file: %s", err)
	}
	token = viper.GetString("api_token")
	url = viper.GetString("splunk_url")
	return url, token
}
