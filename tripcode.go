/*
Package tripcode generates 4chan comapitble tripcodes for use mainly in anonymous forums.
There are different modifications of the tripcode algorythm. This one is based on code
from http://avimedia.livejournal.com/1583.html

Example usage:

  package main

  import "github.com/ComSecNinja/tripcode"

  func main() {
	  print(tripcode.Tripcode("password")
  }
*/
package tripcode

import (
	"encoding/base64"
	"crypto/sha1"
	"strings"

	"github.com/nyarlabo/go-crypt"
	"github.com/qiniu/iconv"
)

const saltTable = "" +
	"................................" +
	".............../0123456789ABCDEF" +
	"GABCDEFGHIJKLMNOPQRSTUVWXYZabcde" +
	"fabcdefghijklmnopqrstuvwxyz....." +
	"................................" +
	"................................" +
	"................................" +
	"................................"

func sjisToUtf8(text string) string {
	cd, err := iconv.Open("SJIS", "utf-8")
	defer cd.Close()
	if err != nil {
		panic("iconv.Open failed!")
	}
	return cd.ConvString(text)
}

func htmlEscape(text string) string {
	r := strings.NewReplacer(
		"&", "&amp;",
		"\"", "&quot;",
		"<", "&lt;",
		">", "&gt;",
	)
	return r.Replace(text)
}

func generateSalt(password string) string {
	var salt [2]rune
	password = substr(password+"H.", 1, 2)
	for i, r := range password {
		salt[i] = rune(saltTable[r%256])
	}
	return string(salt[:])
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func prepare(password string) string {
	password = sjisToUtf8(password)
	password = htmlEscape(password)
	if len(password) > 8 {
		password = substr(password, 0, 8)
	}
}

// Tripcode generates a tripcode for the provided password.
func Tripcode(password string) string {
	password = prepare(password)
	if password == "" {
		return password
	}
	salt := generateSalt(password)
	code := crypt.Crypt(password, salt)
	l := len(code)
	return code[l-10 : l]
}

// SecureTripcode generates a secure tripcode based
// on the provided password and a secure salt combination.
func SecureTripcode(password string, secureSalt string) string {
	// Prepare the password (encoding conversion etc.).
	password = prepare(password)
	// Append password+salt and calculate sha1 hash.
	hash := sha1.New().Sum(append([]byte(password), []byte(secureSalt)...))
	// Encode the hash to base64 string, forming our salt for this tripcode.
	salt := base64.NewEncoding(base64.StdEncoding).EncodeToString(hash)
	// Crypt the password using "_..A." + 4 of the first bytes of the salt.
	code := crypt.Crypt(password, "_..A." + salt[:4])
	l := len(code)
	return code[l-10 : l]
}
