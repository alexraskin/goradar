package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goradar",
	Short: "Aircraft tracking CLI using adsb.lol API",
	Long: `GoRadar is a command-line interface for tracking aircraft using the adsb.lol API.
It provides various commands to search and filter aircraft data.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
