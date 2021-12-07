package session

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type session struct {
	url string
	appid string
	secret string
	amadeusOAuth2Token
}

type amadeusOAuth2Token struct {
	Type            string `json:"type"`
	Username        string `json:"username"`
	ApplicationName string `json:"application_name"`
	ClientId        string `json:"client_id"`
	TokenType       string `json:"token_type"`
	AccessToken     string `json:"access_token"`
	ExpiresIn       int    `json:"expires_in"`
	State           string `json:"state"`
	Scope           string `json:"scope"`
}

func newSession(url string, appid string, secret string) *session {
	return &session{url: url, appid: appid, secret: secret}
}

func (s *session) Token() (string, error){
	if len(s.AccessToken) == 0 {
		return s.AccessToken, fmt.Errorf("session: not autorized")
	}
	return s.AccessToken, nil
}

func (s *session) Init() error {

	data := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", s.appid, s.secret)
	rsp, err := http.Post(
		s.url,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))

	if err != nil{
		return fmt.Errorf("session: %v", err)
	}
	defer rsp.Body.Close()

	err = json.NewDecoder(rsp.Body).Decode(&s.amadeusOAuth2Token)
	if err != nil{
		return fmt.Errorf("session: %v", err)
	}
	log.Printf("INFO: %#v", s.amadeusOAuth2Token)

	return nil
}