package main

// make-pdf.go runs the pdf generation command within the notes file structure.
// I'm too lazy to write the command every time.

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <subject-name> <file-name>\nExample: %s linear-algebra eigenvectors\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}

	subjectName := os.Args[1]
	filename := getFilenameWithoutExtension(os.Args[2])

	// Change directory and check that folder and file exist.
	err := os.Chdir("../src/" + subjectName)
	if err != nil {
		fmt.Printf("ERROR: There is no folder named '%s'.\n", subjectName)
		return
	} else if _, err := os.Stat(filename + ".md"); os.IsNotExist(err) {
		fmt.Printf("ERROR: There is no file '%s' in folder '%s'.\n", filename, subjectName)
		return
	}

	// Execute command
	cmd := exec.Command("pandoc",
		"--latex-engine=xelatex",
		"-s", filename+".md",
		"-o", "../../pdf/"+subjectName+"/"+filename+".pdf")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Printf("There was an error in running the pandoc command. Is pandoc installed?\n")
		return
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
