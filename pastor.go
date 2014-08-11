package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

var passwordLength = 12 - 5
var basePhrase string
var validChars = "abcdefghijklmnopqrstuvwxyz0123456789"
var sanityCheck string

func main() {
	// reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter base phrase: ")
	// basePhrase, _ := reader.ReadString('\n')
	basePhrase := "crap\n"
	basePhrase = strings.TrimSpace(basePhrase)
	fmt.Print(basePhrase)

	data := sha256.Sum256([]byte(basePhrase))

	fmt.Printf("%s", data)

	// for true {
	// 	fmt.Print("Enter door id: ")
	// 	doorID, _ := reader.ReadString('\n')
	// 	doorID = strings.TrimSpace(doorID)
	// 	if doorID == "" {
	// 		os.Exit(0)
	// 	}
	// 	fmt.Print(doorID)
	// }
}
