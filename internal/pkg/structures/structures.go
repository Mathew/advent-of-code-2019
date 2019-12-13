package structures

func GetFirstIntersection(arr, arr2 []string) (string, bool) {
	m := map[string]int{}

	for x, a := range arr {
		m[a] += 1
		m[arr2[x]] += 1

		if m[a] >= 2 {
			return a, true
		}
		if m[arr2[x]] >= 2 {
			return arr2[x], true
		}
	}

	return "", false
}
