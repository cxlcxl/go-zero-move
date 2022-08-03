package curl

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

const (
	AllowMethodGet  = "GET"
	AllowMethodPost = "POST"
)

type Curl struct {
	host       string
	request    *http.Request
	isJsonData bool
	data       io.Reader
	method     string
}

func New(host string) *Curl {
	return &Curl{
		host:       host,
		isJsonData: false,
	}
}

func (c *Curl) Request(v interface{}, headers ...QueryHandlers) (err error) {
	if c.method == "" {
		return NotSetMethod
	}
	c.request, err = http.NewRequest(c.method, c.host, c.data)
	if err != nil {
		return err
	}
	for _, header := range headers {
		header(c)
	}
	res, err := (&http.Client{}).Do(c.request)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	var bs []byte
	bs, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bs, v)
	if err != nil {
		return err
	}
	return nil
}

func (c *Curl) Get() *Curl {
	c.method = AllowMethodGet
	return c
}

func (c *Curl) Post() *Curl {
	c.method = AllowMethodPost
	return c
}
