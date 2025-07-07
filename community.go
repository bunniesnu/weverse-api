package weverse

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (w *Weverse) SearchCommunity(keyword string, limit, pageNo int) (*PageResult[CommunityRecommend], error) {
	target_path := "/community/v1.0/search"
	queryParams := map[string]string{
		"fields": "communityId,communityName,communityAlias,urlPath,logoImage,homeHeaderImage,memberCount,lastArtistContentPublishedAt,recommended",
		"limit": strconv.Itoa(limit),
		"pageNo": strconv.Itoa(pageNo),
		"pagingType": "PAGE_NO",
		"appId": WeverseWebAppId,
		"keyword": keyword,
		"wpf": "pc",
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
	data := new(PageResult[CommunityRecommend])
	if err := json.NewDecoder(reader).Decode(data); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return data, nil
}

func (w *Weverse) GetCommunityByUrlPath(urlPath string) (int, error) {
	target_path := "/community/v1.0/communityIdUrlPathByUrlPathArtistCode"
	queryParams := map[string]string{
		"appId": WeverseWebAppId,
		"keyword": urlPath,
		"language": "en",
		"os": "WEB",
		"platform": "WEB",
		"wpf": "pc",
	}
	resp, err := w.weverseAPICall(http.MethodGet, target_path, queryParams, nil)
	if err != nil {
		return -2, fmt.Errorf("error making API call: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return -2, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	var reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			return -2, fmt.Errorf("error creating gzip reader: %v", err)
		}
		defer gzipReader.Close()
		reader = gzipReader
	}
	var data struct {
		UrlPath     string `json:"urlPath"`
		CommunityId int    `json:"communityId"`
	}
	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		return -2, fmt.Errorf("error decoding response: %v", err)
	}
	return data.CommunityId, nil
}

func (w *Weverse) GetCommunityById(communityId int) (*Community, error) {
	target_path := fmt.Sprintf("/community/v1.0/community-%d", communityId)
	queryParams := map[string]string{
		"appId": WeverseWebAppId,
		"fieldSet": "communityHomeV1_1",
		"fields": "shopUrl,tabs,membership",
		"wpf": "pc",
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
	data := new(Community)
	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return data, nil
}

func (w *Weverse) GetCommunityNotices(communityId, pageNo, limit int) (*PageResult[CommunityNotice], error) {
	target_path := fmt.Sprintf("/notice/v1.0/community-%d/notices", communityId)
	queryParams := map[string]string{
		"fieldSet": "noticesV1",
		"limit": strconv.Itoa(limit),
		"pageNo": strconv.Itoa(pageNo),
		"pagingType": "PAGE_NO",
		"appId": WeverseWebAppId,
		"wpf": "pc",
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
	data := new(PageResult[CommunityNotice])
	if err := json.NewDecoder(reader).Decode(data); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return data, nil
}