package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
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
		info, err := os.Stat(pathToFile)
		if err != nil {
			fmt.Printf("cannot get file info, %s", err)
			os.Exit(1)
		}

		fmt.Printf("%d %s\n", info.Size(), pathToFile)
	}

	if *wantLines {
		file, err := os.Open(pathToFile)
		if err != nil {
			fmt.Printf("cannot open file, %s", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		count := 0

		for scanner.Scan() {
			count++
		}

		fmt.Printf("%d %s\n", count, pathToFile)
	}

	if *wantWords {
		file, err := os.Open(pathToFile)
		if err != nil {
			fmt.Printf("cannot open file, %s", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		count := 0

		for scanner.Scan() {
			for _, str := range strings.Fields(strings.TrimSpace(scanner.Text())) {
				if str != "" {
					count++
				}
			}
		}

		fmt.Printf("%d %s\n", count, pathToFile)
	}

	if *wantCharacters {
		file, err := os.Open(pathToFile)
		if err != nil {
			fmt.Printf("cannot open file, %s", err)
			os.Exit(1)
		}
		defer file.Close()

		count, err := countCharacters(file)
		if err != nil {
			fmt.Printf("cannot count characters, %s", err)
			os.Exit(1)
		}

		fmt.Printf("%d %s\n", count, pathToFile)
	}
}

func countCharacters(file *os.File) (int, error) {
	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		r := bufio.NewReader(strings.NewReader(scanner.Text()))
		for {
			_, _, err := r.ReadRune()
			if err != nil {
				if err == io.EOF {
					break
				}

				return -1, err
			}

			count++
		}
	}

	return count, nil
}
