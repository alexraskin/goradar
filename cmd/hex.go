package cmd

import (
	"fmt"

	"github.com/alexraskin/goradar/internal/api"
	"github.com/alexraskin/goradar/internal/display"

	"github.com/spf13/cobra"
)

var hexCmd = &cobra.Command{
	Use:   "hex [icao_hex]",
	Short: "Search aircraft by ICAO hex code",
	Long: `Search for aircraft using their ICAO hex code.
Example: goradar hex 4CA87C`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.NewClient()
		hex := args[0]

		opts := &api.PaginationOptions{
			Limit:  limit,
			Offset: offset,
		}

		resp, err := client.GetAircraftByHex(hex, opts)
		if err != nil {
			return fmt.Errorf("failed to get aircraft data: %w", err)
		}

		display.DisplayAircraft(resp, opts)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(hexCmd)

	hexCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Number of results to show per page (0 for all)")
	hexCmd.Flags().IntVarP(&offset, "offset", "o", 0, "Starting offset for pagination")
}
