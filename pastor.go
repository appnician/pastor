package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"log"
	"strings"

	"code.google.com/p/go.crypto/pbkdf2"
	"github.com/gcmurphy/getpass"
)

var passwordLength = 16
var validChars = "abcdefghijklmnopqrstuvwxyz0123456789.~!@#$%^&*()_+"

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
			break
		}

		key := pbkdf2.Key([]byte(doorID), []byte(basePhrase), 1000, 32, sha1.New)

		keyData := sha256.Sum256([]byte(key))

		var pass []byte

		for _, e := range keyData[:passwordLength] {
			c := validChars[int(e)%len(validChars)]
			pass = append(pass, c)
		}
		fmt.Println(string(pass))
	}
	clearScreen()
}

func sum(a []byte) int {
	total := 0
	for _, e := range a {
		total += int(e)
	}
	return total
}

func clearScreen() {
	fmt.Print("\033[2J")
}
