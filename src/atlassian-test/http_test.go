package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server *httptest.Server
	reader io.Reader //Ignore this for now
	url    string
)

func init() {
	http.HandleFunc("/filter", filterHandler)
	server = httptest.NewServer(http.DefaultServeMux)
	url = fmt.Sprintf("%s/filter", server.URL)
}

func TestHandler(t *testing.T) {
	assert := assert.New(t)
	testMessage := `Good morning @chris! (megusta) (coffee)`
	testResponse := `{"emoicons": ["megusta", "coffee"], "mentions": ["chris"]}`

	reader = strings.NewReader(testMessage)

	request, err := http.NewRequest("POST", url, reader)

	res, err := http.DefaultClient.Do(request)
	assert.Nil(err)
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
	buf, err := ioutil.ReadAll(res.Body)
	assert.Nil(err)
	res.Body.Close()
	assert.JSONEq(testResponse, string(buf))
}
