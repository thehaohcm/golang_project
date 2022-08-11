package database

import (
	"database/sql"
	"fmt"
	"runtime"
	"sync"

	_ "github.com/lib/pq"
)

var host = "127.0.0.1"

func init() {
	if runtime.GOOS == "darwin" {
		host = "host.docker.internal"
	}
}

const (
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "golang_project"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)

var lock = &sync.Mutex{}

var db *sql.DB

// singleton pattern
func GetInstance() *sql.DB {
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

func connectDatabase() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Print("connected to db")
	return db
}

var dbTest *sql.DB

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
	dbTest, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	fmt.Print("connected to db")
	return dbTest
}
