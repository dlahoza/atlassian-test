package links

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func httpGetMock(url string) (result string) {
	result = `<html>
	<head>
	<lalala>Mock for web page</lalala>
	<title>` + url + `</title>
	</head>
	</html>`
	return
}

func TestFilter(t *testing.T) {
	f := &filter{get: httpGetMock}
	assert := assert.New(t)
	cases := map[string][]linksResult{
		`Multiline string with 2 links
		@chris http://very-very-very-very-very-very-very-very.very/long/url?with=parameter
		sdfsdfasdf http://short.url/without-parameter
		@david asdfasdf`: []linksResult{
			{
				Url:   "http://very-very-very-very-very-very-very-very.very/long/url?with=parameter",
				Title: "http://very-very-very-very-very-very-very-very....",
			},
			{
				Url:   "http://short.url/without-parameter",
				Title: "http://short.url/without-parameter",
			},
		},
	}
	for k, v := range cases {
		res := f.Filter(k)
		vj, _ := json.Marshal(v)
		vres, _ := json.Marshal(res)
		assert.JSONEq(string(vj), string(vres), "Test case %q failed", k)
	}
}

//TestFilterWithInternet works only with Internet
func TestFilterWithInternet(t *testing.T) {
	f := &filter{get: httpGet}
	assert := assert.New(t)
	cases := map[string][]linksResult{
		`@bob @john (success) such a cool feature;
https://twitter.com/jdorfman/status/430511497475670016`: []linksResult{
			{
				Url:   "https://twitter.com/jdorfman/status/430511497475670016",
				Title: "Justin Dorfman on Twitter: \"nice @littlebigdeta...",
			},
		},
	}
	for k, v := range cases {
		res := f.Filter(k)
		vj, _ := json.Marshal(v)
		vres, _ := json.Marshal(res)
		assert.JSONEq(string(vj), string(vres), "Test case %q failed", k)
	}
}
