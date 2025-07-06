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

func (w *Weverse) GetAccountCreationTerms() (*AccountTerms, error) {
	req, err := http.NewRequest(http.MethodGet, "https://accountapi.weverse.io/web/api/v3/terms/ACCOUNT_CREATION?language=en", nil)
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get account creation terms: %s", resp.Status)
	}
	termsResult := new(AccountTerms)
	if err := json.NewDecoder(resp.Body).Decode(termsResult); err != nil {
		return nil, fmt.Errorf("failed to decode terms response: %w", err)
	}
	return termsResult, nil
}

func (w *Weverse) GetAccountServiceTerms() (*AccountTerms, error) {
	req, err := http.NewRequest(http.MethodGet, "https://accountapi.weverse.io/web/api/v3/terms/SERVICE_CONNECTION/weverse/?language=en", nil)
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get account service terms: %s", resp.Status)
	}
	termsResult := new(AccountTerms)
	if err := json.NewDecoder(resp.Body).Decode(termsResult); err != nil {
		return nil, fmt.Errorf("failed to decode terms response: %w", err)
	}
	return termsResult, nil
}

func (w *Weverse) CreateAccount() error {
	terms := new(AccountTerms)
	termsResult, err := w.GetAccountCreationTerms()
	if err != nil {
		return fmt.Errorf("failed to get account creation terms: %w", err)
	}
	terms.Terms = termsResult.Terms
	serviceTermsResult, err := w.GetAccountServiceTerms()
	if err != nil {
		return fmt.Errorf("failed to get account service terms: %w", err)
	}
	terms.Terms = append(terms.Terms, serviceTermsResult.Terms...)
	termsAgreements := "["
	for i, term := range terms.Terms {
		if i > 0 {
			termsAgreements += ","
		}
		if term.Required {
			termsAgreements += fmt.Sprintf(`{"termsDocumentId":"%s","agreed":true}`, term.TermsDocumentID)
		} else {
			termsAgreements += fmt.Sprintf(`{"termsDocumentId":"%s","agreed":false}`, term.TermsDocumentID)
		}
	}
	termsAgreements += "]"
	sendBody := fmt.Sprintf(`{"email":"%s","password":"%s","nickname":"%s","termsAgreements":%s}`, w.Email, w.Password, w.Nickname, termsAgreements)
	req, err := http.NewRequest(http.MethodPost, "https://accountapi.weverse.io/web/api/v4/signup/by-credentials", io.NopCloser(strings.NewReader(sendBody)))
	if err != nil {
		return fmt.Errorf("failed to create account request: %w", err)
	}
	for key, value := range AccountDefaultHeaders {
		req.Header.Set(key, value)
	}
	resp, err := w.Client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to create account: %s", resp.Status)
	}
	return nil
}