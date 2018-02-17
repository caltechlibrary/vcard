#!/bin/bash

function assert_exists() {
    if [ "$#" != "2" ]; then
        echo "wrong number of parameters for $1, $@"
        exit 1
    fi
    if [[ ! -f "$2" && ! -d "$2" ]]; then
        echo "$1: $2 does not exists"
        exit 1
    fi
}

function assert_equal() {
    if [ "$#" != "3" ]; then
        echo "wrong number of parameters for $1, $@"
        exit 1
    fi
    if [ "$2" != "$3" ]; then
        echo "$1: expected |$2| got |$3|"
        exit 1
    fi
}

#
# Tests
#

function test_vcard2json() {
    EXT=".exe"
    OS=$(uname)
    if [ "$OS" != "Windows" ]; then
        EXT=""
    fi
    echo "Testing for bin/vcard2json${EXT}"
    if [[ ! -f "bin/vcard2json${EXT}" ]]; then
        go build -o "bin/vcard2json${EXT}" cmd/vcard2json/vcard2json.go
    fi
    echo "Not implemented!"
    exit 1
}

echo "Testing command line tool"
test_vcard2json
echo 'Success!'
