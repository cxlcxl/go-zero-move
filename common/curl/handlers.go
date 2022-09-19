package curl

import (
	"business/common/vars"
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type HeaderHandlers func(*Curl)

func FormHeader() HeaderHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
}

func JsonHeader() HeaderHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Content-Type", "application/json")
	}
}

func Authorization(token string) HeaderHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Authorization", token)
	}
}

func (c *Curl) JsonData(data interface{}) (*Curl, error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if strings.ToLower(vars.Env) == "dev" {
		fmt.Println()
		fmt.Println("===> 请求报文：", string(bs))
		fmt.Println()
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
