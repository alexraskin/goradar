package cmd

import (
	"fmt"

	"github.com/alexraskin/goradar/internal/api"
	"github.com/alexraskin/goradar/internal/display"

	"github.com/spf13/cobra"
)

var typeCmd = &cobra.Command{
	Use:   "type [aircraft_type]",
	Short: "Search aircraft by type",
	Long: `Search for aircraft using their type designator.
Example: goradar type A320`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.NewClient()
		aircraftType := args[0]

		opts := &api.PaginationOptions{
			Limit:  limit,
			Offset: offset,
		}

		resp, err := client.GetAircraftByType(aircraftType, opts)
		if err != nil {
			return fmt.Errorf("failed to get aircraft data: %w", err)
		}

		display.DisplayAircraft(resp, opts)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(typeCmd)

	typeCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Number of results to show per page (0 for all)")
	typeCmd.Flags().IntVarP(&offset, "offset", "o", 0, "Starting offset for pagination")
}
