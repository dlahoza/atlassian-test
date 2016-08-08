package filter_fabric

import (
	log "github.com/Sirupsen/logrus"
	"sync"
)

type FilteredResult []interface{}

type Filterer interface {
	Filter(string) FilteredResult
}

type Filters struct {
	filters map[string]Filterer
	sync.RWMutex
}

func (f *Filters) Add(name string, n Filterer) {
	f.Lock()
	if f.filters == nil {
		f.filters = map[string]Filterer{}
	}
	f.filters[name] = n
	f.Unlock()
	log.WithField("filter", name).Debugf("Added %q filter to catalog", name)
}

func (f *Filters) FilterAll(input string) (output map[string]FilteredResult) {
	f.RLock()
	for fname, _ := range f.filters {
		output[fname] = f.filters[fname].Filter(input)
	}
	f.RUnlock()
	return
}

var filters Filters

func Register(name string, n Filterer) {
	filters.Add(name, n)
}

func FilterAll(input string) map[string]FilteredResult {
	return filters.FilterAll(input)
}
