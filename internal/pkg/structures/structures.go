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

func ArrayToMatrices(arr []int, width, height int) [][][]int {
	numMatrices := len(arr) / width / height
	var layers [][][]int
	pointer := 0

	for i := 0; i < numMatrices; i++ {
		var matrix [][]int

		for j := 0; j < height; j++ {
			matrix = append(matrix, arr[pointer:pointer+width])
			pointer += width
		}

		layers = append(layers, matrix)
	}

	return layers
}
