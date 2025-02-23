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

	if flag.NArg() == 0 {
		panic("missing path to fil")
	}

	fp := os.Args[len(os.Args)-1]

	if *wantBytes {
		count := countBytes(fp)
		fmt.Printf("%d %s\n", count, fp)
		return
	}

	if *wantLines {
		count := countLines(fp)
		fmt.Printf("%d %s\n", count, fp)
		return
	}

	if *wantWords {
		count := countWords(fp)
		fmt.Printf("%d %s\n", count, fp)
		return
	}

	if *wantCharacters {
		count := countCharacters(fp)
		fmt.Printf("%d %s\n", count, fp)
		return
	}

	fmt.Printf("%d %d %d %s\n", countLines(fp), countWords(fp), countBytes(fp), fp)
}

func countBytes(fp string) int {
	file, err := os.Open(fp)
	if err != nil {
		panic(fmt.Sprintf("cannot open file, %s", err))
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		// Generally you'd want to return the error but for this trivial application we can just
		// return a negative number which is a reasonable indication that something is wrong
		return -1
	}

	return int(info.Size())
}

func countLines(fp string) int {
	file, err := os.Open(fp)
	if err != nil {
		panic(fmt.Sprintf("cannot open file, %s", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count
}

func countWords(fp string) int {
	file, err := os.Open(fp)
	if err != nil {
		panic(fmt.Sprintf("cannot open file, %s", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count
}

func countCharacters(fp string) int {
	file, err := os.Open(fp)
	if err != nil {
		panic(fmt.Sprintf("cannot open file, %s", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count
}
