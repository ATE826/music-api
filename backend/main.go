package main

import (
	"backend/database"
	"backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.ConnectDB()
	routers.SetupRouters(r)
	r.Run(":8080")
}
