package models

import (
	"log"

	db "Gin_Mysql_api/database"
)

type User struct {
	Id        int    `json:"id" form:"id"`
	FirstName string `json:"first_name" form:"first_name" binding:"required"`
	LastName  string `json:"last_name" form:"last_name" binding:"required"`
}

type Users struct {
	Users []User `json:"data" binding:"required"`
}

func (p *User) AddUser() (id int64, err error) {
	stmtIn, err := db.SqlDB.Prepare("INSERT INTO user(firstname, lastname) VALUES(?, ?)")
	if err != nil {
		log.Fatalln(err)
	}

	res, err := stmtIn.Exec(p.FirstName, p.LastName)
	if err != nil {
		log.Fatalln(err)
	}

	id, err = res.LastInsertId()
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func (p *User) DelUser() (ra int64, err error) {
	stmtOut, err := db.SqlDB.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		log.Println(err)
	}

	res, err := stmtOut.Exec(p.Id)
	if err != nil {
		log.Println(err)
	}

	ra, _ = res.RowsAffected()

	return
}

func (p *User) ModUser() (ra int64, err error) {
	stmtInt, err := db.SqlDB.Prepare("UPDATE user SET firstname=?, lastname=? WHERE id=?")
	if err != nil {
		log.Println(err)
	}

	res, err := stmtInt.Exec(p.FirstName, p.LastName, p.Id)
	if err != nil {
		log.Println(err)
	}

	ra, _ = res.RowsAffected()

	return
}

func (p *User) GetUser() (person User, err error) {
	stmtOut, err := db.SqlDB.Prepare("SELECT id, firstname, lastname FROM user WHERE id=?")
	if err != nil {
		log.Println(err)
	}

	err = stmtOut.QueryRow(p.Id).Scan(&person.Id, &person.FirstName, &person.LastName)
	if err != nil {
		log.Println(err)
	}

	return
}

func (p *User) GetUsers() (users []User, err error) {
	users = make([]User, 0)

	stmsOut, err := db.SqlDB.Prepare("SELECT id, firstname, lastname FROM user")
	if err != nil {
		log.Println(err)
	}

	rows, err := stmsOut.Query()
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.FirstName, &user.LastName)
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}

	return
}
