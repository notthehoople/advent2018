package main

import (
	"fmt"
	"bufio"
	"os"
	"flag"
	"strings"
	"strconv"
)

// Read the text file passed in by name into a array of strings
// Returns the array as the first return variable
func readLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
	  return nil, err
	}
	defer file.Close()
  
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
	  lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func printStringArray(tempString []string) {
	// Loop through the array and print each line
	for i:= 0; i < len(tempString); i++ {
		fmt.Println(tempString[i])
	}
}

// Handles everything needed to work out the fabric claim (Day03 part A) and Good Elf (Day03 part B)
func fabricClaim(fileName string, solutionPart string) int {

	// A claim like #123 @ 3,2: 5x4 means that claim ID 123 specifies a rectangle 3 inches from the
	// left edge, 2 inches from the top edge, 5 inches wide, and 4 inches tall

	// We read through the file once, reading in each Elf's fabric claim
	// Each line is #<num> @ <x>,<y>: <a>x<b>
	//   #<num> - claim number
	//   @ - seperator. We don't care about this
	//   <x> - x inches from the left edge
	//   <y> - y inches from the top edge
	//   <a> - a inches wide
	//   <b> - b inches tall

	var resultVar int = 0					// Defining the overall result Variable
	var a, b, x, y int = 0, 0, 0, 0			// Co-ords and size of claim
	var lengthx, lengthy int = 0, 0
	var elfNumber int = 0

	fabricMap := [2000][2000]int{}			// Santa's fabric
	elfList := [2000]int{}					// The list of elves

	// Read contents of file into a string array
	fileContents, _ := readLines(fileName)

	// Loop through the string array; break into component parts then apply to our fabricMap

	for i := 0; i < len(fileContents); i++ {
		words := strings.Fields(fileContents[i])

		for j := 0; j < len(words); j++ {
			switch j {
			case 0:	// First entry is the claim number
				tempElf := strings.Split(words[j], "#")
				elf, _ := strconv.Atoi(tempElf[1])
				elfNumber = elf
			case 1: // Second entry is the '@'
			case 2: // Third entry is the start co-ordinates <x>,<y>
				tempNumber := strings.Split(words[j], ":")
				numberString := strings.Split(tempNumber[0], ",")
				x, _ = strconv.Atoi(numberString[0])
				y, _ = strconv.Atoi(numberString[1])
			case 3: // Fourth entry is the size <a>x<b>
				numberString := strings.Split(words[j], "x")
				a, _ = strconv.Atoi(numberString[0])
				b, _ = strconv.Atoi(numberString[1])
			}
		}

		// Now we have the variables we need it's time to modify the fabricMap
		// Loop through the fabricMap adding the elf number to each relevant area found
		// If an elf number is already there, replace it with '-1' and mark the elf as dud in elfList
		// Once we've built the map, we'll count the number of '-1's for part A

		// k is our "y" axis in the map
		for k := y; k < y + b; k++ {
			// l is the "x" axis in the map. We need to modify from x through to x+a
			for l := x; l < x + a; l++ {
				if fabricMap[l][k] == 0 {
					// This is the first time this square has been used. Allocate it to the Elf
					fabricMap[l][k] = elfNumber
				} else {
					if fabricMap[l][k] > 0 {
						// The fabricMap is already used by another Elf. Mark that elf as bad, our elf as bad
						// and mark the fabricMap as being used by multiple elves
						elfList[fabricMap[l][k]] = -1
						fabricMap[l][k] = -1
						elfList[elfNumber] = -1
					} else {
						if fabricMap[l][k] < 0 {
							// If the fabricMap is already a duplicate entry, just mark the elf as bad
							elfList[elfNumber] = -1
						}
					}
				}
			}
		}
	}	

	if solutionPart == "a" {
		// PART A: Once complete, our result is:
		//    - loop through the fabricMap
		//    - count the number of entries > 1
		resultVar = 0
		lengthy = len(fabricMap)
		lengthx = len(fabricMap[0])
		for k := 0; k < lengthx; k++ {
			// l is the "x" axis in the map. We need to modify from x through to x+a
			for l := 0; l < lengthy; l++ {
				if fabricMap[l][k] < 0 {
					resultVar++
				}
			}
		}
	} else {
		// PART B: Once complete, our answer is the only Elf in elfList that isn't -1
		for i := 1; i <= elfNumber; i++ {
			if elfList[i] != -1 {
				resultVar = i
			}
		}
	}

	return resultVar
}

func main() {
	fileNamePtr := flag.String("file", "input1.txt", "A filename containing input strings")
	execPartPtr := flag.String("part", "a", "Which part of day02 do you want to calc (a or b)")

	flag.Parse()

	switch *execPartPtr {
	case "a":
		fmt.Println("Square inches in two or more claims:", fabricClaim(*fileNamePtr, *execPartPtr))
	case "b":
		fmt.Println("The only good Elf is:", fabricClaim(*fileNamePtr, *execPartPtr))
	}
}