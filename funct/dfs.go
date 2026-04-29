package funct


func FindAllPaths(num int, adjList Links, start, end string) [][]string {
	visited := make(map[string]bool)
	path := []string{start}
	paths := [][]string{}

	dfs(num+1, adjList, start, end, visited, path, &paths)
	if len(paths) == 0 {
		ErrorMsg("ERROR: invalid data format")
	}
	return paths
}

func dfs(num int, adjList Links, current, end string, visited map[string]bool, path []string, paths *[][]string) {
	if num == 15 {
		return
	}
	visited[current] = true

	if current == end {
		// Append a copy of the current path to the paths slice
		*paths = append(*paths, append([]string(nil), path...))
	} else {
		edges := adjList[current]
		for _, edge := range edges {
			if !visited[edge] {
				// Explore the unvisited neighbor

				dfs(num+1, adjList, edge, end, visited, append(path, edge), paths)
			}
		}
	}

	visited[current] = false // Mark the current node as unvisited for other paths
}
