package cmd

import (
	"fmt"

	"github.com/alexraskin/goradar/internal/api"
	"github.com/alexraskin/goradar/internal/display"

	"github.com/spf13/cobra"
)

var squawkCmd = &cobra.Command{
	Use:   "squawk [squawk_code]",
	Short: "Search aircraft by squawk code",
	Long: `Search for aircraft using their squawk code.
Example: goradar squawk 7700`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.NewClient()
		squawk := args[0]

		opts := &api.PaginationOptions{
			Limit:  limit,
			Offset: offset,
		}

		resp, err := client.GetAircraftBySquawk(squawk, opts)
		if err != nil {
			return fmt.Errorf("failed to get aircraft data: %w", err)
		}

		display.DisplayAircraft(resp, opts)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(squawkCmd)

	squawkCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Number of results to show per page (0 for all)")
	squawkCmd.Flags().IntVarP(&offset, "offset", "o", 0, "Starting offset for pagination")
}
