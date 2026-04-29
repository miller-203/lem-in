package funct

type Room struct {
	Name     string
	RoomType string
}

type AntDir struct {
	ant          string
	path         []string
	IsStartToEnd bool
}




type Links map[string][]string

type Group map[string][][]string
