package funct

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)


func SendAntsByWaves(antsNum int, paths [][]string) {
	ants := CreateAnts(antsNum)
	paths = SortPaths(paths)
	paths = RemoveDuplicates(paths)
	antsEachPath := make([]int, len(paths))
	antsWithPaths := AntsDirections(ants, paths, antsEachPath)
	SendWaves(antsWithPaths)
}

func CreateAnts(antsNum int) []string {
	ants := make([]string, antsNum)
	for i := 1; i <= antsNum; i++ {
		ants[i-1] = "L" + strconv.Itoa(i)
	}
	return ants
}

func AntsDirections(ants []string, paths [][]string, antsEachPath []int) []AntDir {
	var antsWithPaths []AntDir
	curPath := 0
	for _, ant := range ants {
		currentAnt := AntDir{
			ant:          ant,
			path:         paths[curPath],
			IsStartToEnd: false,
		}
		if len(currentAnt.path) == 1 {
			currentAnt.IsStartToEnd = true
		}

		antsWithPaths = append(antsWithPaths, currentAnt)
		antsEachPath[curPath]++

		// Logic to balance ants across paths
		if curPath < len(paths)-1 && antsEachPath[curPath]+len(paths[curPath]) > antsEachPath[curPath+1]+len(paths[curPath+1]) {
			curPath++
		} else if curPath == len(paths)-1 {
			curPath = 0
		}
	}
	return antsWithPaths
}

func SendWaves(antsWithPaths []AntDir) {
	endRoom := antsWithPaths[0].path[len(antsWithPaths[0].path)-1]
	for len(antsWithPaths) > 0 {
		var usedStartToEnd bool = false
		var usedRooms []string
		var nextWave []AntDir
		for _, antDir := range antsWithPaths {
			if len(antDir.path) == 0 {
				continue
			}

			if (usedStartToEnd && antDir.IsStartToEnd) || (slices.Contains(usedRooms, antDir.path[0]) && antDir.path[0] != endRoom) {
				nextWave = append(nextWave, antDir)
				continue
			}

			if len(antDir.path) == 1 {
				usedStartToEnd = true
			}
			fmt.Printf("%s-%s ", antDir.ant, antDir.path[0])
			usedRooms = append(usedRooms, antDir.path[0])
			antDir.path = antDir.path[1:]

			if len(antDir.path) > 0 {
				nextWave = append(nextWave, antDir)
			}
		}
		fmt.Println()
		antsWithPaths = nextWave
	}
}

func SortPaths(paths [][]string) [][]string {
	for i := 0; i < len(paths); i++ {
		for j := 0; j < len(paths)-1; j++ {
			if len(paths[j]) > len(paths[j+1]) {
				paths[j], paths[j+1] = paths[j+1], paths[j]
			}
		}
	}
	return paths
}

func ContSl(orig [][]string, ser []string) bool {
	for _, item := range orig {
		if strings.Join(item, ",") == strings.Join(ser, ",") {
			return true
		}
	}
	return false
}

func RemoveDuplicates(orig [][]string) [][]string {
	var newSl [][]string
	for _, path := range orig {
		if !ContSl(newSl, path) {
			newSl = append(newSl, path)
		}
	}
	return newSl
}
