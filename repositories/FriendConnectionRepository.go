package repositories

import (
	"context"
	"database/sql"
	"golang_project/models"
	"golang_project/utils"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type FriendConnectionRepository interface {
	FindFriendsByEmail(email string) []string
	FindCommonFriendsByEmails(emails []string) []string
	CreateFriendConnection(emails []string) (bool, *sql.Tx)
	SubscribeFromEmail(req models.SubscribeRequest) (bool, *sql.Tx)
	BlockSubscribeByEmail(req models.BlockSubscribeRequest) (bool, *sql.Tx)
	GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse
}

type repository struct {
	db  *sql.DB
	ctx context.Context
}

func New() FriendConnectionRepository {
	return &repository{
		db:  utils.GetInstance(),
		ctx: context.Background(),
	}
}

//1.
func (repo *repository) CreateFriendConnection(emails []string) (bool, *sql.Tx) {

	//check empty or invalid email format
	if len(emails) < 2 || len(emails) > 2 {
		panic("invalid request")
	}
	checkValidEmails(emails)

	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec("INSERT INTO friends('user_email','friend_email') VALUES(?,?)", emails[0], emails[1])

	if err != nil {
		// return false
		panic(err)
	}
	return true, tx
}

//2.
func (repo *repository) FindFriendsByEmail(email string) []string {
	checkValidEmails([]string{email})

	rows, err := repo.db.Query("SELECT friend_email FROM friends WHERE user_email=? AND BLOCKED=0", email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var friends []string
	for rows.Next() {
		var friend string
		if err := rows.Scan(&friend); err != nil {
			panic(err)
		}
		friends = append(friends, friend)
	}
	return friends
}

//3.
func (repo *repository) FindCommonFriendsByEmails(emails []string) []string {
	checkValidEmails(emails)
	sqlStatement := "SELECT friend_email FROM friends WHERE"
	for _, email := range emails {
		sqlStatement += " user_email='" + email + "' OR"
	}
	if len(emails) > 1 {
		sqlStatement = sqlStatement[:len(sqlStatement)-3]
	}
	sqlStatement += " AND BLOCKED=0 GROUP BY friend_email HAVING COUNT(*) > 1"
	rows, err := repo.db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var friends []string
	for rows.Next() {
		var friend string
		if err := rows.Scan(&friend); err != nil {
			panic(err)
		}
		friends = append(friends, friend)
	}
	return friends
}

//4.
func (repo *repository) SubscribeFromEmail(req models.SubscribeRequest) (bool, *sql.Tx) {
	checkValidEmails([]string{req.Requestor, req.Target})
	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec("INSERT INTO subscribers('requestor','target') VALUES (?,?)", req.Requestor, req.Target)
	if err != nil {
		panic(err)
		// return false
	}

	return true, tx
}

//5.
func (repo *repository) BlockSubscribeByEmail(req models.BlockSubscribeRequest) (bool, *sql.Tx) {
	checkValidEmails([]string{req.Requestor, req.Target})
	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		panic(err)
	}

	//suppose A block B:

	if repo.hasFriendConnection(req.Requestor, req.Target) {
		//if A and B are friend, A no longer receive notify from B
		res, err := tx.Exec("INSERT OR REPLACE INTO subscribers('requestor','target', 'blocked') VALUES (?,?,1)", req.Requestor, req.Target)
		if err != nil {
			panic(err)
		}

		if _, err = res.LastInsertId(); err != nil {
			panic(err)
		}
	} else {
		//if not friend, no new friend connection added
		res, err := tx.Exec("INSERT OR REPLACE INTO friends('user_email','friend_email', 'blocked') VALUES (?,?,1)", req.Requestor, req.Target)
		if err != nil {
			panic(err)
		}
		if _, err = res.LastInsertId(); err != nil {
			panic(err)
		}
	}

	return true, tx
}

func (repo *repository) hasFriendConnection(requestor string, target string) bool {
	checkValidEmails([]string{requestor, target})
	rows, err := repo.db.Query("SELECT * FROM friends WHERE user_email=? AND friend_email=? AND BLOCKED=0", requestor, target)
	if err != nil {
		return false
	}
	defer rows.Close()

	if rows.Next() {
		return true
	}
	return false
}

//6.
func (repo *repository) GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) models.GetSubscribingEmailListResponse {
	checkValidEmails([]string{req.Sender})
	var res models.GetSubscribingEmailListResponse
	res.Success = false

	//if has a friend connection
	rows, err := repo.db.Query("SELECT friend_email FROM friends WHERE user_email=? AND blocked=0", req.Sender)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var recipients []string
	for rows.Next() {
		var friend string
		if err := rows.Scan(&friend); err != nil {
			panic(err)
		}
		recipients = append(recipients, friend)
	}

	//if has a friend connection, but blocked in subscribers tables
	if len(recipients) > 0 {
		rows, err := repo.db.Query("SELECT target FROM subscribers WHERE requestor=? AND blocked=1", req.Sender)
		if err != nil {
			panic(err)
		}

		defer rows.Close()

		var blockedEmails []string
		for rows.Next() {
			var target string
			if err := rows.Scan(&target); err != nil {
				panic(err)
			}
			blockedEmails = append(blockedEmails, target)
		}
		recipients = utils.GetDifference(recipients, blockedEmails)
	}

	//if subscribed to updates
	rows, err = repo.db.Query("SELECT requestor FROM subscribers WHERE target=? AND blocked=0", req.Sender)
	if err != nil {
		return res
	}

	defer rows.Close()
	for rows.Next() {
		var friend string
		if err := rows.Scan(&friend); err != nil {
			return res
		}
		recipients = append(recipients, friend)
	}

	//if being mentioned in the update
	textArr := strings.Split(req.Text, " ")
	for _, text := range textArr {
		if utils.IsEmailValid(text) {
			recipients = append(recipients, text)
		}
	}

	if len(recipients) > 0 {
		res.Recipients = recipients
		res.Success = true
	}
	return res
}

func checkValidEmails(emails []string) {
	if emails == nil {
		panic("empty request")
	}
	for _, email := range emails {
		if strings.TrimSpace(email) == "" || !utils.IsEmailValid(email) {
			panic("invalid request")
		}
	}
}
