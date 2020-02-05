package routers

import (
	"Gin_Mysql_api/apis"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", apis.IndexApi)
	r.POST("/user", apis.AddUserApi)
	r.DELETE("/user/:id", apis.DelUserApi)
	r.PUT("/user/:id", apis.ModUserApi)
	r.GET("/user/:id", apis.GetUserApi)
	r.GET("/users", apis.GetUsersApi)

	// use json to add and delete multiple persons
	r.POST("/users/DeleteUserByIds", apis.DelUserIdsApi)
	r.POST("/users/AddUsers", apis.AddUsersApi)

	return r
}
