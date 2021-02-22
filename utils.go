package huntingwords

import (
	"math/rand"
	"time"
)

func getInt(min, max int) int {
	rand.Seed(int64(time.Now().Nanosecond()))
	return rand.Intn(max-min+1) + min
}

func getString(sliceString []string) string {
	if len(sliceString) > 0 {
		indexRandom := getInt(0, len(sliceString)-1)
		return sliceString[indexRandom]
	}
	return ""
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
