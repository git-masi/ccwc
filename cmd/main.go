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

	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Printf("cannot open file, %s", err)
		os.Exit(1)
	}
	defer file.Close()

	if *wantBytes {
		count := countBytes(file)
		fmt.Printf("%d %s\n", count, pathToFile)
		return
	}

	if *wantLines {
		count := countLines(file)
		fmt.Printf("%d %s\n", count, pathToFile)
		return
	}

	if *wantWords {
		count := countWords(file)
		fmt.Printf("%d %s\n", count, pathToFile)
		return
	}

	if *wantCharacters {
		count := countCharacters(file)
		fmt.Printf("%d %s\n", count, pathToFile)
		return
	}

	fmt.Printf("%d %d %d %s\n", countLines(file), countWords(file), countBytes(file), pathToFile)
}

func countBytes(file *os.File) int {
	info, err := file.Stat()
	if err != nil {
		// Generally you'd want to return the error but for this trivial application we can just
		// return a negative number which is a reasonable indication that something is wrong
		return -1
	}

	return int(info.Size())
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
