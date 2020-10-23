dbname=utem_gsp.db

all: run

front-sqlite:
	sqlitebrowser $(dbname)

front-mysql:
	mysql-workbench

test:
	go test test/jwt_test.go -count=1

build:
	go build .

copy:
	cp utem-gsp.db utem-gsp-test.db

clear:
	rm utem-gsp

refresh-mysql:
	mysql -u root -p -e "drop database utem_gsp; create database utem_gsp;"

run:
	go run .

log:
	grc tail -f logs.log

.PHONY: test
