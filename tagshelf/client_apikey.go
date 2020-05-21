package tagshelf

import "net/http"

type AppApiKeyClient struct {
	client
}

func (c *AppApiKeyClient) Auth(config Config) (r Requester, err error) {
	c.Config = config
	c.Client = &http.Client{}
	c.Header = http.Header{}
	c.Header.Set("Content-Type", "application/json")
	c.Header.Set("X-TagshelfAPI-Key", c.Config.AppApiKey)
	return c, nil
}
