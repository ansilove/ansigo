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

AnsiGo requires Go 1.11 or later.

Build AnsiGo from the repository root:

	go build

This creates the `ansigo` executable in the current directory. To install it
in your Go binary directory:

	go install

Alternatively, you can easily cross-compile binaries for other systems. See the `Cross-compiling AnsiGo binaries` section for instructions.

## Usage

AnsiGo takes the input file as parameter:

	ansigo inputfile

## Cross-compiling AnsiGo binaries

Go can cross-compile pure Go programs directly. Run the following commands from
the repository root.

Building Linux binaries:

	GOOS=linux GOARCH=amd64 go build -o ansigo-linux-amd64
	GOOS=linux GOARCH=386 go build -o ansigo-linux-386

Building Windows binaries:

	GOOS=windows GOARCH=amd64 go build -o ansigo-windows-amd64.exe
	GOOS=windows GOARCH=386 go build -o ansigo-windows-386.exe

## License

AnsiGo is released under the BSD 2-Clause license. See `LICENSE` file for details.

## Author

AnsiGo is developed by Frederic Cambus.

- Site: https://www.cambus.net

## Resources

GitHub: https://github.com/ansilove/ansigo

Sister project: https://www.ascii-codes.com
