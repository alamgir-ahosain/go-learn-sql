package myapp

import (
	"fmt"

	"github.com/alamgir-ahosain/learn-sql/SQLite/controllers"
	dbsqlite "github.com/alamgir-ahosain/learn-sql/SQLite/db_sqlite"
)

func RunApp() {
	dbsqlite.Connect()
	defer dbsqlite.DB.Close() //close the DB connection when main exist

	//Create user
	controllers.CreateUser("CE21012", "Alamgir", 3.48)

	//Read user By Id
	user, err := controllers.ReadUserById(2)
	controllers.CheckError(err)
	fmt.Println(user)

	//Read Multiple user by Field
	users, er := controllers.ReadUserByField("CE21012")
	controllers.CheckError(er)
	fmt.Println(users)

	//Update user
	controllers.UpdateUser(3, "Md Mustakim Mia")

	//Delete user by ID
	controllers.DeleteUser(7)

}
