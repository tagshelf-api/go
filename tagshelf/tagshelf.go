package tagshelf

import "fmt"

// New creates new client
func New(config Config) (client Requester, err error) {
	if config.AppApiKey != "" {
		c := &AppApiKeyClient{}
		client, err = c.Auth(config)
	}

	if config.ApiKey != "" && config.SecretKey != "" {
		c := &HMACClient{}
		client, err = c.Auth(config)
	}

	if config.GrantType != "" && config.User != "" && config.Pass != "" {
		c := &OAuthClient{}
		client, err = c.Auth(config)
	}

	if client != nil {
		return
	}

	if err == nil {
		err = fmt.Errorf("No credentials provided")
	}
	return
}
