package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	wantBytes := flag.Bool("c", false, "Return the number of bytes in a file")
	wantLines := flag.Bool("l", false, "Return the number of lines in a file")
	wantWords := flag.Bool("w", false, "Return the number of words in a file")
	wantCharacters := flag.Bool("m", false, "Return the number of characters in a file")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("missing path to file")
		os.Exit(1)
	}

	pathToFile := os.Args[len(os.Args)-1]

	if *wantBytes {
		count, err := countBytes(pathToFile)
		if err != nil {
			fmt.Printf("cannot get file info, %s", err)
			os.Exit(1)
		}

		fmt.Printf("%d %s\n", count, pathToFile)
	}

	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Printf("cannot open file, %s", err)
		os.Exit(1)
	}
	defer file.Close()

	if *wantLines {
		count := countLines(file)

		fmt.Printf("%d %s\n", count, pathToFile)
	}

	if *wantWords {
		count := countWords(file)

		fmt.Printf("%d %s\n", count, pathToFile)
	}

	if *wantCharacters {
		count := countCharacters(file)

		fmt.Printf("%d %s\n", count, pathToFile)
	}
}

func countBytes(pathToFile string) (int, error) {
	info, err := os.Stat(pathToFile)
	if err != nil {
		return -1, err
	}

	return int(info.Size()), nil
}

func countLines(file *os.File) int {
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count
}

func countWords(file *os.File) int {
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		for _, str := range strings.Fields(strings.TrimSpace(scanner.Text())) {
			if str != "" {
				count++
			}
		}
	}

	return count
}

func countCharacters(file *os.File) int {
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		count += utf8.RuneCount(scanner.Bytes())
	}

	return count
}
