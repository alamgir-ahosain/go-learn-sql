## GO-Learn-SQL
Go database examples: Connect and perform CRUD operations (`Create`, `Read`, `Update`, `Delete`) with MySQL · PostgreSQL · SQLite
<br>

* [MySQL README](README.md)
* [PostgreSQL README](README_PostgreSql.md)
* [SQLite README](README_SQLite.md)
  <br>
>For SQLite
## Installation
1. **Clone the Repository**
 ```bash
https://github.com/alamgir-ahosain/go-learn-sql.git
```
2. **Install Dependencies**<br>
 ```bash
go get github.com/mattn/go-sqlite3
go get github.com/joho/godotenv
```
3. **Navigate into the prject:**
```bash
cd go-learn-sql/SQlite
```

##  Database Setup
SQLite stores the database in a file(`storage/goDB.db`)
The Table will create automatically if not exist.
```sql
CREATE TABLE IF NOT EXISTS users (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    sid TEXT,
    name TEXT,
    cgpa DECIMAL(5,2)
);
```
## Run the Application:
1. Run the Project:
```go
go run main.go
```
   
##  Project Structure
```plaintext

GO-LEARN-SQL/
├── Mysql/                          # MySql project Structure
├── PostgreSQL/                     # PostgreSQL project Structure
├── SQLite/                         # SQLite Project Structure
│   ├── cmd/
│   │   └── MyApp/
│   │       └── server.go           # Entry point for SQLite server
│   ├── controllers/
│   │   └── user_controller.go      # Handles user-related requests
│   ├── db_sqlite/
│   │   └── sqlite.go               # SQLite database connection setup
│   ├── models/
│   │   └── user.go                 # User model definition
│   ├── storage/
│   │   └── goDB.db                 # database file here
│   └── main.go                     # Main entry file for PostgreSQL
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
