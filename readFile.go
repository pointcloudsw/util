/*
	Example courtesy of gemini
	https://www.google.com/search?q=building+a+file+reader+module+in+golang&sca_esv=1171e7f9e6150e63&source=hp&ei=Hc3_Z4D9GojgkPIPlLrp4AY&iflsig=ACkRmUkAAAAAZ__bLfIMWEpT7Fma7JS5NENvSO8jN-hR&ved=0ahUKEwjA_7m879yMAxUIMEQIHRRdGmwQ4dUDCBo&uact=5&oq=building+a+file+reader+module+in+golang&gs_lp=Egdnd3Mtd2l6IididWlsZGluZyBhIGZpbGUgcmVhZGVyIG1vZHVsZSBpbiBnb2xhbmcyBRAhGKABMgUQIRigATIFECEYoAEyBRAhGKABMgUQIRifBTIFECEYnwUyBRAhGJ8FMgUQIRifBTIFECEYnwUyBRAhGJ8FSJg6UABYpzlwAHgAkAEBmAHWAqABiSmqAQkxNi4yOC4xLjG4AQPIAQD4AQGYAi2gArInwgIOEC4YgAQYsQMY0QMYxwHCAgUQABiABMICCxAAGIAEGLEDGIMBwgIIEAAYgAQYsQPCAg4QLhiABBixAxiDARiKBcICBRAuGIAEwgIOEAAYgAQYsQMYgwEYigXCAggQLhiABBixA8ICERAuGIAEGLEDGIMBGMcBGK8BwgIREC4YgAQYsQMY0QMYgwEYxwHCAgsQLhiABBjRAxjHAcICFBAuGIAEGLEDGIMBGMcBGI4FGK8BwgILEC4YgAQYxwEYrwHCAg4QLhiABBjHARiOBRivAcICCxAAGIAEGLEDGIoFwgIGEAAYFhgewgIFEAAY7wXCAggQABiABBiiBMICBRAhGKsCmAMAkgcHMTQuMjkuMqAHh9QCsgcHMTQuMjkuMrgHsic&sclient=gws-wiz


The module itself
-----------------
*/

package util

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFile reads the entire content of a file and returns it as a string.
func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}
	return string(content), nil
}

// ReadFileLines reads a file line by line and returns a slice of strings.
func ReadFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return lines, nil
}

/*
-------------------

Using the module
----------------
		package main
		import (
			"fmt"
			"log"
			"your_module_name/fileutil"
		)

		func main() {
		// Read entire file
			content, err := fileutil.ReadFile("example.txt")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("File Content:\n", content)

			// Read file line by line
			lines, err := fileutil.ReadFileLines("example.txt")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("\nFile Lines:")
			for _, line := range lines {
				fmt.Println(line)
			}
		}

*/
/*
package util

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func readFile(filename string) error {

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
		}

	} else {
		log.Fatal("Unable to access input file: ", filename)
	}
}
*/
