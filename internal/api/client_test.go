package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAircraftByRegistration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/reg/G-KELS" {
			t.Errorf("Expected path /v2/reg/G-KELS, got %v", r.URL.Path)
		}

		// Return a mock response
		response := APIResponse{
			Ac: []Aircraft{
				{
					Hex:          "4008F5",
					Flight:       "VOI1882",
					Registration: "G-KELS",
					Type:         "A320",
					Category:     "L2J",
				},
			},
			Total: 1,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}))
	defer server.Close()

	client := &Client{
		httpClient: &http.Client{},
	}

	resp, err := client.GetAircraftByRegistration("G-KELS", nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(resp.Ac) != 1 {
		t.Errorf("Expected 1 aircraft, got %d", len(resp.Ac))
	}
	if resp.Ac[0].Registration != "G-KELS" {
		t.Errorf("Expected registration G-KELS, got %s", resp.Ac[0].Registration)
	}
}

func TestAltitudeUnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "number altitude",
			input:    []byte(`35000`),
			expected: "35000",
		},
		{
			name:     "string altitude",
			input:    []byte(`"35000"`),
			expected: "35000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var alt Altitude
			err := alt.UnmarshalJSON(tt.input)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if alt.Value != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, alt.Value)
			}
		})
	}
}

func TestPagination(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("limit") != "5" {
			t.Errorf("Expected limit=5, got %v", r.URL.Query().Get("limit"))
		}
		if r.URL.Query().Get("offset") != "10" {
			t.Errorf("Expected offset=10, got %v", r.URL.Query().Get("offset"))
		}

		// Return empty response
		json.NewEncoder(w).Encode(APIResponse{})
	}))
	defer server.Close()

	client := &Client{
		httpClient: &http.Client{},
	}

	opts := &PaginationOptions{
		Limit:  5,
		Offset: 10,
	}

	_, err := client.GetAircraftByRegistration("G-KELS", opts)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
