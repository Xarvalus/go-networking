package utils

import (
	"bufio"
	"os"
)

func InputFromConsole() string {
	scanner := bufio.NewScanner(os.Stdin)
	LogFatalError(scanner.Err())

	for scanner.Scan() {
		return scanner.Text()
	}

	panic("Scanner did not return input")
}
