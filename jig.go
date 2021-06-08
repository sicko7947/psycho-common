package sickocommon

import (
	"fmt"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
)

func Jig(src string) string {
	faker := gofakeit.New(0)

	charcount := strings.Count(src, "#")
	numcount := strings.Count(src, "$")

	for i := 0; i < charcount; i++ {
		src = strings.Replace(src, "#", faker.LetterN(1), 1)
	}
	for i := 0; i < numcount; i++ {
		src = strings.Replace(src, "$", fmt.Sprint(faker.Number(0, 9)), 1)
	}
	return src
}
