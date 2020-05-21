package tagshelf

import (
	"fmt"
	"net/http"
)

type HMACClient struct {
	client
}

func (c *HMACClient) Auth(config Config) (r Requester, err error) {
	c.Config = config
	c.Client = &http.Client{}
	c.Header = http.Header{}

	c.Header.Set("Content-Type", "application/json")
	return c, nil
}

func (c *HMACClient) sign() (err error) {
	signature := ""
	c.Header.Set("Authorization", fmt.Sprintf("amx %s", signature))
	return
}

func (c *HMACClient) WhoAmI() (r Responder, err error) {
	err = c.sign()
	if err != nil {
		return
	}

	return c.client.WhoAmI()
}
