package generator

type Generator func() (int, bool)

func CreateIntGenerator(inputs ...int) Generator {
	point := 0
	l := len(inputs)

	gen := func() (int, bool) {
		if point >= l {
			return 0, false
		}

		r := inputs[point]
		point += 1
		return r, true
	}

	return gen
}
