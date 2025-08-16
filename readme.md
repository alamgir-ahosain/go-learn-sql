##  Project Structure
```plaintext
GO-LEARN-SQL/
├── Mysql/                          # MySQL implementation
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
├── PostgreSQL/                     # (Future PostgreSQL support)
│
├── SQLite/                         # (Future SQLite support)
│
├── go.mod                          # Go module definition
├── go.sum                          # Go dependencies checksums
└── readme.md                       # Project documentation
```

##  Database Setup

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
