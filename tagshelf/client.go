package tagshelf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

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

func (c *client) do(method, ep string, body io.Reader) (r Responder, err error) {
	req, err := http.NewRequest(method, ep, body)
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
		Payload:    &Payload{},
	}
	err = json.NewDecoder(res.Body).Decode(&j.Payload)
	return j, err
}

func (c *client) Status() (r Responder, err error) {
	return c.do(constant.EpStatusM, constant.EpStatus, nil)
}

func (c *client) WhoAmI() (r Responder, err error) {
	return c.do(constant.EpWhoAmIM, constant.EpWhoAmI, nil)
}

func (c *client) Ping() (r Responder, err error) {
	return c.do(constant.EpPingM, constant.EpPing, nil)
}

func (c *client) FileUpload(files *File) (r Responder, err error) {
	return c.do(
		constant.EpFileUploadM,
		constant.EpFileUpload,
		files.NewReader(),
	)
}

func (c *client) FileDetail(id string) (r Responder, err error) {
	return c.do(
		constant.EpFileDetailM,
		fmt.Sprintf(constant.EpFileDetail, id),
		nil,
	)
}
func (c *client) JobDetail(id string) (r Responder, err error) {
	return c.do(
		constant.EpJobDetailM,
		fmt.Sprintf(constant.EpJobDetail, id),
		nil,
	)
}

type Payload map[string]interface{}

type Response struct {
	StatusCode int `json:"-"`
	*Payload
}

func (c Response) Status() int {
	return c.StatusCode
}

func (c Response) Body() Payload {
	return *c.Payload
}

func (c Response) String() string {
	if b, err := json.Marshal(c.Payload); err == nil {
		return fmt.Sprintf(
			"Status(%v) Payload(%v)", c.StatusCode, string(b),
		)
	}

	return fmt.Sprintf("Status(%v) Payload(%v)", c.StatusCode, c.Payload)
}

type File struct {
	URL   string   `json:"url,omitempty"`
	URLs  []string `json:"urls,omitempty"`
	Merge bool     `json:"merge,omitempty"`

	*SingleFile `json:",omitempty"`
	*MultiFile  `json:",omitempty"`
}

type SingleFile struct {
	MetaData FileMetadata `json:"metadata,omitempty"`
}

type MultiFile struct {
	MetaData []FileMetadata `json:"metadata,omitempty"`
}

type FileMetadata map[string]interface{}

func (f *File) Add(url string, urls ...string) (err error) {
	if len(urls) == 0 {
		f.URL = url
		return
	}

	f.URLs = append(f.URLs, url)
	f.URLs = append(f.URLs, urls...)

	return
}

func (f *File) AddMeta(meta FileMetadata, metas ...FileMetadata) (err error) {
	switch {
	// When the merge field is set to true.
	// - The metadata field should be a JSON object
	case f.Merge == true && len(metas) == 0:
		f.SingleFile = &SingleFile{meta}
		f.MultiFile = nil
	// When merge field is set to false
	// - When using the url field  the metadata field should be a JSON object
	case f.Merge == false && len(f.URLs) == 0 && len(metas) == 0:
		f.SingleFile = &SingleFile{meta}
		f.MultiFile = nil
	// - When using the urls field this should be a JSON object array that
	//   matches the urls field array length
	case f.Merge == false && len(f.URLs) == len(metas)+1:
		f.SingleFile = nil
		f.MultiFile = &MultiFile{
			make([]FileMetadata, len(metas)+1),
		}

		f.MultiFile.MetaData = append(f.MultiFile.MetaData, meta)
		f.MultiFile.MetaData = append(f.MultiFile.MetaData, metas...)
	default:
		err = fmt.Errorf("metadata provided does not match merge value")
	}

	return
}

func (f *File) NewReader() io.Reader {
	if b, err := json.Marshal(f); err == nil {
		return bytes.NewReader(b)
	}

	return strings.NewReader(fmt.Sprintf("%v", *f))
}
