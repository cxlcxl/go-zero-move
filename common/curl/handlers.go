package curl

import (
	"bytes"
	"encoding/json"
	"net/url"
	"strings"
)

type QueryHandlers func(*Curl)

func FormHeader() QueryHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
}

func JsonHeader() QueryHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Content-Type", "application/json")
	}
}

func Authorization(token string) QueryHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Authorization", token)
	}
}

func (c *Curl) JsonData(data interface{}) (*Curl, error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	c.data = bytes.NewReader(bs)
	return c, nil
}

func (c *Curl) QueryData(data map[string]string) *Curl {
	c.data = strings.NewReader(HttpBuildQuery(data))
	return c
}

func HttpBuildQuery(params map[string]string) string {
	v := make(url.Values)
	for key, val := range params {
		v.Set(key, val)
	}
	return v.Encode()
}
