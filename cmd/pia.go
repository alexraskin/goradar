package cmd

import (
	"fmt"

	"github.com/alexraskin/goradar/internal/api"
	"github.com/alexraskin/goradar/internal/display"

	"github.com/spf13/cobra"
)

var piaCmd = &cobra.Command{
	Use:   "pia",
	Short: "List aircraft with PIA addresses",
	Long: `List all aircraft with Privacy ICAO Address (PIA).
Example: goradar pia`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.NewClient()

		opts := &api.PaginationOptions{
			Limit:  limit,
			Offset: offset,
		}

		resp, err := client.GetPIAAircraft(opts)
		if err != nil {
			return fmt.Errorf("failed to get PIA aircraft data: %w", err)
		}

		display.DisplayAircraft(resp, opts)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(piaCmd)

	piaCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Number of results to show per page (0 for all)")
	piaCmd.Flags().IntVarP(&offset, "offset", "o", 0, "Starting offset for pagination")
}
