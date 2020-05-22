package tagshelf

import (
	"encoding/json"
	"net/http"

	"github.com/tagshelf-api/go/tagshelf/constant"
)

// Config client configuration to comunicate with tagshelf API
// user should just provide one of these options
type Config struct {
	// APP API Key Auth
	AppApiKey string

	// HMAC Auth
	SecretKey string
	ApiKey    string

	// OAuth
	GrantType string
	User      string
	Pass      string
}

type client struct {
	Config
	http.Header
	*http.Client
}

func (c *client) Status() (r Responder, err error) {
	req, err := http.NewRequest(constant.EpStatusM, constant.EpStatus, nil)
	if err != nil {
		return nil, err
	}
	req.Header = c.Header.Clone()

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	j := Response{
		StatusCode: res.StatusCode,
		Payload:    map[string]interface{}{},
	}
	err = json.NewDecoder(res.Body).Decode(&j.Payload)
	return j, err
}

func (c *client) WhoAmI() (r Responder, err error) {
	req, err := http.NewRequest(constant.EpWhoAmIM, constant.EpWhoAmI, nil)
	if err != nil {
		return nil, err
	}
	req.Header = c.Header.Clone()

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	j := Response{
		StatusCode: res.StatusCode,
		Payload:    map[string]interface{}{},
	}
	err = json.NewDecoder(res.Body).Decode(&j.Payload)
	return j, err
}

/*func (c *client) Ping() (r Responder, err error) {
	method := "GET"
	url := fmt.Sprintf(constant.EpPing)
	return
}
func (c *client) FileUpload(File) (r Responder, err error) {
	method := "POST"
	url := fmt.Sprintf(constant.EpFileUpload)
	return
}
func (c *client) FileDetail(id string) (r Responder, err error) {
	method := "GET"
	url := fmt.Sprintf(constant.EpFileDetail, id)
	return
}
func (c *client) JobDetail(id string) (r Responder, err error) {
	method := "GET"
	url := fmt.Sprintf(constant.EpJobDetail, id)
	return
}*/

type Response struct {
	StatusCode int `json:"-"`
	Payload    map[string]interface{}
}

func (c Response) Status() int {
	return c.StatusCode
}

func (c Response) Body() interface{} {
	return c.Payload
}

type File struct{}
