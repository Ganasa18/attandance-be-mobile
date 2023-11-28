package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

// var err error

func NewDB() *sql.DB {
	var (
		// connection = os.Getenv("DB_CONNECTION")
		port     = os.Getenv("DB_PORT")
		host     = os.Getenv("DB_HOST")
		dbname   = os.Getenv("DB_NAME")
		username = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASSWORD")
	)

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbname)
	dsn := "postgresql://" + username + ":" + password + "@" + host + ":" + port + "/" + dbname + "?sslmode=disable"
	fmt.Println("CONNECTING....")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatal("CHECK CONNECTING : ", err)
		defer db.Close() // Close the connection when done
	}

	fmt.Println("CONNECTING TO DATABASE")

	return db

	// db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_database_migration")
	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third
	// migrate create -ext sql -dir db/migrations sample_dirty_state
	// migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations up
	// migrate -database "postgres://postgres:admin@localhost:5432/db_attandance_mobile?sslmode=disable" -path db/migrations up
	// migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations down
	// migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations version
	// migrate -database "mysql://root@tcp(localhost:3306)/belajar_golang_database_migration" -path db/migrations force 20220922043738
}

func GetDb() *sql.DB {

	return db
}
