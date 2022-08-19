package config

import (
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

var db *sql.DB

var dbTest *sql.DB

var dbInfo = "host=" + os.Getenv("POSTGRES_HOST") + " port=" + os.Getenv("POSTGRES_PORT") + " user=" + os.Getenv("POSTGRES_USER") + " " +
	"password=" + os.Getenv("POSTGRES_PASSWORD") + " dbname=" + os.Getenv("POSTGRES_DB_NAME") + " sslmode=disable"

// singleton pattern
func GetDBInstance() *sql.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			log.Println("Creating single instance now.")
			db = connectDatabase()
		} else {
			log.Println("Single instance already created.")
		}
	} else {
		log.Println("Single instance already created.")
	}

	return db
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func connectDatabase() *sql.DB {
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	log.Print("connected to db")
	return db
}
