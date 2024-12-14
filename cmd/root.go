package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/mattytmn/blobster"
	"github.com/mattytmn/splunkextractor/pkg"
	"github.com/spf13/cobra"
)

var (
	SplunkMonth    string
	SplunkYear     string
	StorageAccount string
	Container      string
	Directory      string

	ExportCmd = &cobra.Command{
		Use:   "export",
		Short: "Export folder or file to Azure",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Exporting file(s) to Azure...")
			err := blobster.AzureUpload(StorageAccount, Container, Directory)
			if err != nil {
				log.Fatalf("upload error: %v \n", err)
			}
		},
	}
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

	ExportCmd.Flags().StringVarP(&StorageAccount, "storage-account", "s", "stsplunkexport01", "Targeted storage account")
	ExportCmd.Flags().StringVarP(&Container, "container", "c", "test", "Container to target in storage account")
	ExportCmd.Flags().StringVarP(&Directory, "directory", "d", "./", "Directory with files to upload")
	rootCmd.AddCommand(ExportCmd)
}
