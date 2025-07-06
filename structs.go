package weverse

type GetAccountStatusResult struct {
	EmailVerified bool `json:"emailVerified"`
	HasPassword   bool `json:"hasPassword"`
}

type AccountTerms struct {
	Terms []struct {
		TermsCode        string   `json:"termsCode"`
		TermsDocumentID  string   `json:"termsDocumentId"`
		Title            string   `json:"title"`
		Summary          string   `json:"summary"`
		URL              string   `json:"url"`
		URLType          string   `json:"urlType"`
		Required         bool     `json:"required"`
		Tags             []string `json:"tags"`
	} `json:"terms"`
}