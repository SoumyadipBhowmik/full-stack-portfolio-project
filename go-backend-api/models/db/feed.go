package db

type Feed struct {
	Projects *[]Project `db:"projects"`
	Users    *[]User    `db:"users"`
	Posts    *[]Post    `db:"posts"`
}
