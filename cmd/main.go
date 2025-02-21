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

		scanner := bufio.NewScanner(file)
		count := 0

		for scanner.Scan() {
			count += utf8.RuneCount(scanner.Bytes())
		}

		fmt.Printf("%d %s\n", count, pathToFile)
	}
}
