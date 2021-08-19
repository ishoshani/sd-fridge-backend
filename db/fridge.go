package db

import (
	"database/sql"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Fridge struct {
	ID      primitive.ObjectID `json:"_id",`
	Name    string             `json:"name"`
	Address string             `json:"address"`
}

func InitializeFridgeTable(db *sql.DB) {
	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS fridges (
			ID VARCHAR(255) Primary Key UNIQUE,
			NAME VARCHAR(255) UNIQUE,
			ADDRESS VARCHAR(255) UNIQUE
			);`)
	if err != nil {
		panic(err)
	}
	stmt.Exec()
	if err != nil {
		panic(err)
	}
}

func InsertFridge(db *sql.DB, fridge Fridge) {
	stmt, err := db.Prepare(fmt.Sprintf(`INSERT INTO fridges (ID, NAME, ADDRESS) VALUES ('%s', '%s', '%s');`, fridge.ID.Hex(), fridge.Name, fridge.Address))
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec()
	if err != nil {
		panic(err)
	}
	log.Default().Print(result)
}

func GetFridges(db *sql.DB) []Fridge {
	stmt, err := db.Prepare("SELECT * FROM fridges;")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query()
	if err != nil {
		panic(err)
	}
	fridges := make([]Fridge, 0)
	for rows.Next() {
		var fridge *Fridge
		if err := rows.Scan(&fridge); err != nil {
			panic(err)
		}
		fridges = append(fridges, *fridge)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}
	rows.Close()
	return fridges
}
