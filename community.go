package weverse

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
)

func (w *Weverse) SearchCommunity(keyword string) ([]Community, error) {
	target_path := "/community/v1.0/search"
	queryParams := map[string]string{
		"fields": "communityId,communityName,communityAlias,urlPath,logoImage,homeHeaderImage,memberCount,lastArtistContentPublishedAt,recommended",
		"appId": WeverseWebAppId,
		"keyword": keyword,
		"wpf":     "pc",
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
	data := new(PageResult[Community])
	if err := json.NewDecoder(reader).Decode(data); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return data.Data, nil
}