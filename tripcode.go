package tripcode

import (
	"github.com/nyarlabo/go-crypt"
	"regexp"
	"strings"
)

func Tripcode(password string) (string) {
	salt := password + "H."[1:2]
	re := regexp.MustCompile("/[^.\\/0-9:;<=>?@A-Z\\[\\\\]\\^_`a-z]/")
	salt = re.ReplaceAllString(salt, ".")
	r := strings.NewReplacer(
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
	return crypt.Crypt(password, salt)[0:10]
}
