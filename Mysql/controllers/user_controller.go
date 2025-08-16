package controllers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/alamgir-ahosain/learn-sql/Mysql/db"
	"github.com/alamgir-ahosain/learn-sql/Mysql/models"
)

var users = []models.User{}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// _______________________CREATE________________________________

func CreateUser(sid, name string, cgpa float32) {

	query := "insert into users(sid,name,cgpa) values(?,?,?)"
	row, err := db.DB.Exec(query, sid, name, cgpa)
	CheckError(err)
	id, err := row.LastInsertId()
	CheckError(err)

	//add New User to Users slice
	NewUser := models.User{
		ID:   int(id),
		SID:  sid,
		Name: name,
		CGPA: cgpa,
	}
	users = append(users, NewUser)
	fmt.Println("User Created Succesfully")
	fmt.Println(users)
}

// _______________________READ________________________________

func ReadUserById(id int) (*models.User, error) {
	var user models.User
	query := db.DB.QueryRow("select * from users where id=?", id)
	err := query.Scan(&user.ID, &user.SID, &user.Name, &user.CGPA)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal(id, "user not found")
		}
		return nil, err
	}
	return &user, nil
}

func ReadUserByField(sid string) ([]models.User, error) {
	var users []models.User

	query, err := db.DB.Query("select * from users where sid=?", sid)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	// Add users returned by the query to the slice
	for query.Next() {
		var row models.User
		err := query.Scan(&row.ID, &row.SID, &row.Name, &row.CGPA)
		if err != nil {
			return nil, err
		}
		users = append(users, row)
	}
	//check during iteration
	err = query.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// _______________________UPDATE________________________________

func UpdateUser(id int, NewName string) {

	query := "update users set name=? where id=?"
	row, err := db.DB.Exec(query, NewName, id)
	CheckError(err)
	rowsAffected, er := row.RowsAffected()
	CheckError(er)
	if rowsAffected == 0 {
		fmt.Println("No user found with that ID ")
	} else {
		fmt.Println("User Updated Succesfully")

	}
}
