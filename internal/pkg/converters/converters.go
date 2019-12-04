package converters

import (
	"log"
	"strconv"
)

func StringToInt(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Cannot convert: %s", s)
	}

	return r
}

func StringsToInts(ss ...string) []int {
	var is []int
	for _, s := range ss {
		r := StringToInt(s)
		is = append(is, r)
	}

	return is
}
