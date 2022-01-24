# AnsiGo

## Description

AnsiGo is a simple ANSi to PNG converter written in pure Go. It converts files
containing ANSi sequences (.ANS) into PNG images.

For a multi-format general purposes converter and library, check out Ansilove
instead: https://www.ansilove.org

## Features

- ANSi (.ANS) format support
- Small output file size (4-bit PNG)
- 80x25 font support: IBM PC (Code page 437) charset

## Installation

Build and install with the `go` tool:

	go build ansigo
	go install ansigo

Alternatively, you can easily cross-compile binaries for other systems. See the `Cross-compiling AnsiGo binaries` section for instructions.

## Usage

AnsiGo takes the input file as parameter:

	ansigo inputfile

## Cross-compiling AnsiGo binaries

Building Go for the required platform:

	cd /usr/local/go/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ./make.bash

Building Linux binaries:

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o linux/amd64/ansigo ansigo
	CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o linux/i386/ansigo ansigo

Building Windows binaries:

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o windows/amd64/ansigo.exe ansigo
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o windows/i386/ansigo.exe ansigo

## License

AnsiGo is released under the BSD 2-Clause license. See `LICENSE` file for details.

## Author

AnsiGo is developed by Frederic Cambus.

- Site: https://www.cambus.net

## Resources

GitHub: https://github.com/ansilove/ansigo

Sister project: https://www.ascii-codes.com
