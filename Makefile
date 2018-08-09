#
# Simple Makefile
#

PROJECT = vcard

VERSION = $(shell grep -m1 "Version = " $(PROJECT).go | cut -d\` -f 2)

BRANCH = $(shell git branch | grep "* " | cut -d\   -f 2)

PKGASSETS = $(shell which pkgassets)

OS = $(shell uname)

EXT =
ifeq ($(OS),Windows)
	EXT = .exe
endif

build: bin/vcard2json$(EXT) 

bin/vcard2json$(EXT): vcard.go cmd/vcard2json/vcard2json.go
	go build -o bin/vcard2json$(EXT) cmd/vcard2json/vcard2json.go

lint:
	golint vcard.go
	golint vcard_test.go
	golint cmd/vcard2json/vcard2json.go

format:
	gofmt -w vcard.go
	gofmt -w vcard_test.go
	gofmt -w cmd/vcard2json/vcard2json.go

test: bin/vcard2json$(EXT)
	go test
	bash test_cmd.bash

status:
	git status

save:
	if [ "$(msg)" != "" ]; then git commit -am "$(msg)"; else git commit -am "Quick Save"; fi
	git push origin $(BRANCH)

clean:
	if [ -d bin ]; then rm -fR bin; fi
	if [ -d dist ]; then rm -fR dist; fi
	if [ -d man ]; then rm -fR man; fi

man: build
	mkdir -p man/man1
	bin/vcard2json -generate-manpage | nroff -Tutf8 -man > man/man1/vcard2json.1

install:
	env GOBIN=$(GOPATH)/bin go install cmd/vcard2json/vcard2json.go


dist/linux-amd64:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=amd64 go build -o dist/bin/vcard2json cmd/vcard2json/vcard2json.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-linux-amd64.zip README.md LICENSE INSTALL.md bin/* 
	rm -fR dist/bin

dist/windows-amd64:
	mkdir -p dist/bin
	env  GOOS=windows GOARCH=amd64 go build -o dist/bin/vcard2json.exe cmd/vcard2json/vcard2json.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-windows-amd64.zip README.md LICENSE INSTALL.md bin/* 
	rm -fR dist/bin

dist/macosx-amd64:
	mkdir -p dist/bin
	env  GOOS=darwin GOARCH=amd64 go build -o dist/bin/vcard2json cmd/vcard2json/vcard2json.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-macosx-amd64.zip README.md LICENSE INSTALL.md bin/* 
	rm -fR dist/bin

dist/raspbian-arm7:
	mkdir -p dist/bin
	env  GOOS=linux GOARCH=arm GOARM=7 go build -o dist/bin/vcard2json cmd/vcard2json/vcard2json.go
	cd dist && zip -r $(PROJECT)-$(VERSION)-raspbian-arm7.zip README.md LICENSE INSTALL.md bin/*
	rm -fR dist/bin

distribute_docs:
	cp -v README.md dist/
	cp -v LICENSE dist/
	cp -v INSTALL.md dist/
	./package-versions.bash > dist/package-versions.txt

release: distribute_docs dist/linux-amd64 dist/windows-amd64 dist/macosx-amd64 dist/raspbian-arm7

website:
	./mk-website.bash

publish:
	./mk-website.bash
	./publish.bash

