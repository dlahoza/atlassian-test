# atlassian-test
[![Build Status](https://travis-ci.org/DLag/atlassian-test.svg?branch=master)](https://travis-ci.org/DLag/atlassian-test)

Atlassian test task.

It is a microservice with HTTP RESTful API with one endpoint.

### Configuration with ENV
```
LISTEN=":8080"
LOGLEVEL="debug"
```

### Build
```
export GOPATH=$(pwd)
cd src/atlassian
go build
```

### Test
```
# export GOPATH=$(pwd)
# cd src/atlassian
# go test -v -cover ./filter_fabric ./mentions ./emoicons ./links .
=== RUN   TestRegister
--- PASS: TestRegister (0.00s)
=== RUN   TestFilterAll
--- PASS: TestFilterAll (0.00s)
PASS
coverage: 100.0% of statements
ok  	atlassian-test/filter_fabric	0.018s	coverage: 100.0% of statements
=== RUN   TestFilter
--- PASS: TestFilter (0.00s)
PASS
coverage: 100.0% of statements
ok  	atlassian-test/mentions	0.015s	coverage: 100.0% of statements
=== RUN   TestFilter
--- PASS: TestFilter (0.00s)
PASS
coverage: 100.0% of statements
ok  	atlassian-test/emoicons	0.028s	coverage: 100.0% of statements
=== RUN   TestFilter
--- PASS: TestFilter (0.00s)
=== RUN   TestFilterWithInternet
--- PASS: TestFilterWithInternet (0.92s)
PASS
coverage: 95.0% of statements
ok  	atlassian-test/links	0.948s	coverage: 95.0% of statements
=== RUN   TestFilter
--- PASS: TestFilter (0.00s)
PASS
coverage: 33.3% of statements
ok  	atlassian-test	0.013s	coverage: 33.3% of statements
```

### Use it
```
# LOGLEVEL=debug ./atlassian-test &
# curl -v -XPOST http://localhost:8080/filter --data "\@bob @john (success) such a cool feature\;\
   https://twitter.com/jdorfman/status/430511497475670016"
{"emoicons":["success"],"links":[{"url":"https://twitter.com/jdorfman/status/430511497475670016","title":"Justin Dorfman on Twitter: \"nice @littlebigdeta..."}],"mentions":["bob","john"]}
```

