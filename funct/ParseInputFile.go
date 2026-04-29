package funct

import (
	"os"
	"strconv"
	"strings"
)

func ParseInputFile(filename string) (Links, []Room, int, string) {

	content, err1 := os.ReadFile(filename)
	if err1 != nil {
		ErrorMsg(err1.Error())
	}
	MyFileContent := string(content)
	antNumbers := GetAntsNumber(MyFileContent)
	rooms := GetRooms(MyFileContent)
	relations := GetRelations(MyFileContent, rooms)

	return relations, rooms, antNumbers, MyFileContent
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func GetAntsNumber(content string) int {
	lines := strings.Split(content, "\n")
	antNum, err := strconv.Atoi(lines[0])
	if err != nil || antNum <= 0 {
		ErrorMsg("ERROR: invalid data format, Invalid Ants number")
	}
	return antNum
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func GetRooms(MyFile string) []Room {
	rooms := []Room{}
	startCount := strings.Count(MyFile, "##start")
	endCount := strings.Count(MyFile, "##end")
	lines := strings.Split(MyFile, "\n")

	if !strings.Contains(MyFile, "##start") || !strings.Contains(MyFile, "##end") {
		ErrorMsg("ERROR: invalid data format,No start or End room found")
	}

	if startCount > 1 || endCount > 1 {
		ErrorMsg("ERROR: invalid data format, More than one start/end command found")
	}

	for i, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) == 3 {
			room := checkRoom(parts, lines, i)
			if !Contains(room, rooms) {
				rooms = append(rooms, room)
			} else {
				ErrorMsg("ERROR: invalid data format,Duplicated Room")
			}
		}
	}
	return rooms
}

// +++++++++++++++++++++++++++++++++++++++++++++++++++++

func Contains(value Room, arr []Room) bool {
	for _, v := range arr {
		if value.Name == v.Name {
			return true
		}
	}
	return false
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

func checkRoom(parts []string, lines []string, i int) Room {
	name := parts[0]
	roomType := "normal"
	if string(name[0]) == "L" {
		ErrorMsg("ERROR: invalid data format,Name should never start with L")
	}

	if lines[i-1] == "##start" {
		roomType = "start"
	} else if lines[i-1] == "##end" {
		roomType = "end"
	}
	return Room{
		Name:     name,
		RoomType: roomType,
	}
}
