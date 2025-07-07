package weverse

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
)

func weverseAPICall(client *http.Client, method, url string, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	for key, value := range WeverseDefaultHeaders {
		req.Header.Set(key, value)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	return resp, nil
}

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
	url, err := generateWeverseURL(target_path, queryParams)
	if err != nil {
		return nil, fmt.Errorf("error generating HMAC: %v", err)
	}
	resp, err := weverseAPICall(w.Client, http.MethodGet, url, nil)
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