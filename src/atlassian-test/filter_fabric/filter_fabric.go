package filter_fabric

import "sync"

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
}

var filters Filters

func Register(name string, n Filterer) {
	filters.Add(name, n)
}
