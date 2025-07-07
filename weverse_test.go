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

func TestWeverseAccount(t *testing.T) {
	proxyURL := "" // Set your proxy URL if needed, otherwise leave empty
	timeout := 30 * time.Second

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
	w, err := New(email, password, proxyURL, timeout)
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
				err = clickLink(finalLink)
				if err != nil {
					t.Errorf("error clicking verification link: %v", err)
				} else {
					t.Log("Verification link clicked successfully")
				}
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

	// Test saving and loading session
	saveSessionPath := "session.json"
	err = w.SaveSession(saveSessionPath)
	if err != nil {
		t.Errorf("error saving session: %v", err)
	}
	err = w.LoadSession(saveSessionPath, proxyURL, timeout)
	if err != nil {
		t.Errorf("error loading session: %v", err)
	}
	info, err = w.GetAccountInfo()
	if err != nil {
		t.Errorf("error getting account info after loading session: %v", err)
	}
	if info.Email != email || info.Nickname != nickname {
		t.Errorf("Expected email %s, got %s", email, info.Email)
	} else {
		t.Log("Session save/load success")
	}
}

func TestWeverseAPI(t *testing.T) {
	w, err := New("", "", "", 0)
	if err != nil {
		t.Errorf("error creating Weverse client: %v", err)
		return
	}
	err = w.LoadSession("session.json", "", 0)
	if err != nil {
		t.Errorf("error loading session: %v", err)
		return
	}

	// Test Recommendations
	_, err = w.Home()
	if err != nil {
		t.Errorf("error getting home data: %v", err)
		return
	}
	_, err = w.GetDMRecommendations()
	if err != nil {
		t.Errorf("error getting DM recommendations: %v", err)
		return
	}
	t.Log("Recommendations API calls success")
	
	// Table driven tests for community API
	queries := []string {
		"NewJeans",
		"ENHYPEN",
		"TXT",
		"SEVENTEEN",
		"LE SSERAFIM",
		"BABYMONSTER",
		"BTS",
		"Hearts2Hearts",
	}
	for _, query := range queries {
		t.Run("SearchCommunity_"+query, func(t *testing.T) {
			// Test Search
			res, err := w.SearchCommunity(query, 10, 1)
			if err != nil {
				t.Errorf("error searching community: %v", err)
				return
			}
			urlPath := res.Data[0].UrlPath
			t.Log("Search Community success")

			// Test Get Community by URL Path
			communityId, err := w.GetCommunityByUrlPath(urlPath)
			if err != nil {
				t.Errorf("error getting community by URL path: %v", err)
				return
			}
			if communityId != res.Data[0].CommunityID {
				t.Errorf("Expected community ID %d, got %d", res.Data[0].CommunityID, communityId)
				return
			}
			t.Log("Get Community by URL Path success")

			// Test Get Community by ID
			communityInfo, err := w.GetCommunityById(communityId)
			if err != nil {
				t.Errorf("error getting community info: %v", err)
				return
			}
			if communityInfo.CommunityID != communityId {
				t.Errorf("Expected community ID %d, got %d", communityId, communityInfo.CommunityID)
				return
			}
			t.Log("Get Community Info success")

			// Test Get Community User Info
			communityUserInfo, err := w.GetCommunityUserInfo(communityId)
			if err != nil {
				t.Errorf("error getting community user info: %v", err)
				return
			}
			if communityUserInfo.CommunityId != communityId {
				t.Errorf("Expected community ID %d, got %d", communityId, communityUserInfo.CommunityId)
				return
			}
			t.Log("Get Community User Info success")

			// Test Get Community Artists
			artists, err := w.GetCommunityArtists(communityId)
			if err != nil {
				t.Errorf("error getting community artists: %v", err)
				return
			}
			if len(artists) == 0 {
				t.Error("Expected at least one artist, got none")
				return
			}
			t.Log("Get Community Artists success")

			// Test Get Community Notices
			notices, err := w.GetCommunityNotices(communityId, 10, 1)
			if err != nil {
				t.Errorf("error getting community notices: %v", err)
				return
			}
			if len(notices.Data) == 0 {
				t.Error("Expected at least one notice, got none")
				return
			}
			t.Log("Get Community Notices success")
		})
	}
}