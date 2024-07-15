package geoip

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"
)

type Location struct {
	City      string  `json:"city,omitempty"`
	Country   string  `json:"country,omitempty"`
	IP        string  `json:"ip,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
}

type Resolver struct {
	*http.Client
	BaseURL string
}

func NewResolver(baseURL string, timeout time.Duration) *Resolver {
	return &Resolver{
		Client:  &http.Client{Timeout: timeout},
		BaseURL: baseURL,
	}
}

func NewDefaultResolver() *Resolver {
	return NewResolver("http://ip-api.com/json", 15*time.Second)
}

func (r *Resolver) Resolve(ip string) (*Location, error) {
	urlPath, err := url.JoinPath(r.BaseURL, ip)
	if err != nil {
		return nil, err
	}

	resp, err := r.Get(urlPath)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to retrieve data: " + resp.Status)
	}

	var result struct {
		City      string  `json:"city"`
		Country   string  `json:"country"`
		IP        string  `json:"query"`
		Latitude  float64 `json:"lat"`
		Longitude float64 `json:"lon"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &Location{
		City:      result.City,
		Country:   result.Country,
		IP:        result.IP,
		Latitude:  result.Latitude,
		Longitude: result.Longitude,
	}, nil
}
