package cat

import (
	"bufio"
	"fmt"
	"os"
)

func CopyFiles(files []string, fileWritedToIt string) {
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			err := os.WriteFile(fileWritedToIt, []byte(line), 0777)
			fmt.Fprintln(os.Stderr, err.Error())

			os.Exit(2)
		}
	}
}
