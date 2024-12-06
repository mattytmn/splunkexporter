package internal

import (
	"fmt"
	"os"
	"time"
)

const (
	timeLayout = "January 2006"
	// mm/dd/yyyy:hr:min:sec
	splunkLayout = "01/02/2006:15:04:05"
)

func CreateLogsDir(monthAndYear string) (logsDirPath string) {
	// Check if dir exists
	// either log or month-year directory
	// if not, create dir
	logsDirPath = fmt.Sprintf("logs/%s", monthAndYear)
	err := os.MkdirAll(logsDirPath, 0755)
	if err != nil {
		fmt.Println("Logs directory already exists")
	}
	return logsDirPath
}

func CreateLogFile(filepath string) error {
	fullFileName := fmt.Sprintf("./%s/%s.raw", filepath)
	f, err := os.Create(fullFileName)
	Check(err)
	defer f.Close()

	n, err := f.WriteString("Test write")
	Check(err)
	fmt.Printf("Wrote %d bytes \n", n)
	return nil
}

func GetDaysInMonth(monthAndYear string) (days int, t time.Time) {
	t, _ = time.Parse(timeLayout, monthAndYear)
	fmt.Println(t.Date())
	// lastOfYear := time.Date(2024, 12, 01, 0, 0, 0, 0, time.UTC)

	days = time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	return days, t
}

func QueryDates(t time.Time) (earliest, latest string) {
	earliest = t.Format(splunkLayout)
	latest = t.AddDate(0, 0, 1).Format(splunkLayout)

	return earliest, latest
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
