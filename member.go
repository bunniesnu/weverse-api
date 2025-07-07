package weverse

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"net/http"
)

func (w *Weverse) GetCommunityUserInfo(communityId int) (*CommunityUserInfo, error) {
	target_path := fmt.Sprintf("/member/v1.0/community-%d/me", communityId)
	queryParams := map[string]string{
		"fields": "memberId,communityId,joined,joinedDate,profileType,profileName,profileImageUrl,profileCoverImageUrl,profileComment,hidden,blinded,memberJoinStatus,firstJoinAt,followCount,hasMembership,hasOfficialMark,artistOfficialProfile,profileSpaceStatus,availableActions,badges",
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
	data := new(CommunityUserInfo)
	if err := json.NewDecoder(reader).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return data, nil
}