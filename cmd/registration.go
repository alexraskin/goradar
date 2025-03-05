package cmd

import (
	"fmt"

	"github.com/alexraskin/goradar/internal/api"
	"github.com/alexraskin/goradar/internal/display"

	"github.com/spf13/cobra"
)

var registrationCmd = &cobra.Command{
	Use:   "registration [registration]",
	Short: "Search aircraft by registration number",
	Long: `Search for aircraft using their registration number.
Example: goradar registration G-KELS`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.NewClient()
		registration := args[0]

		opts := &api.PaginationOptions{
			Limit:  limit,
			Offset: offset,
		}

		resp, err := client.GetAircraftByRegistration(registration, opts)
		if err != nil {
			return fmt.Errorf("failed to get aircraft data: %w", err)
		}

		display.DisplayAircraft(resp, opts)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(registrationCmd)

	registrationCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Number of results to show per page (0 for all)")
	registrationCmd.Flags().IntVarP(&offset, "offset", "o", 0, "Starting offset for pagination")
}
