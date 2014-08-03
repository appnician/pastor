package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gcmurphy/getpass"
)

var passwordLength = 12 - 5
var validChars = "abcdefghijklmnopqrstuvwxyz0123456789"

func main() {
	log.SetFlags(log.Lshortfile)

	basePhrase, _ := getpass.GetPassWithOptions("Enter base phrase: ", 0, getpass.DefaultMaxPass)
	basePhrase = strings.TrimSpace(basePhrase)

	digest := sha256.Sum256([]byte(basePhrase))

	fmt.Println(sum(digest[:]))

	for true {
		doorID, _ := getpass.GetPassWithOptions("Enter door id: ", 0, getpass.DefaultMaxPass)
		doorID = strings.TrimSpace(doorID)
		if doorID == "" {
			os.Exit(0)
		}

		keyData := sha256.Sum256([]byte(basePhrase + " - " + doorID))

		var pass []byte

		for _, e := range keyData[:passwordLength] {
			c := validChars[int(e)%len(validChars)]
			pass = append(pass, c)
		}
		fmt.Println(string(pass))
	}
}

func sum(a []byte) int {
	total := 0
	for _, e := range a {
		total += int(e)
	}
	return total
}
