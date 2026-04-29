package funct

import (
	"strings"
)

func GetRelations(content string, rooms []Room) Links {
	relations := Links{}
	lines := strings.Split(content, "\n")

	for _, line := range lines {
		parts := strings.Split(line, "-")
		if len(parts) == 2 {
			if IsRoomExist(parts[0], rooms) && IsRoomExist(parts[1], rooms) {
				if parts[0] == parts[1] {
					ErrorMsg("ERROR: invalid data format,a room can't link to himself")
				}

				room1 := getRoom(parts[0], rooms)
				room2 := getRoom(parts[1], rooms)

				relations[room1.Name] = append(relations[room1.Name], room2.Name)

				relations[room2.Name] = append(relations[room2.Name],room1.Name)

			} else {
				ErrorMsg("ERROR: invalid data format,Link to unknown room")
			}
		}
	}
	return relations
}
