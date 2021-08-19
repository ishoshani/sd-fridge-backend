package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sdfridgeproject.com/m/v2/db"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func AddFridge(c *gin.Context) {
	var JsonFridge struct {
		Name    string `json:"name"`
		Address string `json:"address"`
	}
	if err := c.BindJSON(&JsonFridge); err != nil {
		panic(err)
	}
	fridge := db.Fridge{
		ID:      primitive.NewObjectID(),
		Name:    JsonFridge.Name,
		Address: JsonFridge.Address,
	}
	db.InsertFridge(db.DB, fridge)
}

func GetFridges(c *gin.Context) {
	fridges := db.GetFridges(db.DB)
	c.JSON(http.StatusOK, fridges)
}
