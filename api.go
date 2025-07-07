package weverse

import (
	"fmt"
	"net/http"
)

func (w *Weverse) weverseAPICall(method, target_path string, queryParams map[string]string, headers map[string]string) (*http.Response, error) {
	url, err := generateWeverseURL(target_path, queryParams)
	if err != nil {
		return nil, fmt.Errorf("error generating HMAC: %v", err)
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+w.AccessToken)
	for key, value := range WeverseDefaultHeaders {
		req.Header.Set(key, value)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := w.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	return resp, nil
}