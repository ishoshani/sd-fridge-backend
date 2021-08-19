package db

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID   primitive.ObjectID `json:"_id"`
	Name string             `json:"name"`
}

func InitializeItemsTable(db *sql.DB) {
	stmt, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS items (" +
			"ID VARCHAR(255) Primary Key UNIQUE," +
			"NAME VARCHAR(255) UNIQUE" +
			");")
	if err != nil {
		panic(err)
	}
	stmt.Exec()
	if err != nil {
		panic(err)
	}
}
