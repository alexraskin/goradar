package cmd

import (
	"fmt"

	"github.com/alexraskin/goradar/internal/api"
	"github.com/alexraskin/goradar/internal/display"

	"github.com/spf13/cobra"
)

var laddCmd = &cobra.Command{
	Use:   "ladd",
	Short: "List aircraft on LADD filter",
	Long: `List all aircraft on the Limiting Aircraft Data Displayed (LADD) filter.
Example: goradar ladd`,
	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := api.NewClient()

		opts := &api.PaginationOptions{
			Limit:  limit,
			Offset: offset,
		}

		resp, err := client.GetLADDAircraft(opts)
		if err != nil {
			return fmt.Errorf("failed to get LADD aircraft data: %w", err)
		}

		display.DisplayAircraft(resp, opts)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(laddCmd)

	laddCmd.Flags().IntVarP(&limit, "limit", "l", 0, "Number of results to show per page (0 for all)")
	laddCmd.Flags().IntVarP(&offset, "offset", "o", 0, "Starting offset for pagination")
}
