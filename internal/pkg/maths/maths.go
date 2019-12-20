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

func Permutations(arr ...int) [][]int {
	var helper func([]int, int)
	var res [][]int

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
