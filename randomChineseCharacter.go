package sickocommon

import "math/rand"

func randInt(min, max int64) int64 {
	return min + rand.Int63n(max-min)
}

func RandomChineseCharacterN(length int) string {
	a := make([]rune, length)
	for i := range a {
		a[i] = rune(randInt(19968, 40869))
	}
	return string(a)
}
