package config

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var lock = &sync.Mutex{}

var db *sql.DB

var dbTest *sql.DB

// singleton pattern
func GetDBInstance() *sql.DB {
	if db == nil {
		lock.Lock()
		defer lock.Unlock()
		if db == nil {
			fmt.Println("Creating single instance now.")
			db = connectDatabase()
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return db
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}

func getDBInfo() string {
	return fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB_NAME"))
}

func connectDatabase() *sql.DB {
	db, err := sql.Open("postgres", getDBInfo())
	if err != nil {
		panic(err)
	}
	fmt.Print("connected to db")
	return db
}

func GetTestDB() *sql.DB {
	if dbTest == nil {
		lock.Lock()
		defer lock.Unlock()
		if dbTest == nil {
			fmt.Println("Creating single instance now.")
			db = connectDatabase()
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}

	return dbTest
}

func connectTestingDatabase() *sql.DB {
	dbTest, err := sql.Open("postgres", getDBInfo())
	if err != nil {
		panic(err)
	}
	fmt.Print("connected to db")
	return dbTest
}
