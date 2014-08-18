package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <subject-name> <filename>\nExample: %s linear-algebra\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}

	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <subject-name> <file-name>\nExample: %s linear-algebra eigenvectors\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}

	subjectName := os.Args[1]
	filename := getFilenameWithoutExtension(os.Args[2])

	if _, err := os.Stat("../src/" + subjectName); os.IsNotExist(err) {
		os.Mkdir("../src/"+subjectName, 0700)
	} else if _, err := os.Stat("../src/" + subjectName); os.IsNotExist(err) {
		os.Mkdir("../pdf/"+subjectName, 0700)
	}

}

func getFilenameWithoutExtension(rawFilename string) string {
	for i := 0; i < len(rawFilename); i++ {
		if string(rawFilename[i]) == "." {
			return rawFilename[:i]
		}
	}
	return rawFilename
}
