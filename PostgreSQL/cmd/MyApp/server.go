package myapp

import (
	

	"github.com/alamgir-ahosain/learn-sql/PostgreSQL/controllers"
	"github.com/alamgir-ahosain/learn-sql/PostgreSQL/db"
)

func RunApp() {
	db.Connect()
	defer db.DB.Close() //close the DB connection when main exist

	//Create user
	controllers.CreateUser("CE21012", "Alamgir", 3.48)

	// //Read user By Id
	// user, err := controllers.ReadUserById(3)
	// controllers.CheckError(err)
	// fmt.Println(user)

	// //Read Multiple user by Field
	// users, er := controllers.ReadUserByField("CE21012")
	// controllers.CheckError(er)
	// fmt.Println(users)

	// //Update user
	// controllers.UpdateUser(33, "Alamgir Hosain")

	// //Delete user by ID
	// controllers.DeleteUser(3)

}
