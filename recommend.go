package weverse

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
)

type HomeResponse struct {
	MainBanners    []Banner     `json:"mainBanners"`
	FeaturedArtist []Community  `json:"featuredArtist"`
	Ads            []Ad         `json:"ads"`
}

func (w *Weverse) Home() (*HomeResponse, error) {
	target_path := "/home/v1.0/home/pc"
	queryParams := map[string]string{
		"wpf":"pc",
	}
	resp, err := w.weverseAPICall(http.MethodGet, target_path, queryParams, nil)
	if err != nil {
		return nil, fmt.Errorf("error making API call: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error creating gzip reader: %v", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}
	homeResponse := new(HomeResponse)
	if err := json.NewDecoder(reader).Decode(homeResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return homeResponse, nil
}

func (w *Weverse) GetDMRecommendations() ([]Community, error) {
	target_path := "/dm/v1.0/dm/recommend-communities"
	queryParams := map[string]string{
		"wpf":"pc",
	}
	resp, err := w.weverseAPICall(http.MethodGet, target_path, queryParams, nil)
	if err != nil {
		return nil, fmt.Errorf("error making API call: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("error creating gzip reader: %v", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}
	var recommendations []Community
	if err := json.NewDecoder(reader).Decode(&recommendations); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return recommendations, nil
}