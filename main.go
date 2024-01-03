package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func tabsToSpaces(lines []string) {
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, "\t", "    ")
	}
}

func calculateMaxWidth(lines []string) int {
	var width int = 0
	for _, line := range lines {
		if len(line) > width {
			width = len(line)
		}
	}
	return width
}

func addSpaceStrings(lines []string) {
	var width int = calculateMaxWidth(lines)
	for i, line := range lines {
		lines[i] = line + strings.Repeat(" ", width-len(line))
	}
}

func addBorderStrings(lines []string) (topRow string, botRow string) {
	var width int = calculateMaxWidth(lines)

	for i, line := range lines {
		lines[i] = fmt.Sprintf("│ %s │", line)
	}

	var row = strings.Repeat("─", width+2)
	topRow = fmt.Sprintf("┌%s┐", row)
	botRow = fmt.Sprintf("└%s┘", row)

	return topRow, botRow
}

func main() {
	info, _ := os.Stdin.Stat()

	if info.Mode()&os.ModeCharDevice != 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: fortune | gocowsay")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	var lines []string

	for {
		line, _, err := reader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	const cow = `
        \  ^__^
         \ (oo)\_______
           (__)\       )\
                ||--u||  \
                ||   ||`

	tabsToSpaces(lines)
	addSpaceStrings(lines)
	topRow, bottomRow := addBorderStrings(lines)

	//Print Balloon
	fmt.Println(topRow)
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println(bottomRow)

	//Cow say it !!
	fmt.Println(cow[1:])
}
