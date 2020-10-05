all: run

front:
	sqlitebrowser utem-gsp-test.db

test:
	go test test/endpoints_test.go -count=1

build:
	go build .

clear:
	rm utem-gsp

refresh:
	mysql -u root -p -e "drop database utem_gsp; create database utem_gsp;"

run:
	go run .
