all: run

front:
	sqlitebrowser utem-gsp-test.db

test:
	go test test/endpoints_test.go -count=1

run:
	go run .
