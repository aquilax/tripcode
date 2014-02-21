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
	"regexp"
	"strings"
)

// Tripcode function generates tripcode for the provided password
func Tripcode(password string) string {
	r := strings.NewReplacer(
		"&", "&amp;",
		"\"", "&quot;",
		"'", "&#39;",
		"<", "&lt;",
		">", "&gt;",
	)
	password = r.Replace(password)
	salt := (password + "H.")[1:3]
	re := regexp.MustCompile("/[^.\\/0-9:;<=>?@A-Z\\[\\\\]\\^_`a-z]/")
	salt = re.ReplaceAllString(salt, ".")
	r = strings.NewReplacer(
		":", "A",
		";", "B",
		"<", "C",
		"=", "F",
		">", "E",
		"?", "F",
		"@", "G",
		"[", "a",
		"\\", "b",
		"]", "c",
		"^", "d",
		"_", "e",
		"`", "f")
	salt = r.Replace(salt)
	code := crypt.Crypt(password, salt)
	l := len(code)
	return code[l-10 : l]
}
