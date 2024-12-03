package internal

import (
	"fmt"
	"os"
	"time"
)

const (
	timeLayout   = "January 2006"
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
	fullFileName := fmt.Sprintf("./%s/01.raw", filepath)
	f, err := os.Create(fullFileName)
	check(err)
	defer f.Close()

	n, err := f.WriteString("Test write")
	check(err)
	fmt.Printf("Wrote %d bytes \n", n)
	return nil
}

func GetDaysInMonth(monthAndYear string) int {
	t, _ := time.Parse(timeLayout, monthAndYear)
	fmt.Println(t.Date())
	lastOfYear := time.Date(2024, 12, 01, 0, 0, 0, 0, time.UTC)
	fmt.Println(lastOfYear.Format(splunkLayout))

	fmt.Println(lastOfYear.AddDate(0, 1, 0))
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
