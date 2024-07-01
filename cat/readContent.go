package cat

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFiles(files []string) {

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			line := scanner.Text()
			fmt.Println(line)
		}
	}

}
