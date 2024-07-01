package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/eslamward/cat/cat"
)

func main() {

	lineNumber := flag.Bool("n", false, "number of line in one file")
	help := flag.Bool("help", false, "Help for")
	version := flag.Bool("v", false, "number of line in one file")

	flag.Parse()

	if len(os.Args) < 2 {
		os.Exit(1)
	}

	if *lineNumber {
		cat.ReadFilesWithLinesNumber(os.Args[2:])
		return
	}
	if *help {
		fmt.Printf(`
		-you can display content of file by: "cat file.ext"
		-you can display more than file at once by : "cat file.ext file2.ext"
		-to copy file into other file use this cmd : "cat sourceFile.ext > destFile.ext"
		-you can copy more then file into file by: "cat sorce1.ext source2.ext > dest.ext"
		-user -n to display line number: "cat -n file.ext"
		-to create file user : "cat > file.ext"
		-for version user: "cat -v"
		-for help use: "cat -help"

				`)
		return
	}
	if *version {
		fmt.Println("Cat version 1.0.0")
		return
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
