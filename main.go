package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func Interpret(codestr string) int {
	code := strings.Split(codestr, "")
	startTime := MilliTime()
	codePos := 0
	var mem [30_000]uint8
	memPos := 0
	var loopStack []int

	reader := bufio.NewReader(os.Stdin)

	for codePos < len(code) {
		char := code[codePos]
		switch char {
		case "+":
			mem[memPos]++
		case "-":
			mem[memPos]--
		case ">":
			memPos++
		case "<":
			memPos--
		case ".":
			fmt.Print(string(rune(mem[memPos])))
		case ",":
			value, _ := reader.ReadByte()
			mem[memPos] = uint8(value)
		// This part is copied from my old terrible code, I should probably switch this out with new better code.
		case "[":
			if mem[memPos] != 0 {
				loopStack = append(loopStack, codePos)
			} else {
				searchSlice := code[codePos:]
				searchIndex := 0
				openings := 0
				for true {
					if searchSlice[searchIndex] == "[" {
						openings++
					}
					if searchSlice[searchIndex] == "]" && openings < 2 {
						break
					} else if searchSlice[searchIndex] == "]" {
						openings--
					}
					searchIndex++
				}
				codePos += searchIndex
			}
		case "]":
			if mem[memPos] != 0 {
				codePos = loopStack[len(loopStack)-1]
			} else {
				loopStack = loopStack[:len(loopStack)-1]
			}
		}
		// End of old terrible code
		codePos++
	}

	return int(MilliTime()) - int(startTime)
}

func MilliTime() int64 {
	return time.Now().UnixNano() / 1e6
}

func IsValidPath(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You have to provide a valid file path. Like `./main.bf`")
		return
	}

	bffile := os.Args[1]

	if !IsValidPath(bffile) {
		fmt.Println(bffile, "Isn't a valid path. You have to provide a valid file path. Like `./main.bf`")
		return
	}

	data, err := ioutil.ReadFile(bffile)
	if err != nil {
		fmt.Println("Couldn't read file", bffile)
		return
	}

	execTime := Interpret(string(data))

	green := "\033[32m"
	reset := "\033[0m"

	fmt.Println(green, "\nRan in", execTime, "milliseconds")
	fmt.Print(reset)
}
