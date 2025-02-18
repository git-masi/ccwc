package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	c := flag.Bool("c", false, "Return the number of bytes in a file")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("missing path to file")
		os.Exit(1)
	}

	p := os.Args[len(os.Args)-1]

	if *c {
		info, err := os.Stat(p)
		if err != nil {
			fmt.Printf("cannot get file info, %s", err)
			os.Exit(1)
		}

		fmt.Printf("%d %s\n", info.Size(), p)
	}
}
