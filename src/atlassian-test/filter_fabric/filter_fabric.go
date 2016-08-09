package filter_fabric

import (
	log "github.com/Sirupsen/logrus"
	"sync"
)

// Regular result returned by Filterer
type FilteredResult []interface{}

// Result from some Filterers
type FullFilteredResult map[string]FilteredResult

// Regular filter interface
type Filterer interface {
	Filter(string) FilteredResult
}

// Filters catalog type
type Filters struct {
	filters map[string]Filterer
	sync.RWMutex
}

// Add function adds some filter to catalog
func (f *Filters) Add(name string, n Filterer) {
	// Locking mutex to write into the map
	f.Lock()
	// If catalog wasn't initialized
	if f.filters == nil {
		// Initialize it
		f.filters = map[string]Filterer{}
	}
	// Adding filter to the catalog
	f.filters[name] = n
	f.Unlock()
	log.WithField("filter", name).Debugf("Added %q filter to catalog", name)
}

//FilterAll runs filtering in all filters in parallel
func (f *Filters) FilterAll(input string) (output map[string]FilteredResult) {
	type nameResult struct {
		name   string
		result FilteredResult
	}
	f.RLock()
	var count int
	// Channel for parallel results
	c := make(chan nameResult)
	// Do not forget to close channels
	defer close(c)
	output = map[string]FilteredResult{}
	// Starting workers with filters
	for fname := range f.filters {
		count++
		go func(name string, filter Filterer) {
			c <- nameResult{name, filter.Filter(input)}
		}(fname, f.filters[fname])
	}
	f.RUnlock()
	// Receiving results
	for i := 0; i < count; i++ {
		res := <-c
		output[res.name] = res.result
	}
	return
}

// Global filters catalog
var filters Filters

// Register registers filter in the global catalog
func Register(name string, n Filterer) {
	filters.Add(name, n)
}

// FilterAll runs FilterAll of the filters in global catalog
func FilterAll(input string) map[string]FilteredResult {
	return filters.FilterAll(input)
}
