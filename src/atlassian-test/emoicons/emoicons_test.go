package emoicons

import (
	fabric "atlassian-test/filter_fabric"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmoicons(t *testing.T) {
	f := new(filterEmoicons)
	assert := assert.New(t)
	cases := map[string]fabric.FilteredResult{
		`Good morning! (megusta) (coffee)`: {"megusta", "coffee"},
		`Multiline string with 2 emoicons
		(test) asdfsdf
		sdfsdfasdf
		(test2) asdfasdf`: {"test", "test2"},
		`Another multiline string with 2 emoicons and one broken
		(test) asdfsdf
		sdfsdfasdf(bro
		ken)
		(test2) asdfasdf`: {"test", "test2"},
	}
	for k, v := range cases {
		res := f.Filter(k)
		assert.EqualValues(v, res, "Test case %q failed", k)
	}
}
