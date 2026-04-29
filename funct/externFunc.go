package funct

import (
	"fmt"
	"os"
)

func ErrorMsg(err string) {
	fmt.Println(err)
	os.Exit(0)
}

func RemoveStartRoom(paths [][]string) [][]string {
	pathWithoutStart := [][]string{}
	for _, v := range paths {
		if len(v) > 0 {
			pathWithoutStart = append(pathWithoutStart, v[1:])
		}
	}
	return pathWithoutStart
}

func MakeGroups(paths [][]string) Group {
	groups := Group{}

	for i, path := range paths {
		if len(path) > 0 {
			groups[paths[i][0]] = append(groups[paths[i][0]], path)
		}
	}
	return groups
}

func IsRoomExist(name string, rooms []Room) bool {
	for _, v := range rooms {
		if v.Name == name {
			return true
		}
	}
	return false
}

func getRoom(name string, rooms []Room) Room {
	for _, v := range rooms {
		if v.Name == name {
			return v
		}
	}
	return Room{}
}

func GetStartRoom(rooms []Room) Room {
	for _, v := range rooms {
		if v.RoomType == "start" {
			return v
		}
	}
	return Room{}
}

func GetEndRoom(rooms []Room) Room {
	for _, v := range rooms {
		if v.RoomType == "end" {
			return v
		}
	}
	return Room{}
}
