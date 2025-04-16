package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// loadEnv sets environment variables based on content input
func LoadEnv(filename string) error {

	var inputString []string

	if filename != "" {
		f, e := os.Open(filename)
		if e != nil {
			log.Fatal("Error opening input file:", e)
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			inputString = strings.Split(scanner.Text(), "=")
			os.Setenv(inputString[0], inputString[1])
		}

	} else {
		log.Fatal("Unable to access input file: ", filename)

	}
	return nil
}
