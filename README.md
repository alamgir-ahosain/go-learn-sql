## GO-Learn-SQL
Go database examples: Connect and perform CRUD operations (`Create`, `Read`, `Update`, `Delete`) with MySQL · PostgreSQL · SQLite
<br>

* [MySQL README](README.md)
* [PostgreSQL README](README_PostgreSql.md)
* [SQLite README](README_SQLite.md)
  <br>
>For MySQL
## Installation
1. **Clone the Repository**
 ```bash
https://github.com/alamgir-ahosain/go-learn-sql.git
```
2. **Install Dependencies**<br>
 ```bash
go get -u github.com/go-sql-driver/mysql
```
3. **Navigate into the prject:**
```bash
cd go-learn-sql/Mysql
```

##  Database Setup
Login to MySQL and create the database and table:
```sql
CREATE DATABASE goDB;
USE goDB;

CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL,
    sid VARCHAR(128) NOT NULL,
    name VARCHAR(128) NOT NULL,
    cgpa DECIMAL(5,2),
    PRIMARY KEY (`id`)
);
```
## Run the Application:
1. Start MySQL server.  
2. Update database connection in `db/mysql.go`:  

   i. With Password:  
   ```go
   dsn := "root:password@tcp(127.0.0.1:3306)/goDB"
   ```
   ii. Without Password<br>
```go
    dsn := "root:password@tcp(127.0.0.1:3306)/goDB"
 ```
3.Run the Project:
```go
go run main.go
```
   
##  Project Structure
```plaintext
GO-LEARN-SQL/
├── Mysql/                          # MySql project Structure
│   ├── cmd/
│   │   └── MyApp/
│   │       └── server.go           # Entry point for MySQL server
│   ├── controllers/
│   │   └── user_controller.go      # Handles user-related requests
│   ├── db/
│   │   └── mysql.go                # MySQL database connection setup
│   ├── models/
│   │   └── user.go                 # User model definition
│   └── main.go                     # Main entry file for MySQL
│
├── PostgreSQL/                     # PostgreSQL project Structure
│
├── SQLite/                         # SQLite project Structure
│
├── go.mod                          # Go module definition
├── go.sum                          # Go dependencies checksums
└── README.md                       # Project documentation for MySQL
└── README_PostgreSql.md            # Project documentation for PostgreSQL
└── README_SQLite.md                # Project documentation for SQLite
```
---
# Thank You for Checking Out This Repository!
Your feedback is valuable! Please share your thoughts and suggestions for improvement via [GitHub Issues](https://github.com/alamgir-ahosain/go-learn-sql/issues).

# Contributing  
Contributions are welcome! Feel free to fork the repo and create a pull request.
