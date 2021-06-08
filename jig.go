package sickocommon

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

func Jig(src string) (dst string) {
	faker := gofakeit.New(0)

	charcount := strings.Count(src, "#")
	numcount := strings.Count(src, "$")

	for i := 0; i < charcount; i++ {
		dst = strings.Replace(src, "#", faker.LetterN(1), 1)
	}
	for i := 0; i < numcount; i++ {
		dst = strings.Replace(dst, "$", fmt.Sprint(faker.Number(0, 9)), 1)
	}
	return dst
}
