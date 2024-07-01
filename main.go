package main

import (
	"fmt"
	"os"

	"github.com/eslamward/cat/cat"
)

func main() {

	if len(os.Args) < 2 {
		os.Exit(1)
	}
	for i := range os.Args[1:] {
		if os.Args[i] == ">" {
			if i == 0 {
				os.Create(os.Args[i])
				break

			} else if i == len(os.Args)-1 {
				fmt.Fprintln(os.Stderr, "error in trailing: > ")
				os.Exit(1)
			} else {
				cat.CopyFiles(os.Args[1:i], os.Args[i+1])
			}
		} else {
			cat.ReadFiles(os.Args[1:])

		}
	}

}
