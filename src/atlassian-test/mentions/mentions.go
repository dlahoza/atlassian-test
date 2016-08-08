package mentions

import (
	fabric "atlassian-test/filter_fabric"
	log "github.com/Sirupsen/logrus"
	"regexp"
)

const FILTER_NAME = "mentions"

// `(?m)` for multiline mode
var filter_re = regexp.MustCompile(`(?m)@(\w+)`)

// Empty struct for filter object
type filter struct {
}

// Filter filters input message and returns founded objects
func (f *filter) Filter(input string) (output fabric.FilteredResult) {
	//Looking for matches regarding to filter expression
	m := filter_re.FindAllStringSubmatch(input, -1)
	log.WithField("filter", FILTER_NAME).Debugf("Found %d matches", len(m))
	for _, s := range m {
		//Adding founded objects
		output = append(output, s[1])
	}
	return
}

func init() {
	//Registering in filters catalog
	fabric.Register(FILTER_NAME, new(filter))
}
