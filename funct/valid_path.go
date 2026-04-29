package funct

import (
	"math"
)

func Elector(eligeables [][]string, thePromised []string, elected [][]string, groups Group) [][]string {
	found := 0
	min := math.MaxInt32

	allPossibilities := GenerateAllPossibilities(thePromised, eligeables)
	for _, v := range allPossibilities {
		if !HasCommonElements2(v) {
			found++
			flat := Flat2DArray(v)
			if len(flat) < min {
				min = len(flat)
				elected = v
			}
		}
	}

	if found == 0 {
		min = math.MaxInt32
		choosen := [][]string{}
		for key, subPath := range groups {
			if key != thePromised[0] {
				flat := Flat2DArray(subPath)
				if len(flat) > 0 && len(flat) < min {
					for _, sub := range subPath {
						if !HasCommonElements(sub, thePromised) {
							choosen = [][]string{sub}
							break
						}
					}
				}
			}
		}

		elected = append(elected, choosen...)

	}
	return elected
}

func GenerateAllPossibilities(thePromised []string, eligeables [][]string) [][][]string {
	groups := MakeGroups(eligeables)
	allPossibilities := [][][]string{}

	for key := range groups {
		sources, targets := getSourcesAndTargets(groups, key)
		allPossibilities = append(allPossibilities, GenerateCombinations(thePromised, targets, sources)...)
	}

	return allPossibilities
}

func HasCommonElements2(arr [][]string) bool {
	set := make(map[string]bool)

	for _, subArr := range arr {
		if len(subArr) > 0 {
			for _, element := range subArr[:len(subArr)-1] {
				if set[element] {
					return true
				}
				set[element] = true
			}
		}
	}

	return false
}

func Flat2DArray(v [][]string) []string {
	flat := []string{}
	for _, val := range v {
		flat = append(flat, val...)
	}
	return flat
}

func getSourcesAndTargets(groups Group, key string) ([][]string, [][]string) {
	sources := groups[key]
	targets := [][]string{}

	for k, v := range groups {
		if k != key {
			targets = append(targets, v...)
		}
	}
	return sources, targets
}

func GenerateCombinations(thePromised []string, sources [][]string, targets [][]string) [][][]string {
	combinations := [][][]string{}

	for range thePromised {
		for _, target := range targets {
			for _, source := range sources {
				combination := [][]string{
					thePromised,
					target,
					source,
				}
				combinations = append(combinations, combination)
			}
		}
	}
	return combinations
}

func HasCommonElements(arr1, arr2 []string) bool {
	set := make(map[string]bool)

	for _, element := range arr1[:len(arr1)-1] {
		set[element] = true
	}
	for _, element := range arr2[:len(arr2)-1] {
		if set[element] {
			return true
		}
	}
	return false
}
