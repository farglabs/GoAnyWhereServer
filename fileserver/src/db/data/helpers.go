package data

import (
	"database/sql"
	"time"
)

type user struct {
	username      string
	department_id int
	created       time.Time
}

func AddUser(u *user, d *sql.DB) (int64, error) {
	res, err := d.Exec(u.username, u.department_id, u.created)

	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	return id, err
}
