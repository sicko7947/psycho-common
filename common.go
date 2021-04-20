package sickocommon

import (
	"math/rand"
	"strings"
)

func getBase() string {
	slice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b", "c", "d", "e", "f"}
	index := rand.Intn(len(slice))
	return slice[index]
}

// GetUUID : Get Formatted Nike Type UUID as string
func GetUUID() string {
	data := "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"
	for i := 0; i <= 29; i++ {
		data = strings.Replace(data, "x", getBase(), 1)
	}
	data = strings.Replace(data, "y", getBase(), 1)
	return data
}
