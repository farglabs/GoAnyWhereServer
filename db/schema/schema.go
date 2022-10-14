package schema

import (
	"database/sql"
)

type schema struct{}

func (d *schema) Error() string {
	return d.Error()
}

func CreateTables(db *sql.DB) (string, error) {
	users_table := `CREATE TABLE "files" (
        "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
        "username" VARCHAR(64) NULL,
        "departname" VARCHAR(64) NULL,
        "created" DATE NULL
    );`
	query, err := db.Prepare(users_table)
	if err != nil {
		return "", err
	}
	query.Exec()
	return "Table created successfully!", &schema{}
}
