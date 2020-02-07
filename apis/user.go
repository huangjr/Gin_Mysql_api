package apis

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"Gin_Mysql_api/models"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works.")
}

func CommonAddUser(user models.User) (msg string) {

	firstName := user.FirstName
	lastName := user.LastName

	log.Println(firstName, lastName)

	p := models.User{FirstName: firstName, LastName: lastName}

	ra, err := p.AddUser()
	if err != nil {
		log.Println(err)
	}

	msg = fmt.Sprintf("insert successful %d", ra)
	return msg
}

func AddUserApi(c *gin.Context) {

	var user models.User

	err := c.ShouldBind(&user)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	msg := CommonAddUser(user)

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func AddUsersApi(c *gin.Context) {

	var users models.Users

	err := c.ShouldBindJSON(&users)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(users.Users) == 0 {
		log.Println("empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty"})
		return
	}

	for _, user := range users.Users {

		var msg string

		if user.FirstName == "" || user.LastName == "" {
			msg = fmt.Sprintf("invalid user.FirstName and user.LastName")
			c.JSON(http.StatusBadRequest, gin.H{"msg": msg})

		} else {
			msg = CommonAddUser(user)
			c.JSON(http.StatusOK, gin.H{"msg": msg})
		}

		log.Println(msg)

	}
}

func DelUserIdsApi(c *gin.Context) {
	var users models.Users
	// var persons []models.Person
	err := c.ShouldBindJSON(&users)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(users.Users) == 0 {
		log.Println("empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty"})
		return
	}

	for _, user := range users.Users {

		id := user.Id
		log.Println(id)
		p := models.User{Id: id}
		ra, err := p.DelUser()
		if err != nil {
			log.Println(err)
		}
		var msg string

		if ra == 0 {
			msg = fmt.Sprintf("Affected %d row.", ra)
			c.JSON(http.StatusBadRequest, gin.H{"msg": msg})
		} else {
			msg = fmt.Sprintf("Delete person %d successful, Affected %d row.", id, ra)
			c.JSON(http.StatusOK, gin.H{"msg": msg})
		}
	}
}

func DelUserApi(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	p := models.User{Id: id}
	ra, err := p.DelUser()

	if err != nil {
		log.Println(err)
	}
	var msg string

	if ra == 0 {
		msg = fmt.Sprintf("Affected %d row.", ra)
		c.JSON(http.StatusBadRequest, gin.H{"msg": msg})
	} else {
		msg = fmt.Sprintf("Delete person %d successful, Affected %d row.", id, ra)
		c.JSON(http.StatusOK, gin.H{"msg": msg})
	}

}

func ModUserApi(c *gin.Context) {
	// cid := c.Param("id")
	// id, _ := strconv.Atoi(cid)

	// p := models.User{Id: id}

	// err := c.ShouldBind(&p)
	// if err != nil {
	// 	log.Println(err)
	// }

	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)

	p := models.User{Id: id}

	err := c.ShouldBind(&p)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ra, err := p.ModUser()
	if err != nil {
		log.Println(err)
	}

	var msg string
	if ra == 0 {
		msg = fmt.Sprintf("Update %d person", ra)
		c.JSON(http.StatusBadRequest, gin.H{"msg": msg})
	} else {
		msg = fmt.Sprintf("Update person %d successful", p.Id)
		c.JSON(http.StatusOK, gin.H{"msg": msg})
	}
}

func GetUserApi(c *gin.Context) {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)
	p := models.User{Id: id}

	user, err := p.GetUser()
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func GetUsersApi(c *gin.Context) {
	p := models.User{}

	users, err := p.GetUsers()
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
