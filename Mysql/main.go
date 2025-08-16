package main

import (
	"github.com/alamgir-ahosain/learn-sql/Mysql/controllers"
	"github.com/alamgir-ahosain/learn-sql/Mysql/db"
)

func main() {
	db.Connect()
	defer db.DB.Close() //close the DB connection when main exist
	controllers.CreateUser("CE21012","Alamgir",3.48)



}