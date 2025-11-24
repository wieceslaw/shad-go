//go:build !solution

package hogwarts

func dfs(key string, graph map[string][]string, used map[string]int, result *[]string, i int) {
	used[key] = i
	for _, k := range graph[key] {
		if used[k] == i {
			panic("cycle")
		}
		if used[k] == 0 {
			dfs(k, graph, used, result, i)
		}
	}
	*result = append(*result, key)
}

func GetCourseList(prereqs map[string][]string) []string {
	// topological order
	result := make([]string, 0)
	used := make(map[string]int)
	i := 1
	for k := range prereqs {
		if used[k] == 0 {
			dfs(k, prereqs, used, &result, i)
		}
		i++
	}
	// slices.Reverse(result)
	return result
}
