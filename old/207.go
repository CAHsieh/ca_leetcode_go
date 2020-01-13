package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	ref := make([]int, numCourses)
	adj := make([][]bool, numCourses)
	for i := 0; i < numCourses; i++ {
		adj[i] = make([]bool, numCourses)
	}

	for _, prerequisite := range prerequisites {
		count := len(prerequisite)
		for i := 0; i < count-1; i++ {
			adj[prerequisite[i+1]][prerequisite[i]] = true
			ref[prerequisite[i]]++
		}
	}

	for i := 0; i < numCourses; i++ {
		s := 0
		for s < numCourses && ref[s] != 0 {
			s++
		}

		if s == numCourses {
			return false
		}

		ref[s] = -1

		for t := 0; t < numCourses; t++ {
			if adj[s][t] {
				ref[t]--
			}
		}
	}

	return true
}
