package dao

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

type DaoNames struct {
	DB *sql.DB
}

func (d *DaoNames) InsertName(name string) (int, string, error) {
	sqlStatement := `INSERT INTO hello_world (hw_text) VALUES ($1) RETURNING id, hw_text`
	id, hw_text := 0, ""
	err := d.DB.QueryRow(sqlStatement, name).Scan(&id, &hw_text)
	if err != nil {
		return 0, "", errors.Wrap(err, "there was an error while inserting the name")
	}
	log.Printf("New hello_world record with ID '%v' and text '%s' was inserted\n", id, hw_text)
	return id, hw_text, nil
}
