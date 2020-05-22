package tagshelf

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tagshelf-api/go/constant"

	"github.com/google/uuid"
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

func (c *HMACClient) Sign(method, uri string, body io.Reader) (
	err error,
) {
	uri = strings.ToLower(url.QueryEscape(uri))
	key, err := base64.StdEncoding.DecodeString(c.Config.ApiKey)
	if err != nil {
		return
	}
	appID := c.Config.SecretKey
	timestamp := time.Now().UTC().Unix()
	nonce := uuid.New()
	bodyChecksum := ""
	if body != nil {
		h := md5.New()
		if _, err = io.Copy(h, body); err != nil {
			return
		}
		bodyChecksum = base64.StdEncoding.EncodeToString(h.Sum(nil))
	}

	signatureRaw := fmt.Sprintf(
		"%s%s%s%d%s%s",
		appID, method, uri, timestamp, nonce, bodyChecksum,
	)

	mac := hmac.New(sha256.New, key)
	_, err = mac.Write([]byte(signatureRaw))
	if err != nil {
		return
	}
	bundle := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	signature := fmt.Sprintf("%s:%s:%s:%d", appID, bundle, nonce, timestamp)
	c.Header.Set("Authorization", fmt.Sprintf("amx %s", signature))
	return
}

func (c *HMACClient) Status() (r Responder, err error) {
	method := "GET"
	url := fmt.Sprintf(constant.EpStatus)

	err = c.Sign(method, url, nil)
	if err != nil {
		return
	}

	return c.client.Status()
}

func (c *HMACClient) WhoAmI() (r Responder, err error) {
	method := "GET"
	url := fmt.Sprintf(constant.EpWhoAmI)

	err = c.Sign(method, url, nil)
	if err != nil {
		return
	}

	return c.client.WhoAmI()
}
