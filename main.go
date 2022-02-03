package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stringValue := CleanedUserInput("Please Enter a string to reverse")
	reveresedString := ReverseString(stringValue)
	fmt.Println("Reveresed Stirng: " + reveresedString)
}

func ReverseString(stringValue string) string {
	var reveresedString string
	for i := len(stringValue) - 1; i >= 0; i-- {
		reveresedString = reveresedString + string(stringValue[i])
	}

	return reveresedString
}

func CleanedUserInput(consoleMessage string) string {
	fmt.Println(consoleMessage)
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')
	value = strings.Trim(value, "\n")
	return value
}
