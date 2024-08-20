package model

import "time"

type Profile struct {
	Id          uint64
	Username    string
	Description string
	NumFriends  string
}

type Biography struct {
	Firstname  string
	Secondname string
	Birthday   time.Time
}
