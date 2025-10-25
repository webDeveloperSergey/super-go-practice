package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)


type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city == "" {
		return &GeoData{
			City: city,
		}, nil
	}

	resp, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("NOT200")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var geo GeoData
	json.Unmarshal(body, &geo)

	return &geo, nil
}