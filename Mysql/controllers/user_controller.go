package controllers

import (
	"fmt"
	"log"

	"github.com/alamgir-ahosain/learn-sql/Mysql/db"
)

func CreateUser(sid, name string, cgpa float32) {
	
	query := "insert into users(sid,name,cgpa) values(?,?,?)"
	_, err := db.DB.Exec(query, sid, name, cgpa)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User Created Succesfully")

}
