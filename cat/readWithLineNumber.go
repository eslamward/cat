package cat

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFilesWithLinesNumber(files []string) {

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			continue
		}

		scanner := bufio.NewScanner(f)
		i := 0
		for scanner.Scan() {
			i++
			line := scanner.Text()
			numberOfline := fmt.Sprintf("%d ", i) + line
			fmt.Println(numberOfline)
		}
	}

}
