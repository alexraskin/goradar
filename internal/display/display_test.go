package display

import (
	"testing"

	"github.com/alexraskin/goradar/internal/api"
)

func TestFormatAircraftType(t *testing.T) {
	tests := []struct {
		name     string
		aircraft api.Aircraft
		expected string
	}{
		{
			name: "with type and category",
			aircraft: api.Aircraft{
				Type:     "A320",
				Category: "L2J",
			},
			expected: "A320 (L2J)",
		},
		{
			name: "with type only",
			aircraft: api.Aircraft{
				Type: "A320",
			},
			expected: "A320",
		},
		{
			name:     "empty",
			aircraft: api.Aircraft{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatAircraftType(tt.aircraft)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestFormatPosition(t *testing.T) {
	tests := []struct {
		name     string
		aircraft api.Aircraft
		expected string
	}{
		{
			name: "with position",
			aircraft: api.Aircraft{
				Lat: 51.5074,
				Lon: -0.1278,
			},
			expected: "51.507400, -0.127800",
		},
		{
			name:     "empty",
			aircraft: api.Aircraft{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatPosition(tt.aircraft)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestFormatAltitude(t *testing.T) {
	tests := []struct {
		name     string
		aircraft api.Aircraft
		expected string
	}{
		{
			name: "with altitude",
			aircraft: api.Aircraft{
				AltBaro: api.Altitude{Value: "35000"},
			},
			expected: "35000 ft",
		},
		{
			name:     "empty",
			aircraft: api.Aircraft{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatAltitude(tt.aircraft)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestFormatSpeed(t *testing.T) {
	tests := []struct {
		name     string
		aircraft api.Aircraft
		expected string
	}{
		{
			name: "with speed",
			aircraft: api.Aircraft{
				Speed: 450.0,
			},
			expected: "450.0 kts",
		},
		{
			name:     "empty",
			aircraft: api.Aircraft{},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatSpeed(tt.aircraft)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestGetFlightRadar24URL(t *testing.T) {
	tests := []struct {
		name     string
		flight   string
		expected string
	}{
		{
			name:     "with flight number",
			flight:   "VOI1882",
			expected: "https://www.flightradar24.com/VOI1882",
		},
		{
			name:     "empty",
			flight:   "",
			expected: "https://www.flightradar24.com/",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getFlightRadar24URL(tt.flight)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
