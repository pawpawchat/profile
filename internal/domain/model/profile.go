package model

import "time"

type Profile struct {
	ID          int64     `db:"profile_id"`
	Username    string    `db:"username"`
	Description string    `db:"description"`
	NumFriends  int32     `db:"number_of_friends"`
	Online      bool      `db:"online"`
	LastSeen    time.Time `db:"last_seen"`
	CreatedAt   time.Time `db:"created_at"`
	Biography   Biography `db:"biography"`
}

type Biography struct {
	ID         int64      `db:"biography_id"`
	FirstName  string     `db:"first_name"`
	SecondName string     `db:"second_name"`
	Birthday   *time.Time `db:"birthday"`
	ProfileID  int64      `db:"profile_id"`
}

type Avatar struct {
	ID         int64     `db:"avatar_id"`
	ProfileID  int64     `db:"profile_id"`
	IsSelected bool      `db:"is_selected"`
	OrigURL    string    `db:"orig_url"`
	AddedAt    time.Time `db:"added_at"`
}
