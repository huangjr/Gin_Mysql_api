# Gin_Mysql_api

----
Restful API includes using gin and Mysql

### Download  
```
git clone https://github.com/huangjr/Gin_Mysql_api.git
```
### Quick Start
**Requirement**
* go version go1.13.6
* MySQL 8.0.11

**MySQL Setting**

Database Name: mydb  
User Account: adm_user1   
Password: 12345678  
Network Address: 127.0.0.1  
TCP/IP Port: 3306  
db.sql code:  
~~~sql
-- MySql 

create database if not exists mydb;
use mydb;
create table user (
	id int(11) not null auto_increment primary key,
	firstname varchar(40) not null default "",
	lastname  varchar(40) not null default ""
) engine=innodb default charset=utf8;
~~~

**Build & Run**
```
cd Gin_Mysql_api  
go build main.go  
./main
```

**API Demo**

* GET  
    * Test whether it could run on port 8000 
        > http://localhost:8000/

		e.g.
		> curl -X GET http://localhost:8000/

    * Check the user's information from database by giving the user's id  

        > http://localhost:8000/user/:id
		
		e.g.
		> curl -X GET http://localhost:8000/user/2

     * Check all the users' information from database  
        > http://localhost:8000/users

		e.g.
		> curl -X GET http://localhost:8000/users

* POST
    * Add one user into database, giving first_name and last_name as keys in body 
        > http://localhost:8000/user

		e.g.  
		> curl -X POST http://localhost:8000/user -d 'first_name=Amy&last_name=Huang'

    * Add users into database, and it needs json as input
        > http://localhost:8000/users/AddUsers

		e.g.
		> curl -X POST http://localhost:8000/users/AddUsers -d '{"users":[{"first_name":"Amy","last_name":"Huang"},{"first_name":"John","last_name":"Chang"}]}'
        
        Below is the example for json to add users
        ~~~sql
        {
        "users":[{
        	"first_name":"Amy",
        	"last_name":"Huang"
            },{
        	"first_name":"John",
        	"last_name":"Chang"
            }]
        }
        ~~~
    * "Delete" users by its ids and it needs json as input

        > http://localhost:8000/users/DeleteUserByIds
        
		e.g.  
		> curl -X POST http://localhost:8000/users/DeleteUserByIds -d '{"users":[{"id":99},{"id":100}]}'

        Below is the example for json to delete users by its ids
        ~~~sql
        {
        "users":[{
        	"id":5
            },{
        	"id":6
            }]
        }
        ~~~
* DELETE
    * Delete one user by its id 
        > http://localhost:8000/user/:id
		e.g.
		> curl -X DELETE http://localhost:8000/user/50

* PUT
    * Modify the user's information by his id, including his first_name and last_name

        > http://localhost:8000/user/:id
		e.g. 
		> curl -X PUT http://localhost:8000/user/2 -d 'first_name=Amy&last_name=Huang'

