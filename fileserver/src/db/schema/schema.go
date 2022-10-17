package schema

import (
	"database/sql"
)

func CreateTables(db *sql.DB) {
	users_table := `CREATE TABLE "files" (
        "uid" INTEGER PRIMARY KEY AUTOINCREMENT,
        "username" VARCHAR(64) NULL,
        "departname" VARCHAR(64) NULL,
        "created" DATE NULL
    );`

	query, err := db.Prepare(users_table)
	if err != nil {
		panic(err)
	}

	query.Exec()
}
