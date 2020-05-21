package tagshelf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type OAuthClient struct {
	client
	token *OAuthToken
}

type OAuthToken struct {
	AccessToken string `json:"access_token"`
	Type        string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Username    string `json:"userName"`
	Issued      string `json:".issued"`
	Expires     string `json:".expires"`
}

type LoginResponse struct {
	Response
	Payload OAuthToken
}

func (c *OAuthClient) Auth(config Config) (r Requester, err error) {
	c.Config = config
	c.Client = &http.Client{}
	c.Header = http.Header{}

	token, err := c.login()
	if err != nil {
		return nil, err
	}
	c.token = token

	c.Header.Set("Content-Type", "application/json")
	c.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	return c, nil
}

func (c OAuthClient) login() (token *OAuthToken, err error) {
	c.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	method := "POST"
	URL := fmt.Sprintf(
		"%s/token", "https://staging.tagshelf.com",
	)
	data := url.Values{}
	data.Set("grant_type", c.GrantType)
	data.Set("username", c.User)
	data.Set("password", c.Pass)

	req, err := http.NewRequest(method, URL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header = c.Header.Clone()

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	j := LoginResponse{}
	err = json.NewDecoder(res.Body).Decode(&j.Payload)
	if err != nil {
		return nil, err
	}

	j.StatusCode = res.StatusCode
	if j.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("error getting OAuth token")
	}

	return &j.Payload, nil
}
