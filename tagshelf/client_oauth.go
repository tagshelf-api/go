package tagshelf

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/tagshelf-api/go/tagshelf/constant"
)

type OAuthClient struct {
	client
	token *OAuthToken
}

// OAuthToken these are the fields of an OAuth token from tagshelf
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

	c.Header.Set("Content-Type", constant.ContentTypeJSON)
	c.Header.Set("User-Agent", constant.UserAgentHeader)
	c.Header.Set(
		"Authorization",
		fmt.Sprintf(constant.AuthHBearer, token.AccessToken),
	)
	return c, nil
}

func (c OAuthClient) login() (token *OAuthToken, err error) {
	c.Header.Set("Content-Type", constant.ContentTypeXForm)
	data := url.Values{}
	data.Set(constant.OAuthGrantType, c.GrantType)
	data.Set(constant.OAuthUSR, c.User)
	data.Set(constant.OAuthPWD, c.Pass)

	req, err := http.NewRequest(
		constant.EndpointTokenMethod, constant.EndpointToken,
		strings.NewReader(data.Encode()),
	)
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
