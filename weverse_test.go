package weverse

import (
	"testing"

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
	gmail, err := gmailnator.NewGmailnator()
	if err != nil {
		t.Errorf("error creating Gmailnator client: %v", err)
	}
	err = gmail.GenerateEmail()
	if err != nil {
		t.Errorf("error generating email: %v", err)
	}
	email := gmail.Email.Email
	password := generatePassword(16)
	w := New(email, password)
	nickname, err := w.GetAccountNicknameSuggestion()
	if err != nil {
		t.Errorf("error getting nickname suggestion: %v", err)
	}
	w.Nickname = nickname
	t.Log("Nickname suggestion success")
	isValid, err := w.CheckNickname(nickname)
	if err != nil {
		t.Errorf("error checking nickname: %v", err)
	}
	if !isValid {
		t.Errorf("nickname %s is not valid", nickname)
	} else {
		t.Log("Nickname check success")
	}
}