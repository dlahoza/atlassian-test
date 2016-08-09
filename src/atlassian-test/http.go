package main

import (
	fabric "atlassian-test/filter_fabric"
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

// filterHandler
func filterHandler(w http.ResponseWriter, r *http.Request) {
	// Checking that HTTP method is POST
	if r.Method != "POST" {
		w.WriteHeader(404)
		return
	}
	// Reading request payload
	payload, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.WithError(err).Error("Can't read payload from HTTP request")
		w.WriteHeader(500)
		w.Write([]byte("Error when reading request"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	j := json.NewEncoder(w)
	// Writing filtered and encoded to JSON response
	err = j.Encode(fabric.FilterAll(string(payload)))
	if err != nil {
		log.WithError(err).Error("Can't encode JSON and write HTTP response")
		return
	}
}
