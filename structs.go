package weverse

type GetAccountStatusResult struct {
	EmailVerified bool `json:"emailVerified"`
	HasPassword   bool `json:"hasPassword"`
}