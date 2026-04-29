package main

import (
	"fmt"
	"os"

	"funct/funct"
)

func main() {
	args := os.Args

	// Check if the file name is provided as an argument
	if len(args) <= 1 {
		funct.ErrorMsg("You Forget File Name")
	}

	// Parse the input file and extract relations, rooms, ants, and file content
	relations, rooms, ants, MyFileContent := funct.ParseInputFile("examples/" + args[1])
	
	// Get the names of the start and end rooms
	start := funct.GetStartRoom(rooms).Name
	end := funct.GetEndRoom(rooms).Name

	allPaths := funct.FindAllPaths(0, relations, start, end) // Find all possible paths from start to end
	
	rmStart := funct.RemoveStartRoom(allPaths) // Remove the start room from all paths
	
	
	makeGroups := funct.MakeGroups(rmStart) // Group paths by their starting points

	thePromised := funct.SmallestPaths(makeGroups) // Find the shortest paths in the groups

	
	elected := [][]string{thePromised} // Initialize elected paths with the shortest path
	
	eligeables := funct.GetEligeables(thePromised, makeGroups) // Get paths eligible for further selection
	
	elected = funct.Elector(eligeables, thePromised, elected, makeGroups) // Elect the best group of paths based on eligibility
	
	fmt.Println(MyFileContent + "\n") // Print the original file content

	funct.SendAntsByWaves(ants, elected) // Send ants through the selected paths in waves
}
