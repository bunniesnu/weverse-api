package weverse

import (
	"net/url"
	"regexp"
	"strings"
	"testing"
	"time"

	"crypto/rand"
	"math/big"

	"github.com/bunniesnu/go-gmailnator"
)

const (
	lower       = "abcdefghijklmnopqrstuvwxyz"
	upper       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits      = "0123456789"
	specials    = "!@#%^_=+"
	allChars    = lower + upper + digits + specials
	passwordLen = 16
)

func getRandomChar(charset string) byte {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	return charset[n.Int64()]
}

func shuffle(bytes []byte) {
	for i := len(bytes) - 1; i > 0; i-- {
		j, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		bytes[i], bytes[j.Int64()] = bytes[j.Int64()], bytes[i]
	}
}

func generatePassword(length int) string {
	if length < 4 {
		panic("Password length must be at least 4 to include all character types.")
	}

	// Ensure at least one character from each required set
	password := []byte{
		getRandomChar(lower),
		getRandomChar(upper),
		getRandomChar(digits),
		getRandomChar(specials),
	}

	// Fill the rest with random characters from all sets
	for i := 4; i < length; i++ {
		password = append(password, getRandomChar(allChars))
	}

	// Shuffle the final password
	shuffle(password)

	return string(password)
}

func TestWeverse(t *testing.T) {
	// Generate a random email using Gmailnator
	gmail, err := gmailnator.NewGmailnator()
	if err != nil {
		t.Errorf("error creating Gmailnator client: %v", err)
	}
	err = gmail.GenerateEmail()
	if err != nil {
		t.Errorf("error generating email: %v", err)
	}
	email := gmail.Email.Email

	// Test nickname suggestion
	password := generatePassword(16)
	w, err := New(email, password, "", 30*time.Second)
	if err != nil {
		t.Errorf("error creating Weverse client: %v", err)
	}
	nickname, err := w.GetAccountNicknameSuggestion()
	if err != nil {
		t.Errorf("error getting nickname suggestion: %v", err)
	} else {
		t.Log("Nickname suggestion success")
	}
	w.Nickname = nickname

	// Test nickname check
	isValid, err := w.CheckNickname(nickname)
	if err != nil {
		t.Errorf("error checking nickname: %v", err)
	}
	if !isValid {
		t.Errorf("nickname %s is not valid", nickname)
	} else {
		t.Log("Nickname check success")
	}

	// Test account creation
	err = w.CreateAccount()
	if err != nil {
		t.Errorf("error creating account: %v", err)
	} else {
		t.Log("Account creation success")
	}

	// Test account email verification
	status, err := w.GetAccountStatus()
	if err != nil {
		t.Errorf("error getting account status: %v", err)
	}
	if !status.EmailVerified {
		t.Log("Email not verified. Proceeding to verification.")
	} else {
		t.Fatal("Email is already verified.")
	}
	found := false
	for range 12 {
		time.Sleep(5 * time.Second)
		mailList, err := gmail.GetMails()
		if err != nil {
			t.Errorf("error getting mails: %v", err)
		}
		for _, mail := range mailList {
			mailBody, err := gmail.GetMailBody(mail.Mid)
			if err != nil {
				t.Errorf("error getting mail body: %v", err)
			}
			re := regexp.MustCompile(`https://account\.weverse\.io/signup[^\s"'<>]+`)
			matches := re.FindAllString(mailBody, -1)
			for _, link := range matches {
				parsed, err := url.QueryUnescape(link)
				if err != nil {
					t.Errorf("error unescaping url: %v", err)
					continue
				}
				finalLink := strings.ReplaceAll(parsed, "&amp;", "&")
				found = true
				clickLink(finalLink)
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}
	if !found {
		t.Fatal("Verification link not found in emails.")
	}
	status, err = w.GetAccountStatus()
	if err != nil {
		t.Errorf("error getting account status: %v", err)
	}
	if !status.EmailVerified {
		t.Fatal("Email verification failed.")
	} else {
		t.Log("Email verification success")
	}
	
	// Test login
	err = w.Login()
	if err != nil {
		t.Errorf("error logging in: %v", err)
	} else {
		t.Log("Login success")
	}

	// Test account info retrieval
	info, err := w.GetAccountInfo()
	if err != nil {
		t.Errorf("error getting account info: %v", err)
	}
	if info.Email != email || info.Nickname != nickname {
		t.Errorf("Expected email %s, got %s", email, info.Email)
	} else {
		t.Log("Account info retrieval success")
	}
}