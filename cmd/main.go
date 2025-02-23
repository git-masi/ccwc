package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	wantBytes := flag.Bool("c", false, "Return the number of bytes in a file")
	wantLines := flag.Bool("l", false, "Return the number of lines in a file")
	wantWords := flag.Bool("w", false, "Return the number of words in a file")
	wantCharacters := flag.Bool("m", false, "Return the number of characters in a file")
	flag.Parse()

	fp := ""

	if flag.NArg() > 0 {
		fp = os.Args[len(os.Args)-1]
	}

	getInput := initGetInput(fp)

	if *wantBytes {
		count := countBytes(getInput)
		fmt.Printf("%d %s\n", count, fp)
		return
	}

	if *wantLines {
		count := countLines(getInput)
		fmt.Printf("%d %s\n", count, fp)
		return
	}

	if *wantWords {
		count := countWords(getInput)
		fmt.Printf("%d %s\n", count, fp)
		return
	}

	if *wantCharacters {
		count := countCharacters(getInput)
		fmt.Printf("%d %s\n", count, fp)
		return
	}

	fmt.Printf("%d %d %d %s\n", countLines(getInput), countWords(getInput), countBytes(getInput), fp)
}

type inputGetter func() (file *os.File, close func() error)

func initGetInput(fp string) inputGetter {
	return func() (file *os.File, close func() error) {
		if fp == "" {
			return os.Stdin, func() error { return nil }
		}

		file, err := os.Open(fp)
		if err != nil {
			panic(fmt.Sprintf("cannot open file, %s", err))
		}

		return file, file.Close
	}
}

func countBytes(getInput inputGetter) int {
	file, close := getInput()
	defer close()

	info, err := file.Stat()
	if err != nil {
		// Generally you'd want to return the error but for this trivial application we can just
		// return a negative number which is a reasonable indication that something is wrong
		return -1
	}

	return int(info.Size())
}

func countLines(getInput inputGetter) int {
	file, close := getInput()
	defer close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count
}

func countWords(getInput inputGetter) int {
	file, close := getInput()
	defer close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count
}

func countCharacters(getInput inputGetter) int {
	file, close := getInput()
	defer close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count
}
