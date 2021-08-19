package db

import (
	"database/sql"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FridgeItemDB struct {
	ID          primitive.ObjectID `json:"_id"`
	FridgeID    primitive.ObjectID `json:"fridgeID"`
	ItemID      primitive.ObjectID `json:"itemID"`
	NeedsRefill bool               `json:"needsRefill"`
}

func InitializeFridgeItemTable(db *sql.DB) {
	stmt, err := db.Prepare(
		`CREATE TABLE IF NOT EXISTS fridgeitems ( 
				ID VARCHAR(255) Primary Key UNIQUE, 
				fridgeid VARCHAR(255) NOT NULL, 
				itemid VARCHAR(255) NOT NULL, 
				needsrefill boolean, 
				CONSTRAINT fk_fridge 
				FOREIGN KEY(fridgeid) 
				REFERENCES fridges(ID) 
				ON DELETE CASCADE, 
				CONSTRAINT fk_item 
				FOREIGN KEY(itemid) 
				REFERENCES items(id) 
				ON DELETE CASCADE 
				);`)

	if err != nil {
		panic(err)
	}
	stmt.Exec()
	if err != nil {
		panic(err)
	}
}
