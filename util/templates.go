package util

// String literals do not process escape sequences

const MainTemplate = `package main

import (
	. "fmt"
)

func main() {
	Println("This is an example program generated by GoLink.")
}

`
const GoadTemplate = `#!/bin/bash

# Where is this script located?
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

# Project details
pkg="PACKAGE"
name="NAME"

(
	cd "$DIR"

	export GOPATH="$PWD"/.gopath/
	export BASEDIR="$PWD"

	go build -o $name $pkg/$name/main
)

`
