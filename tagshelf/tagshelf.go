package tagshelf

import "fmt"

// New creates new client
func New(config Config) (c Requester, err error) {
	if config.AppApiKey != "" {
		client := &AppApiKeyClient{}
		c, err = client.Auth(config)
	}

	if config.ApiKey != "" && config.SecretKey != "" {
		client := &HMACClient{}
		c, err = client.Auth(config)
	}

	if config.GrantType != "" && config.User != "" && config.Pass != "" {
		client := &OAuthClient{}
		c, err = client.Auth(config)
	}

	if c != nil {
		return
	}

	if err == nil {
		err = fmt.Errorf("No credentials provided")
	}
	return
}
