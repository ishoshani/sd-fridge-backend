package routes

import "github.com/gin-gonic/gin"

func App() {
	r := gin.Default()
	r.GET("/ping", Ping)

	r.POST("/addFridges", AddFridge)
	r.GET("/getFridges", GetFridges)

	r.Run()
}
