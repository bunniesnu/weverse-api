package weverse

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// GetAccountStatusResult retrieves the account status for the account.
func (w *Weverse) GetAccountStatus() (*GetAccountStatusResult, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://accountapi.weverse.io/web/api/v2/signup/email/status?email=%s", w.Email), nil)
	if err != nil {
		return nil, err
	}
	for key, value := range AccountDefaultHeaders {
		req.Header.Set(key, value)
	}
	resp, err := w.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("failed to get account status: %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := new(GetAccountStatusResult)
	if err := json.Unmarshal(body, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccountNicknameSuggestion retrieves a nickname suggestion for the account.
func (w *Weverse) GetAccountNicknameSuggestion() (string, error) {
	req, err := http.NewRequest(http.MethodGet, "https://accountapi.weverse.io/web/api/v2/signup/nickname/suggestion", nil)
	if err != nil {
		return "", err
	}
	for key, value := range AccountDefaultHeaders {
		req.Header.Set(key, value)
	}
	resp, err := w.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result struct {
		Nickname string `json:"nickname"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	return result.Nickname, nil
}

// CheckNickname checks if the given nickname contains any bad words.
func (w *Weverse) CheckNickname(nickname string) (bool, error) {
	sendBody := fmt.Sprintf(`{"text":"%s"}`, nickname)
	bodyReader := io.NopCloser(strings.NewReader(sendBody))
	req, err := http.NewRequest(http.MethodPost, "https://accountapi.weverse.io/web/api/v2/resources/bad-words", bodyReader)
	if err != nil {
		return false, err
	}
	for key, value := range AccountDefaultHeaders {
		req.Header.Set(key, value)
	}
	resp, err := w.Client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	var result struct {
		ContainsBadWords bool `json:"containsBadWords"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return false, err
	}
	return !result.ContainsBadWords, nil
}