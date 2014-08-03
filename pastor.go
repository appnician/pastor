package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"code.google.com/p/go.crypto/pbkdf2"
	"github.com/gcmurphy/getpass"
)

var passwordLength = 16
var validChars = "abcdefghijklmnopqrstuvwxyz0123456789.~!@#$%^&*()_+"

func main() {
	basePhrase, _ := getpass.GetPassWithOptions("Enter base phrase: ", 0, getpass.DefaultMaxPass)

	digest := sha256.Sum256([]byte(basePhrase))

	fmt.Println(sum(digest[:]))

	for true {
		doorID, _ := getpass.GetPassWithOptions("Enter door id: ", 0, getpass.DefaultMaxPass)
		if doorID == "" {
			break
		}

		var pass []byte

		keyData := crypt(doorID, basePhrase, 1000)
		rawHash := sha256.Sum256(keyData)

		for _, e := range rawHash[:passwordLength] {
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

// A stripped down version of crypt using pbkdf2
// Copies from:
// https://github.com/dlitz/python-pbkdf2/blob/master/pbkdf2.py#L230
// meant to be used in limited cases
func crypt(word string, salt string, iterations int) []byte {
	salt = fmt.Sprintf("$p5k2$%x$%s", iterations, salt)
	length := 24
	key := pbkdf2.Key([]byte(word), []byte(salt), iterations, length, sha1.New)
	hash := base64.StdEncoding.EncodeToString(key)
	out := fmt.Sprintf("%s$%s", salt, hash)
	return []byte(out)
}
