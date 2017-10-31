package main

//Label used to record leader
type Label struct {
	root  *Label
	label int
}

func findCircleNum(M [][]int) int {

	labelSlice := make([]Label, len(M))
	for i := range labelSlice {
		labelSlice[i].root = &labelSlice[i]
		labelSlice[i].label = i
	}

	for i := range M {
		for j := i + 1; j < len(M[i]); j++ {
			if 1 == M[i][j] {
				join(&labelSlice[i], &labelSlice[j])
			}
		}
	}

	m := make(map[int]int)
	count := 0

	for i := range labelSlice {
		label := getRoot(&labelSlice[i]).label
		_, ok := m[label]
		if !ok {
			m[label] = 1
			count++
		}
	}

	return count
}

func getRoot(a *Label) *Label {
	if a.root != a {
		a.root = getRoot(a.root)
	}
	return a.root
}

func join(a *Label, b *Label) {
	getRoot(b).root = getRoot(a).root
}
