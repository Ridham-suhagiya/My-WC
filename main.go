package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {

	var scanner *bufio.Scanner
	var pathToTheFile string

	fileInfo, err := os.Stdin.Stat()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if fileInfo.Mode()&os.ModeNamedPipe != 0 {
		scanner = bufio.NewScanner(os.Stdin)

	} else {
		listOfArguments := os.Args[1:]
		if len(listOfArguments) < 1 {
			fmt.Printf("error: no file path given\n")
			return
		}
		var pathToTheFile string = listOfArguments[0]

		isValid := validateFile(pathToTheFile)
		if !isValid {
			return
		}
		file, err := os.Open(pathToTheFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		scanner = bufio.NewScanner(file)
		defer file.Close()
	}
	arguments := os.Args
	var operator string
	if len(arguments) > 1 {
		operator = arguments[len(arguments)-1]
		operatorMime := []string{"-c", "-l", "-w", "-m"}
		isValid := arrayIncludesElement(operator, operatorMime)
		if !isValid {
			operator = ""
		}
	}
	detialsArray := getFileDetails(scanner)

	if operator == "-c" || operator == "" {
		fmt.Printf("%s ", detialsArray[0])
	}
	if operator == "-l" || operator == "" {
		fmt.Printf("%s ", detialsArray[1])
	}

	if operator == "-w" || operator == "" {
		fmt.Printf("%s ", detialsArray[2])
	}
	if operator == "-m" || operator == "" {
		fmt.Printf("%s ", detialsArray[3])
	}
	if len(pathToTheFile) > 0 {
		fmt.Printf("%s \n", pathToTheFile)
	}
	fmt.Printf("\n")

}

func getFileDetails(scanner *bufio.Scanner) [4]string {

	var listArray [4]string
	var totalBytes int64
	var totalNumberOfWords int
	var totalNumberOfCharacters int
	var totalLines int
	for scanner.Scan() {

		bytesInLine := int64(len(scanner.Bytes()))
		totalBytes += (bytesInLine + 1)

		line := scanner.Text()
		words := strings.Fields(line)

		totalNumberOfWords += (len(words))

		totalNumberOfCharacters += (len(line) + 1)

		totalLines++
	}
	listArray[0] = strconv.FormatInt(totalBytes, 10)
	listArray[1] = strconv.Itoa(totalLines)
	listArray[2] = strconv.Itoa(totalNumberOfWords)
	listArray[3] = strconv.Itoa(totalNumberOfCharacters)
	return listArray
}

func validateFile(fileName string) bool {
	validExtensions := []string{".txt", ".csv", ".pdf"}
	ext := filepath.Ext(fileName)
	ext = strings.ToLower(ext)
	isValid := arrayIncludesElement(ext, validExtensions)
	if !isValid {
		fmt.Printf("error: File type should be from %v \n", validExtensions)
		return isValid
	}
	return isValid

}

func arrayIncludesElement(element string, array []string) bool {
	for _, arrayElement := range array {
		if arrayElement == element {
			return true
		}
	}
	return false
}
