package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func readFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func matchRegex(regexString string, content string) bool {
	var regex = regexp.MustCompile(regexString)
	return regex.MatchString(content)
}

func analyzeArgs(args []string) ([]string, bool, bool) {
	var isRecursive, isNameFile bool
	//Search the options first
	i := 0
	for i < len(args) {
		if !matchRegex("-.*", args[i]) {
			break
		}
		switch args[i] {
		case "-r":
			isRecursive = true
		case "-l":
			isNameFile = true
		}
		i++
	}
	fmt.Printf("The args are %v\n", args[i:])
	return args[i:], isRecursive, isNameFile
}

func main() {
	var content string
	var nameFiles []string
	var isMatch, isRecursive, isNameFile bool

	nameFiles, isRecursive, isNameFile = analyzeArgs(os.Args[1:])
	_ = isNameFile
	_ = isRecursive

	fmt.Printf("The file is %s\n", nameFiles[0])
	content = readFile(nameFiles[0])
	isMatch = matchRegex("[a-z]*Alarde[a-z]*", content)
	fmt.Printf("Concordo?: %t\n", isMatch)

}
