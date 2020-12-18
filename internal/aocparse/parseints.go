package aocparse

import (
	"log"
	"strconv"
)

// ParseInts converts an
func ParseInts(str []string) (res []int) {
	for _, s := range str {
		i, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, int(i))
	}
	return
}

func SplitInts(str string, split string) []int {
	return nil
}
