package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	AUTHOR = "Steven Schmatz, University of Michigan: College of Engineering"
	EMAIL  = "stevenschmatz@gmail.com"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <subject-name> <file-name>\nExample: %s linear-algebra eigenvectors\n", os.Args[0], os.Args[0])
		os.Exit(1)
	}

	subjectName := os.Args[1]
	filename := getFilenameWithoutExtension(os.Args[2])

	if _, err := os.Stat("../src/" + subjectName); os.IsNotExist(err) {
		os.Mkdir("../src/"+subjectName, 0700)
	} else if _, err := os.Stat("../pdf/" + subjectName); os.IsNotExist(err) {
		os.Mkdir("../pdf/"+subjectName, 0700)
	}

	err := os.Chdir("../src/" + subjectName)
	if err != nil {
		fmt.Println("There was an error opening the src folder.")
		return
	}

	err = ioutil.WriteFile(filename+".md", []byte(getLatexTitlePage(filename, subjectName)), 0700)
	if err != nil {
		fmt.Println("There was an error in writing the file.")
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

func getLatexTitlePage(filename, subjectName string) string {
	latexTitle := "\\title{" + subjectName + ": " + filename + "}\n"

	month := time.Now().Month().String()
	day := strconv.Itoa(time.Now().Day())
	year := strconv.Itoa(time.Now().Year())

	latexDate := "\\date{" + month + " " + day + ", " + year + "}\n"
	latexAuthor := "\\author{" + AUTHOR + "}\n\n"

	latexEmail := "\\begin{center} " + EMAIL + " \\end{center}"

	return latexTitle + latexDate + latexAuthor + "\\maketitle\n" + latexEmail + " \\pagebreak\n\n<!---Content goes here-->"
}
