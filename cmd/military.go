package cmd

import (
	"fmt"

	"github.com/alexraskin/goradar/internal/api"
	"github.com/alexraskin/goradar/internal/display"

	"github.com/spf13/cobra"
)

var militaryCmd = &cobra.Command{
	Use:   "military",
	Short: "List military aircraft",
	Long: `List all military registered aircraft currently tracked.
Example: goradar military`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.NewClient()

		opts := &api.PaginationOptions{
			Limit:  limit,
			Offset: offset,
		}

		resp, err := client.GetMilitaryAircraft(opts)
		if err != nil {
			return fmt.Errorf("failed to get military aircraft data: %w", err)
		}

		display.DisplayAircraft(resp, opts)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(militaryCmd)

	militaryCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Number of results to show per page (0 for all)")
	militaryCmd.Flags().IntVarP(&offset, "offset", "o", 0, "Starting offset for pagination")
}
