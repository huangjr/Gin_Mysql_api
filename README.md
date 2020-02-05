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

    * Check the user's information from database and the user's id as input  

        > http://localhost:8000/user/:id

     * Check all the users' information from database  
        > http://localhost:8000/users

* POST
    * Add one user into database, and it needs first_name and last_name as keys in body 
        > http://localhost:8000/user

    * Add users into database, and it needs json as input
        > http://localhost:8000/users/AddUsers
        
        Below is the example for json to add users
        ~~~sql
        {
        "data":[{
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
        
        Below is the example for json to delete users by its ids
        ~~~sql
        {
        "data":[{
        	"id":5
            },{
        	"id":6
            }]
        }
        ~~~
* DELETE
    * Delete one user by its id 
        > http://localhost:8000/user/:id

* PUT
    * Modify the user's information by his id, including his first_name and last_name

        > http://localhost:8000/user/:id

