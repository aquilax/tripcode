/*
Package tripcode generates 4chan comapitble tripcodes for use mainly in anonymous forums.
There are different modifications of the tripcode algorythm. This one is based on code
from http://avimedia.livejournal.com/1583.html

Example usage:

  package main

  import "github.com/aquilax/tripcode"

  func main() {
	  print(tripcode.Tripcode("password")
  }
*/
package tripcode

import (
	"github.com/nyarlabo/go-crypt"
	"github.com/qiniu/iconv"
	"strings"
)

const SALT_TABLE = ".............................................../0123456789ABCDEFGABCDEFGHIJKLMNOPQRSTUVWXYZabcdefabcdefghijklmnopqrstuvwxyz....................................................................................................................................."

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
		"'", "&#39;",
		"<", "&lt;",
		">", "&gt;",
	)
	return r.Replace(text)
}

func generateSalt(password string) string {
	var salt [2]rune
	password = substr(password+"H.", 1, 2)
	for i, r := range password {
		salt[i] = rune(SALT_TABLE[r%256])
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

// Tripcode function generates tripcode for the provided password
func Tripcode(password string) string {
	password = sjisToUtf8(password)
	password = htmlEscape(password)
	if password == "" {
		return password
	}
	if len(password) > 8 {
		password = substr(password, 0, 8)
	}
	salt := generateSalt(password)
	code := crypt.Crypt(password, salt)
	l := len(code)
	return code[l-10 : l]
}
