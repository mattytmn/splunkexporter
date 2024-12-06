package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
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
	for i := 0; i < days; i++ {
		fmt.Println(i)
		earliest, latest := internal.QueryDates(first.AddDate(0, 0, i))
		sendHttpRequest(earliest, latest, dir)
	}
}

func sendHttpRequest(earliest, latest, dir string) {
	splunkUrl := getConfigValues("splunk_api")
	fmt.Println(splunkUrl)
	token := getConfigValues("api_token")
	search := fmt.Sprintf("search index=* earliest=%s latest=%s", earliest, latest)
	formBody := url.Values{"search": []string{search}, "output_mode": []string{"raw"}}
	data := formBody.Encode()
	req, err := http.NewRequest("POST", splunkUrl, strings.NewReader(data))
	internal.Check(err)
	req.Header.Add("Content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", token)
	client := &http.Client{}
	res, err := client.Do(req)
	defer res.Body.Close()

	filename := strings.ReplaceAll(earliest, "/", "")
	filename = strings.ReplaceAll(filename, ":00:00:00", "")
	fullFilename := fmt.Sprintf("%s/%s.raw", dir, filename)
	out, err := os.Create(fullFilename)
	internal.Check(err)
	defer out.Close()
	io.Copy(out, res.Body)
}

func writeRespToFile(dir string, data []byte) error {
	return nil
}

func setupSplunkExport(month, year string) (filepath string, firstDay time.Time, days int) {
	directoryName := fmt.Sprintf("%s_%s", month, year)
	filepath = internal.CreateLogsDir(directoryName)
	// internal.CreateLogFile(filepath)
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
