package main

import (
	"Gin_Mysql_api/routers"
)

func main() {

	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "It works")
	// })
	router := routers.InitRouter()

	router.Run(":8000")
}
