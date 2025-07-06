package weverse

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

type Weverse struct {
	Client      *http.Client `json:"-"`
	Email       string       `json:"email"`
	Password    string       `json:"password"`
	Nickname    string       `json:"nickname"`
	AccessToken string       `json:"access_token"`
}

func New(email, password, proxyURL string, timeout time.Duration) (*Weverse, error) {
	client, err := MakeProxyClient(proxyURL, timeout)
	if err != nil {
		return nil, err
	}
	return &Weverse{
		Client: client,
		Email: email,
		Password: password,
		Nickname: "",
		AccessToken: "",
	}, nil
}

func (w *Weverse) SaveSession(destPath string) error {
	file, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer file.Close()
	session := struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		Nickname    string `json:"nickname"`
		AccessToken string `json:"access_token"`
	}{
		Email:       w.Email,
		Password:    w.Password,
		Nickname:    w.Nickname,
		AccessToken: w.AccessToken,
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(session)
}

func (w *Weverse) LoadSession(srcPath, proxyURL string, timeout time.Duration) error {
	file, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer file.Close()
	session := struct {
		Email       string `json:"email"`
		Password    string `json:"password"`
		Nickname    string `json:"nickname"`
		AccessToken string `json:"access_token"`
	}{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&session); err != nil {
		return err
	}
	w.Client, err = MakeProxyClient(proxyURL, timeout)
	if err != nil {
		return err
	}
	w.Email = session.Email
	w.Password = session.Password
	w.Nickname = session.Nickname
	w.AccessToken = session.AccessToken
	return nil
}