package db

import "time"

type Feed struct {
	Projects *[]Project `db:"projects"`
	Users    *[]User    `db:"users"`
	Time     time.Time  `db:"time"`
}
