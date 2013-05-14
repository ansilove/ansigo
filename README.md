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

## License

AnsiGo is released under the MIT license. See `LICENSE` file for details.

## Author

AnsiGo is developed by Frederic Cambus

- Site : http://www.cambus.net
- Twitter: http://twitter.com/fcambus

## Resources

Project Homepage : https://github.com/fcambus/ansigo

Sister project : http://www.ascii-codes.com
