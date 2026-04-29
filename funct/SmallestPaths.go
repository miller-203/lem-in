package funct

import (
	"sort"
)

func SmallestPaths(groups Group) []string {
	paths := extractShortestPaths(GroupTo3D(groups))
	return MaxCandidates(paths, groups)
}

func MaxCandidates(paths [][]string, groups Group) []string {
	max, bestPath := -1, []string{}
	for _, path := range paths {
		eligible := FindNonCrossingPaths(GetEligeables(path, groups))
		size := len(flattenPaths(eligible))
		if size > max {
			max = size
			bestPath = path
		}
	}
	return bestPath
}

func GetEligeables(targetPath []string, groups Group) [][]string {
	var eligible [][]string
	for _, paths := range groups {
		for _, path := range paths {
			if !sharesNodes(targetPath, path) {
				eligible = append(eligible, path)
			}
		}
	}
	return eligible
}

func sharesNodes(path1, path2 []string) bool {
	nodes := make(map[string]bool)
	for _, node := range path1[:len(path1)-1] {
		nodes[node] = true
	}
	for _, node := range path2[:len(path2)-1] {
		if nodes[node] {
			return true
		}
	}
	return false
}

func GroupTo3D(groups Group) [][][]string {
	var result [][][]string
	for _, v := range groups {
		result = append(result, v)
	}
	sort.Slice(result, func(i, j int) bool { return len(result[i]) < len(result[j]) })
	return result
}

func FindNonCrossingPaths(paths [][]string) [][]string {
	var nonCrossing [][]string
	for i, path1 := range paths {
		if !hasConflict(path1, paths, i) {
			nonCrossing = append(nonCrossing, path1)
		}
	}
	return nonCrossing
}

func hasConflict(path1 []string, paths [][]string, index int) bool {
	for j, path2 := range paths {
		if j != index && pathsIntersect(path1, path2) {
			return true
		}
	}
	return false
}

func pathsIntersect(path1, path2 []string) bool {
	for i := 0; i < len(path1)-1; i++ {
		for j := 0; j < len(path2)-1; j++ {
			if path1[i] == path2[j+1] && path1[i+1] == path2[j] {
				return true
			}
		}
	}
	return false
}

func extractShortestPaths(groups [][][]string) [][]string {
	var shortest [][]string
	for _, group := range groups {
		sort.Slice(group, func(i, j int) bool { return len(group[i]) < len(group[j]) })
		shortest = append(shortest, group[0])
	}
	sort.Slice(shortest, func(i, j int) bool { return len(shortest[i]) < len(shortest[j]) })
	return shortest
}

func flattenPaths(paths [][]string) []string {
	var flat []string
	for _, path := range paths {
		flat = append(flat, path...)
	}
	return flat
}
