package weverse

import "net/http"

type Weverse struct {
	Client *http.Client
	Email string
	Password string
	Nickname string
	AccessToken string
}

func New(email, password string) *Weverse {
	return &Weverse{
		Client: &http.Client{},
		Email: email,
		Password: password,
		Nickname: "",
		AccessToken: "",
	}
}