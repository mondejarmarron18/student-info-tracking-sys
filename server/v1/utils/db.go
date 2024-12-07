package utils

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	conn *sql.DB
	once sync.Once
)

func NewDB(dbUrl string) {
	once.Do(func() {
		log.Println("Initializing database connection...")
		db, err := sql.Open("postgres", dbUrl)

		if err != nil {
			log.Fatalln("Error opening database:", err)
		}

		if err := db.Ping(); err != nil {
			log.Fatalln("Error pinging database:", err)
		}

		conn = db
		log.Println("Database connection initialized successfully.")
	})

}

func GetConn() *sql.DB {
	if conn == nil {
		log.Println("Database not initialized")
		return nil
	}

	return conn
}

func CloseConn() {
	if conn == nil {
		log.Println("Database not initialized")
		return
	}

	conn.Close()
}
