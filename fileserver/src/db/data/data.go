package data

import (
	"database/sql"
	"fmt"
	"time"
)

func Seed(db *sql.DB) {
	// insert
	stmt, err := db.Prepare("INSERT INTO files(username, departname, created) values(?,?,?)")
	CheckDBErr(err)

	type user struct {
		username string
		department_id int
		created time.Time
	}
	res, err := stmt.Exec("romanjordan", "1", "2012-12-09")
	CheckDBErr(err)

	id, err := res.LastInsertId()
	CheckDBErr(err)

	fmt.Printf("Last Insert ID %d\n",id)
	// update
	stmt, err = db.Prepare("update files set username=? where uid=?")
	CheckDBErr(err)

	res, err = stmt.Exec("Roman", id)
	CheckDBErr(err)

	affect, err := res.RowsAffected()
	CheckDBErr(err)

	fmt.Printf("Number of rows `Updated` %d\n", affect)
	
	// query
	rows, err := db.Query("SELECT * FROM files")
	CheckDBErr(err)
	var uid int
	var username string
	var department string
	var created time.Time
	
	for rows.Next() {
		err = rows.Scan(&uid, &username, &department, &created)
		CheckDBErr(err)
		fmt.Printf("User ID:%d\n",uid)
		fmt.Printf("Username: %s\n", username)
	}
	
	rows.Close() //good habit to close
	
	// delete
	stmt, err = db.Prepare("delete from files where uid=?")
	CheckDBErr(err)
	
	res, err = stmt.Exec(id)
	CheckDBErr(err)
	
	affect, err = res.RowsAffected()
	CheckDBErr(err)
	
	fmt.Printf("Number of rows `Deleted` %d\n", affect)

	db.Close()

}

func CheckDBErr(err error) {
	if err != nil {
		panic(err)
	}
}
