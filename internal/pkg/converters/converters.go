package converters

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func StringToInt(s string) int {
	r, err := strconv.Atoi(s)

	if err != nil {
		log.Fatalf("Cannot convert: %s, %v", s, err)
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

func IntToDigits(i int) []int {
	remainder := i
	var digits []int

	if i < 10 {
		return []int{i}
	}

	for true {
		d := remainder % 10
		digits = append(digits, d)

		if remainder < 10 {
			break
		}
		remainder = remainder / 10
	}

	return Reverse(digits)
}

func StringToDigits(s string) []int {
	var digits []int

	for _, ss := range s {
		ds := IntToDigits(StringToInt(string(ss)))
		digits = append(digits, ds...)
	}
	return digits
}

func Reverse(is []int) []int {
	var n []int
	for i := len(is) - 1; i > -1; i-- {
		n = append(n, is[i])
	}

	return n
}

func IntsToString(delim string, is ...int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(is), " ", delim, -1), "[]")
}
