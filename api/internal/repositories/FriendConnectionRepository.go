package repositories

import (
	"context"
	"database/sql"
	"strings"

	"golang_project/api/cmd/serverd/utils"
	"golang_project/api/internal/models"

	_ "github.com/lib/pq"
)

type FriendConnectionRepository interface {
	FindFriendsByEmail(email string) ([]string, error)
	FindCommonFriendsByEmails(emails []string) ([]string, error)
	CreateFriendConnection(emails []string) (bool, error)
	SubscribeFromEmail(req models.SubscribeRequest) (bool, error)
	BlockSubscribeByEmail(req models.BlockSubscribeRequest) (bool, error)
	GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) ([]string, error)
}

type repository struct {
	db  *sql.DB
	ctx context.Context
}

func New(db *sql.DB) FriendConnectionRepository {
	return &repository{
		db:  db,
		ctx: context.Background(),
	}
}

// 1.
func (repo *repository) CreateFriendConnection(emails []string) (bool, error) {

	//check empty or invalid email format
	if len(emails) < 2 || len(emails) > 2 {
		panic("invalid request")
	}
	if valid, err := utils.CheckValidEmails(emails); !valid {
		return valid, err
	}

	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		return false, err
	}
	_, err = tx.Exec("INSERT INTO public.friends(user_email,friend_email) VALUES($1,$2)", emails[0], emails[1])

	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			panic(err1)
		}
		return false, err
	}

	tx.Commit()
	return true, nil
}

// 2.
func (repo *repository) FindFriendsByEmail(email string) ([]string, error) {
	if valid, err := utils.CheckValidEmails([]string{email}); !valid {
		panic(err)
	}

	rows, err := repo.db.Query("SELECT friend_email FROM public.friends WHERE user_email=$1 AND BLOCKED=$2", email, false)
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
	return friends, nil
}

// 3.
func (repo *repository) FindCommonFriendsByEmails(emails []string) ([]string, error) {
	if valid, err := utils.CheckValidEmails(emails); !valid {
		panic(err)
	}
	sqlStatement := "SELECT friend_email FROM public.friends WHERE"
	for _, email := range emails {
		sqlStatement += " user_email='" + email + "' OR"
	}
	if len(emails) > 1 {
		sqlStatement = sqlStatement[:len(sqlStatement)-3]
	}
	sqlStatement += " AND BLOCKED=false GROUP BY friend_email HAVING COUNT(*) > 1"
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
	return friends, nil
}

// 4.
func (repo *repository) SubscribeFromEmail(req models.SubscribeRequest) (bool, error) {
	if valid, err := utils.CheckValidEmails([]string{req.Requestor, req.Target}); !valid {
		panic(err)
	}
	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec("INSERT INTO public.subscribers(requestor,target) VALUES ($1,$2)", req.Requestor, req.Target)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}
	tx.Commit()

	return true, nil
}

// 5.
func (repo *repository) BlockSubscribeByEmail(req models.BlockSubscribeRequest) (bool, error) {
	if valid, err := utils.CheckValidEmails([]string{req.Requestor, req.Target}); !valid {
		panic(err)
	}
	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		panic(err)
	}

	//suppose A block B:
	if isFriend, _ := repo.hasFriendConnection(req.Requestor, req.Target); isFriend == true {
		//if A and B are friend, A no longer receive notify from B
		_, err := tx.Exec("INSERT INTO public.subscribers(requestor,target,blocked) VALUES ($1,$2,true) ON CONFLICT (requestor,target) DO UPDATE SET blocked = EXCLUDED.blocked", req.Requestor, req.Target)
		if err != nil {
			err1 := tx.Rollback()
			if err1 != nil {
				panic(err1)
			}
			panic(err)
		}

		// if _, err = res.LastInsertId(); err != nil {
		// 	panic(err)
		// }
		tx.Commit()
	} else {
		//if not friend, no new friend connection added
		_, err := tx.Exec("INSERT INTO public.friends(user_email,friend_email,blocked) VALUES ($1,$2,true) ON CONFLICT (user_email,friend_email) DO UPDATE SET blocked = EXCLUDED.blocked", req.Requestor, req.Target)
		if err != nil {
			err1 := tx.Rollback()
			if err1 != nil {
				panic(err1)
			}
			panic(err)
		}
		// if _, err = res.LastInsertId(); err != nil {
		// 	panic(err)
		// }
		tx.Commit()
	}

	return true, nil
}

func (repo *repository) hasFriendConnection(requestor string, target string) (bool, error) {
	if valid, err := utils.CheckValidEmails([]string{requestor, target}); !valid {
		panic(err)
	}
	rows, err := repo.db.Query("SELECT * FROM public.friends WHERE user_email=$1 AND friend_email=$2 AND BLOCKED=false", requestor, target)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if rows.Next() {
		return true, nil
	}
	return false, nil
}

// 6.
func (repo *repository) GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) ([]string, error) {
	if valid, err := utils.CheckValidEmails([]string{req.Sender}); !valid {
		panic(err)
	}

	//if has a friend connection
	rows, err := repo.db.Query("SELECT friend_email FROM public.friends WHERE user_email=$1 AND blocked=$2", req.Sender, false)
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
		rows, err := repo.db.Query("SELECT target FROM public.subscribers WHERE requestor=$1 AND blocked=true", req.Sender)
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
	rows, err = repo.db.Query("SELECT requestor FROM public.subscribers WHERE target=$1 AND blocked=false", req.Sender)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var friend string
		if err := rows.Scan(&friend); err != nil {
			panic(err)
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

	return recipients, nil
}
