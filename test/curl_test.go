package test

import (
	"business/common/curl"
	"testing"
)

func TestCurl(t *testing.T) {
	var aa interface{}
	req, err := curl.New("http://192.168.120.62:30010/login").Post().JsonData(map[string]string{
		"email": "",
		"pass":  "",
	})
	if err != nil {
		t.Fatal(err)
	}
	err = req.Request(&aa)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(aa)
}
