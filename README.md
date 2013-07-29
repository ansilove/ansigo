# AnsiGo

## Description

AnsiGo is a simple ANSi to PNG converter written in pure Go. It converts files
containing ANSi sequences (.ANS) into PNG images.

For a multi-format general purposes converter and library, check out Ansilove 
instead : http://ansilove.sourceforge.net

## Features

- ANSi (.ANS) format support
- Small output file size (4-bit PNG)
- 80x25 font support : IBM PC (Code page 437) charset

## Installation

Build and install with the `go` tool :

	go build ansigo
	go install ansigo

Alternatively, you can download pre-built binaries (32-bit or 64-bit), check
the directory corresponding to your operating system. Currently, only Windows
binaries are available.

## Cross-compiling AnsiGo binaries for Windows

Building Go for the required platform :

	cd /usr/local/go/
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ./make.bash

Building Windows binaries :

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o windows/amd64/ansigo.exe ansigo
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o windows/i386/ansigo.exe ansigo

## License

AnsiGo is released under the MIT license. See `LICENSE` file for details.

## Author

AnsiGo is developed by Frederic Cambus

- Site : http://www.cambus.net
- Twitter: http://twitter.com/fcambus

## Resources

Project Homepage : http://fcambus.github.io/ansigo/

Sister project : http://www.ascii-codes.com
