package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	countBytes := flag.Bool("c", false, "Return the number of bytes in a file")
	countLines := flag.Bool("l", false, "Return the number of lines in a file")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("missing path to file")
		os.Exit(1)
	}

	pathToFile := os.Args[len(os.Args)-1]

	if *countBytes {
		info, err := os.Stat(pathToFile)
		if err != nil {
			fmt.Printf("cannot get file info, %s", err)
			os.Exit(1)
		}

		fmt.Printf("%d %s\n", info.Size(), pathToFile)
	}

	if *countLines {
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
}
