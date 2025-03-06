package display

import (
	"fmt"
	"os"
	"time"

	"github.com/alexraskin/goradar/internal/api"
	"github.com/olekukonko/tablewriter"
)

func getFlightRadar24URL(flight string) string {
	return fmt.Sprintf("https://www.flightradar24.com/%s", flight)
}

func DisplayAircraft(ac *api.APIResponse, opts *api.PaginationOptions) {
	if ac == nil || len(ac.Ac) == 0 {
		fmt.Println("No aircraft found")
		return
	}

	var filteredAc []api.Aircraft
	for _, aircraft := range ac.Ac {
		if aircraft.Flight != "" {
			filteredAc = append(filteredAc, aircraft)
		}
	}

	if len(filteredAc) == 0 {
		fmt.Println("No aircraft with flight numbers found")
		return
	}

	fmt.Printf("Found %d aircraft with flight numbers", len(filteredAc))
	if opts != nil {
		if opts.Limit > 0 {
			fmt.Printf(" (showing %d per page", opts.Limit)
			if opts.Offset > 0 {
				fmt.Printf(", starting at offset %d", opts.Offset)
			}
			fmt.Print(")")
		}
	}
	fmt.Println()
	fmt.Println()

	start := 0
	end := len(filteredAc)
	if opts != nil {
		if opts.Limit > 0 {
			start = opts.Offset
			end = min(start+opts.Limit, len(filteredAc))
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Flight", "Registration", "Type", "Category", "Emergency", "Position", "Altitude", "Speed", "FR24"})
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetCenterSeparator("+")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.SetHeaderLine(true)
	table.SetBorder(true)
	table.SetTablePadding(" ")
	table.SetNoWhiteSpace(false)

	for _, aircraft := range filteredAc[start:end] {
		fr24URL := ""
		if aircraft.Flight != "" {
			fr24URL = getFlightRadar24URL(aircraft.Flight)
		}

		row := []string{
			aircraft.Flight,
			aircraft.Registration,
			aircraft.Type,
			aircraft.Category,
			formatEmergency(aircraft.Emergency),
			formatPosition(aircraft),
			formatAltitude(aircraft),
			formatSpeed(aircraft),
			fr24URL,
		}
		table.Append(row)
	}

	table.Render()
	fmt.Printf("\nLast updated: %s\n", time.Unix(ac.Now, 0).Format(time.RFC3339))
}

func formatAircraftType(ac api.Aircraft) string {
	if ac.Type == "" {
		return ""
	}
	if ac.Category != "" {
		return fmt.Sprintf("%s (%s)", ac.Type, ac.Category)
	}
	return ac.Type
}

func formatPosition(ac api.Aircraft) string {
	if ac.Lat == 0 && ac.Lon == 0 {
		return ""
	}
	return fmt.Sprintf("%.6f, %.6f", ac.Lat, ac.Lon)
}

func formatAltitude(ac api.Aircraft) string {
	if ac.AltBaro.Value == "" {
		return ""
	}
	return fmt.Sprintf("%s ft", ac.AltBaro.Value)
}

func formatSpeed(ac api.Aircraft) string {
	if ac.Speed == 0 {
		return ""
	}
	return fmt.Sprintf("%.1f kts", ac.Speed)
}

func formatEmergency(emergency string) string {
	if emergency == "" {
		return ""
	}
	return emergency
}
