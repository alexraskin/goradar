package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	baseURL = "https://api.adsb.lol/v2"
)

type PaginationOptions struct {
	Limit  int
	Offset int
}

type Altitude struct {
	Value string
}

func (a *Altitude) UnmarshalJSON(data []byte) error {
	var v any
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	switch val := v.(type) {
	case float64:
		a.Value = fmt.Sprintf("%.0f", val)
	case string:
		a.Value = val
	default:
		return fmt.Errorf("unexpected type for altitude: %T", v)
	}
	return nil
}

type Aircraft struct {
	Hex          string   `json:"hex"`
	Flight       string   `json:"flight,omitempty"`
	Lat          float64  `json:"lat,omitempty"`
	Lon          float64  `json:"lon,omitempty"`
	AltBaro      Altitude `json:"alt_baro,omitempty"`
	Speed        float64  `json:"gs,omitempty"`
	Type         string   `json:"t,omitempty"`
	Category     string   `json:"category,omitempty"`
	Registration string   `json:"r,omitempty"`
	Emergency    string   `json:"emergency,omitempty"`
}

type APIResponse struct {
	Ac    []Aircraft `json:"ac"`
	Msg   string     `json:"msg"`
	Now   int64      `json:"now"`
	Total int        `json:"total"`
}

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (c *Client) GetAircraftByRegistration(registration string, opts *PaginationOptions) (*APIResponse, error) {
	return c.get(fmt.Sprintf("%s/reg/%s", baseURL, registration), opts)
}

func (c *Client) GetAircraftByHex(hex string, opts *PaginationOptions) (*APIResponse, error) {
	return c.get(fmt.Sprintf("%s/hex/%s", baseURL, hex), opts)
}

func (c *Client) GetAircraftByType(aircraftType string, opts *PaginationOptions) (*APIResponse, error) {
	return c.get(fmt.Sprintf("%s/type/%s", baseURL, aircraftType), opts)
}

func (c *Client) GetAircraftBySquawk(squawk string, opts *PaginationOptions) (*APIResponse, error) {
	return c.get(fmt.Sprintf("%s/squawk/%s", baseURL, squawk), opts)
}

func (c *Client) GetMilitaryAircraft(opts *PaginationOptions) (*APIResponse, error) {
	return c.get(fmt.Sprintf("%s/mil", baseURL), opts)
}

func (c *Client) GetPIAAircraft(opts *PaginationOptions) (*APIResponse, error) {
	return c.get(fmt.Sprintf("%s/pia", baseURL), opts)
}

func (c *Client) GetLADDAircraft(opts *PaginationOptions) (*APIResponse, error) {
	return c.get(fmt.Sprintf("%s/ladd", baseURL), opts)
}

func (c *Client) get(url string, opts *PaginationOptions) (*APIResponse, error) {
	if opts != nil {
		if opts.Limit > 0 {
			url = fmt.Sprintf("%s?limit=%d", url, opts.Limit)
		}
		if opts.Offset > 0 {
			separator := "?"
			if opts.Limit > 0 {
				separator = "&"
			}
			url = fmt.Sprintf("%s%soffset=%d", url, separator, opts.Offset)
		}
	}

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned non-200 status code: %d", resp.StatusCode)
	}

	var apiResp APIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &apiResp, nil
}
