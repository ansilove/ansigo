// AnsiGo 1.00 (c) by Frederic Cambus 2012-2015
// http://www.github.com/fcambus/ansigo
//
// Created:      2012/02/14
// Last Updated: 2015/03/29
//
// AnsiGo is released under the BSD 3-Clause license.
// See LICENSE file for details.
package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("-------------------------------------------------------------------------------\n                  AnsiGo 1.00 (c) by Frederic CAMBUS 2012-2015\n-------------------------------------------------------------------------------\n")

	// Check input parameters and show usage
	if len(os.Args) != 2 {
		fmt.Println("USAGE:    ansigo inputfile\n")
		fmt.Println("EXAMPLES: ansigo ansi.ans\n")
		os.Exit(1)
	}

	input := os.Args[1]
	output := input + ".png"

	fmt.Println("Input File:", input)
	fmt.Println("Output File:", output)

	var ansi Ansi
	ansi.SetPalette()
	ansi.SetFont()

	// Load input file
	data, err := ioutil.ReadFile(input)
	if err != nil {
		fmt.Println("\nERROR: Can't open or read", input, "\n")
		os.Exit(1)
	}

	// Process ANSI
	for i := 0; i < len(data); i++ {
		ansi.character = data[i]

		// 80th column wrapping
		if ansi.positionX == 80 {
			ansi.positionY++
			ansi.positionX = 0
		}

		// CR (Carriage Return)
		if ansi.character == '\r' {
			continue
		}

		// LF (Line Feed)
		if ansi.character == '\n' {
			ansi.positionY++
			ansi.positionX = 0
			continue
		}

		// HT (Horizontal Tabulation)
		if ansi.character == '\t' {
			ansi.positionX += 8
			continue
		}

		// SUB (Substitute)
		if ansi.character == '\x1a' {
			break
		}

		// ANSI Sequence : ESC (Escape) + [
		if ansi.character == '\x1b' && data[i+1] == '[' {
			ansiSequence := []byte{}

		sequence:
			for j := 0; j < 12; j++ {
				ansiSequenceCharacter := data[i+2+j]

				switch ansiSequenceCharacter {
				case 'H', 'f': // Cursor Position
					ansiSequenceValues := strings.SplitN(string(ansiSequence), ";", -1)

					valueY, _ := strconv.Atoi(ansiSequenceValues[0])
					ansi.positionY = valueY - 1

					valueX, _ := strconv.Atoi(ansiSequenceValues[1])
					ansi.positionX = valueX - 1

					i += j + 2
					break sequence

				case 'A': // Cursor Up
					valueY, _ := strconv.Atoi(string(ansiSequence))
					if valueY == 0 {
						valueY++
					}

					ansi.positionY = ansi.positionY - valueY

					i += j + 2
					break sequence

				case 'B': // Cursor Down
					valueY, _ := strconv.Atoi(string(ansiSequence))
					if valueY == 0 {
						valueY++
					}

					ansi.positionY = ansi.positionY + valueY

					i += j + 2
					break sequence

				case 'C': // Cursor Forward
					valueX, _ := strconv.Atoi(string(ansiSequence))
					if valueX == 0 {
						valueX++
					}

					ansi.positionX = ansi.positionX + valueX
					if ansi.positionX > 80 {
						ansi.positionX = 80
					}

					i += j + 2
					break sequence

				case 'D': // Cursor Backward
					valueX, _ := strconv.Atoi(string(ansiSequence))
					if valueX == 0 {
						valueX++
					}

					ansi.positionX = ansi.positionX - valueX
					if ansi.positionX < 0 {
						ansi.positionX = 0
					}

					i += j + 2
					break sequence

				case 's': // Save Cursor Position
					ansi.savedPositionY = ansi.positionY
					ansi.savedPositionX = ansi.positionX

					i += j + 2
					break sequence

				case 'u': // Restore Cursor Position
					ansi.positionY = ansi.savedPositionY
					ansi.positionX = ansi.savedPositionX

					i += j + 2
					break sequence

				case 'J': // Erase Display
					value, _ := strconv.Atoi(string(ansiSequence))

					if value == 2 {
						ansi.buffer = nil

						ansi.positionX = 0
						ansi.positionY = 0
						ansi.sizeX = 0
						ansi.sizeY = 0
					}

					i += j + 2
					break sequence

				case 'm': // Set Graphic Rendition
					ansiSequenceValues := strings.SplitN(string(ansiSequence), ";", -1)

					for j := 0; j < len(ansiSequenceValues); j++ {
						valueColor, _ := strconv.Atoi(ansiSequenceValues[j])

						switch valueColor {
						case 0:
							ansi.colorBackground = 0
							ansi.colorForeground = 7
							ansi.bold = false

						case 1:
							ansi.colorForeground += 8
							ansi.bold = true

						case 5:
							ansi.colorBackground += 8

						case 30, 31, 32, 33, 34, 35, 36, 37:
							ansi.colorForeground = valueColor - 30
							if ansi.bold {
								ansi.colorForeground += 8
							}

						case 40, 41, 42, 43, 44, 45, 46, 47:
							ansi.colorBackground = valueColor - 40
						}
					}

					i += j + 2
					break sequence

				case 'h', 'l': // Skipping Set Mode And Reset Mode Sequences
					i += j + 2
					break sequence
				}

				ansiSequence = append(ansiSequence, ansiSequenceCharacter)
			}
		} else {
			// Record Number Of Columns And Lines Used
			if ansi.positionX > ansi.sizeX {
				ansi.sizeX = ansi.positionX
			}
			if ansi.positionY > ansi.sizeY {
				ansi.sizeY = ansi.positionY
			}

			// Write Current Character Info In A Temporary Array
			ansi.buffer = append(ansi.buffer, Character{ansi.colorBackground, ansi.colorForeground, ansi.positionX, ansi.positionY, ansi.character})

			ansi.positionX++
		}
	}

	// Allocate Image Buffer Memory
	canvasSize := image.Rect(0, 0, 640, (ansi.sizeY+1)*16)
	canvas := image.NewRGBA(canvasSize)

	// Draw The Canvas Background
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{ansi.palette[0]}, image.ZP, draw.Src)

	// Render ANSI
	for i := 0; i < len(ansi.buffer); i++ {
		character := ansi.buffer[i]

		// Set Background
		draw.Draw(canvas, image.Rect(character.positionX*8, character.positionY*16, character.positionX*8+8, character.positionY*16+16), &image.Uniform{ansi.palette[character.colorBackground]}, image.ZP, draw.Src)

		// Draw Character
		for line := 0; line < 16; line++ {
			for column := 0; column < 8; column++ {
				if (ansi.font[line+(int(character.code)*16)] & (0x80 >> uint(column))) != 0 {
					canvas.Set(character.positionX*8+column, line+character.positionY*16, ansi.palette[character.colorForeground])
				}
			}
		}
	}

	// Create Output File
	outputFile, err := os.Create(output)
	if err != nil {
		fmt.Println("ERROR: Can't create ouput file.")
		os.Exit(1)
	}

	// Encode PNG image
	if err = png.Encode(outputFile, canvas); err != nil {
		fmt.Println("ERROR: Can't encode PNG file.")
		os.Exit(1)
	}

	outputFile.Close()

	fmt.Println("\nSuccessfully created file", output, "\n")
}
