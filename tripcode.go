/*
Package tripcode generates 4chan compatible tripcodes mainly for anonymous forums.
There are different implementations of the tripcode algorithm. This one is based on code
from http://avimedia.livejournal.com/1583.html

Example usage:

  package main

  import "github.com/ComSecNinja/tripcode"

  func main() {
	  print(tripcode.Tripcode("password")
	  print(tripcode.SecureTripcode("password", "secure salt"))
  }
*/
package tripcode

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"strings"

	"github.com/nyarlabo/go-crypt"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
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

func convert(text string) string {
	var s bytes.Buffer
	transform.NewWriter(&s, japanese.ShiftJIS.NewEncoder()).Write([]byte(text))
	return s.String()
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
	password = convert(password)
	password = htmlEscape(password)
	if len(password) > 8 {
		password = password[:8]
	}
	return password
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
	password = prepare(password)
	// Append password+salt and calculate sha1 hash.
	hash := sha1.New().Sum(append([]byte(password), []byte(secureSalt)...))
	salt := base64.StdEncoding.EncodeToString(hash)
	code := crypt.Crypt(password, "_..A."+salt[:4])
	l := len(code)
	return code[l-10 : l]
}
