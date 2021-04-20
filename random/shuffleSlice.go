package random

import "math/rand"

func ShuffleSlice(s interface{}) {
	ss := s.([]interface{})
	rand.Shuffle(len(ss), func(i, j int) {
		ss[i], ss[j] = ss[j], ss[i]
	})
}
