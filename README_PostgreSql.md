## GO-Learn-SQL
Go database examples: Connect and perform CRUD operations (`Create`, `Read`, `Update`, `Delete`) with MySQL · PostgreSQL · SQLite
<br>

* [MySQL README](README.md)
* [PostgreSQL README](README_PostgreSql.md)
* [SQLite README](README_SQLite.md)
  <br>
>For PostgreSQL
## Installation
1. **Clone the Repository**
 ```bash
https://github.com/alamgir-ahosain/go-learn-sql.git
```
2. **Install Dependencies**<br>
 ```bash
go get github.com/lib/pql
go get github.com/joho/godotenv
```
3. **Navigate into the prject:**
```bash
cd go-learn-sql/PostgreSQL
```

##  Database Setup
Login to PostgreSQL and create the database and table:
```sql
CREATE DATABASE goDB;
USE goDB;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,         
    sid VARCHAR(128) NOT NULL,
    name VARCHAR(128) NOT NULL,
    cgpa NUMERIC(5,2)             
);
```
## Run the Application:
1. Start PostgreSQL server.  
2. Update database connection in `db/postgreSql.go`
3. Configure `.env` in `config/.env`
4. Run the Project:
```go
go run main.go
```
   
##  Project Structure
```plaintext

GO-LEARN-SQL/

├── Mysql/                          # MySql project Structure
├── PostgreSQL/                     # PostgreSQL implementation
│   ├── cmd/
│   │   └── MyApp/
│   │       └── server.go           # Entry point for PostgreSQL server
│   ├── controllers/
│   │   └── user_controller.go      # Handles user-related requests
│   ├── db/
│   │   └── PostgreSql.go           # PostgreSQL database connection setup
│   ├── models/
│   │   └── user.go                 # User model definition
│   └── main.go                     # Main entry file for PostgreSQL
│
│
├── SQLite/                         # SQLite project Structure
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
