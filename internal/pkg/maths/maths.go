package maths

func Abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func MinOf(is ...int) int {
	min := is[0]

	for _, i := range is {
		if min > i {
			min = i
		}
	}

	return min
}

func ZeroOrOne(i int) int {
	if i > 0 {
		return 1
	}

	if i < 0 {
		return -1
	}

	return 0
}
