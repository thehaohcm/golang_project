package unit_testing

import (
	"fmt"
	"golang_project/repositories"
	"testing"

	// db "github.com/mattn/go-sqlite3"
	"database/sql"
	"os"

	"github.com/stretchr/testify/assert"
)

var (
	repo repositories.FriendConnectionRepository = repositories.New()
)

func TestMain(m *testing.M) {
	// os.Exit skips defer calls
	// so we need to call another function
	code, err := run(m)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}

func run(m *testing.M) (code int, err error) {
	// pseudo-code, some implementation excluded:
	//
	// 1. create test.db if it does not exist
	// 2. run our DDL statements to create the required tables if they do not exist
	// 3. run our tests
	// 4. truncate the test db tables

	db, err := sql.Open("sqlite3", "file: golang_project.db")
	if err != nil {
		return -1, fmt.Errorf("could not connect to database: %w", err)
	}

	// truncates all test data after the tests are run
	defer func() {
		for _, t := range []string{"friends", "subscribers", "users"} {
			_, _ = db.Exec(fmt.Sprintf("DELETE FROM %s", t))
		}

		db.Close()
	}()

	return m.Run(), nil
}

func TestCreateFriendConnection(t *testing.T) {

	result, tx := repo.CreateFriendConnection([]string{
		"abc@def.com",
		"abc1@def.com",
	})

	//rollback db
	fmt.Println("hao")
	if tx != nil {
		fmt.Println("dfa")
	} else {
		fmt.Println("fda")
	}
	assert.Equal(t, true, result)

}
