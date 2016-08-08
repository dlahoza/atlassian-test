package mentions

import (
	fabric "atlassian-test/filter_fabric"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	f := new(filter)
	assert := assert.New(t)
	cases := map[string]fabric.FilteredResult{
		`@chris you around?`: {"chris"},
		`Multiline string with 2 mentions
		@chris asdfsdf
		sdfsdfasdf
		@david asdfasdf`: {"chris", "david"},
		`Another multiline string with 2 mentions and one broken
		@test asdfsdf
		sdfsdfasdf @
		broken
		@test2 asdfasdf`: {"test", "test2"},
	}
	for k, v := range cases {
		res := f.Filter(k)
		assert.EqualValues(v, res, "Test case %q failed", k)
	}
}
