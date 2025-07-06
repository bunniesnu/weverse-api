package weverse

import "net/http"

type Weverse struct {
	Client *http.Client
	Email string
	Password string
	Username string
	AccessToken string
}

func New(email, password string) *Weverse {
	return &Weverse{
		Client: &http.Client{},
		Email: email,
		Password: password,
		Username: "",
		AccessToken: "",
	}
}