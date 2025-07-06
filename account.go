package weverse

import (
	"fmt"
	"io"
	"net/http"
)

func (w *Weverse) GetAccountStatus() (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://accountapi.weverse.io/web/api/v2/signup/email/status?email=%s", w.Email), nil)
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
	return string(body), nil
}