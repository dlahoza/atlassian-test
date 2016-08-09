package links

import (
	fabric "atlassian-test/filter_fabric"
	log "github.com/Sirupsen/logrus"
	"regexp"
)

const FILTER_NAME = "links"

// `(?m)` for multiline mode
var filter_re = regexp.MustCompile(`(?m)https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,4}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)

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
