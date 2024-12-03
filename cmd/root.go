package cmd

import (
	"fmt"
	"time"

	"github.com/mattytmn/splunkextractor/pkg"
	"github.com/spf13/cobra"
)

var (
	SplunkMonth string
	SplunkYear  string
)

var rootCmd = &cobra.Command{
	Use:   "splunky [-month]",
	Short: "Get splunk logs for a given month",
	Long:  "A fast way to extract all information from splunk for a given query and time period",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Getting splunk logs")
		pkg.RunSplunkQuery(SplunkMonth, SplunkYear)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&SplunkYear, "year", "y", string(time.Now().Year()), "Year to run query on")
	rootCmd.PersistentFlags().StringVarP(&SplunkMonth, "month", "m", "January", "Month to run query on")
}
