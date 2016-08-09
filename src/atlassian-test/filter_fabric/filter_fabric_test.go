package filter_fabric

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type filterMock struct{}

func (f *filterMock) Filter(input string) FilteredResult {
	return FilteredResult{input}
}

var catalogMock = map[string]Filterer{
	"test":  new(filterMock),
	"test2": new(filterMock),
}

const messageToFilterMock = "testMessage"

var fullFilteredResultMock = FullFilteredResult{
	"test":  {messageToFilterMock},
	"test2": {messageToFilterMock},
}

func TestRegister(t *testing.T) {
	assert := assert.New(t)
	for k, v := range catalogMock {
		Register(k, v)
	}
	assert.EqualValues(catalogMock, filters.filters)
}

func TestFilterAll(t *testing.T) {
	TestRegister(t)
	assert := assert.New(t)
	res := FilterAll(messageToFilterMock)
	assert.EqualValues(fullFilteredResultMock, res)
}
