package sickocommon

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"strings"
)

func Jig(src string) (dst string) {
	faker := gofakeit.New(0)
	dst = strings.ReplaceAll(src, `#`, faker.LetterN(1))
	dst = strings.ReplaceAll(dst, `$`, fmt.Sprint(faker.Number(0, 9)))
	return dst
}
