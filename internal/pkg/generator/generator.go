package generator

type Generator func() (int, bool)

func CreateIntGenerator(inputs ...int) Generator {
	point := 0

	gen := func() (int, bool) {
		if point >= len(inputs) {
			return 0, false
		}

		r := inputs[point]
		point += 1
		return r, true
	}

	return gen
}
