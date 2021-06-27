package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const usageMessage = "usage: bigdigits [-b|--bar] <whole-number>\n" +
	"-b --bar draw an underbar and an overbar"
var helpFlag = flag.Bool("help", false, usageMessage)
var barFlag = flag.Bool("bar", false, "-b --bar draw an underbar and an overbar")

func init() {
	flag.BoolVar(helpFlag, "h", false, usageMessage)
	flag.BoolVar(barFlag, "b", false, "-b --bar draw an underbar and an overbar")
}

func main() {

	flag.Parse()
	if *helpFlag {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if len(os.Args) == 1 {
		fmt.Printf("usage: %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringOfDigits := ""
	if *barFlag {
		stringOfDigits = os.Args[2]
	} else {
		stringOfDigits = os.Args[1]
	}
	lenLine := 0
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + "  "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		lenLine = len(line)
		if row == 0 && *barFlag {
			fmt.Println(strings.Repeat("*", lenLine - 1))
		}
		fmt.Println(line)
		}
	if *barFlag {
		fmt.Println(strings.Repeat("*", lenLine-1))
	}
}

var bigDigits = [][]string{
	{"  000  ", " 0   0 ", "0     0", "0     0", "0     0", " 0   0 ", "  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",  "   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}