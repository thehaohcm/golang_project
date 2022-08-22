package repositories

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"

	"golang_project/api/internal/models"
	"golang_project/api/internal/pkg"

	_ "github.com/lib/pq"
)

type FriendConnectionRepository interface {
	CreateUser(models.CreatingUserRequest) (models.User, error)
	FindFriendsByEmail(models.FriendListRequest) ([]models.Relationship, error)
	FindCommonFriendsByEmails(models.CommonFriendListRequest) ([]models.Relationship, error)
	CreateFriendConnection(friendConnectionRequest models.FriendConnectionRequest) (models.Relationship, error)
	SubscribeFromEmail(req models.SubscribeRequest) (models.Relationship, error)
	BlockSubscribeByEmail(req models.BlockSubscribeRequest) (models.Relationship, error)
	GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) ([]models.Relationship, error)
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

// CreateUser function used to insert data of a new user into user table
// pass a CreatingUserRequest model as parameter
// return a User model and an error type
func (repo *repository) CreateUser(request models.CreatingUserRequest) (models.User, error) {
	if err := pkg.CheckValidEmail(request.Email); err != nil {
		return models.User{}, err
	}

	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		return models.User{}, err
	}
	_, err = tx.Exec("INSERT INTO public.user_account(user_email) VALUES($1)", request.Email)

	if err != nil {
		tx.Rollback()
		return models.User{}, err
	}

	tx.Commit()
	return models.User{Email: request.Email}, nil
}

// CreateFriendConnection function used to insert data of a new friend connection into relationship table
// pass a FriendConnectionRequest model as parameter
// return a Relationship model and an error type
func (repo *repository) CreateFriendConnection(friendConnectionRequest models.FriendConnectionRequest) (models.Relationship, error) {
	//check empty or invalid email format
	if len(friendConnectionRequest.Friends) != 2 {
		return models.Relationship{}, errors.New("Invalid Request")
	}
	if err := pkg.CheckValidEmails([]string{friendConnectionRequest.Friends[0], friendConnectionRequest.Friends[1]}); err != nil {
		return models.Relationship{}, err
	}

	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		return models.Relationship{}, err
	}
	_, err = tx.Exec("INSERT INTO public.relationship(requestor, target, is_friend) VALUES($1,$2,true),($2,$1,true) ON CONFLICT (requestor,target) DO UPDATE SET is_friend = EXCLUDED.is_friend", friendConnectionRequest.Friends[0], friendConnectionRequest.Friends[1])

	if err != nil {
		tx.Rollback()
		return models.Relationship{}, err
	}

	tx.Commit()
	return models.Relationship{Requestor: friendConnectionRequest.Friends[0], Target: friendConnectionRequest.Friends[1], IsFriend: true}, nil
}

// FindFriendsByEmail function used to query data from relationship table to get a list of friend emails by an email address
// pass a FriendListRequest model as parameter
// return an array of Relationship model and an error type
func (repo *repository) FindFriendsByEmail(request models.FriendListRequest) ([]models.Relationship, error) {
	if err := pkg.CheckValidEmail(request.Email); err != nil {
		return []models.Relationship{}, err
	}

	rows, err := repo.db.Query("SELECT requestor FROM public.relationship WHERE target=$1 and is_friend=true AND friend_blocked=false UNION SELECT target FROM public.relationship WHERE requestor=$1 and is_friend=true AND friend_blocked=false", request.Email)
	if err != nil {
		return []models.Relationship{}, err
	}
	defer rows.Close()

	var relationships []models.Relationship
	for rows.Next() {
		relationshipTmp := models.Relationship{Requestor: request.Email}
		if err := rows.Scan(&relationshipTmp.Target); err != nil {
			return []models.Relationship{}, err
		}
		relationships = append(relationships, relationshipTmp)
	}
	return relationships, nil
}

// FindCommonFriendsByEmails function used to query data from relationship table to get a list of common friend emails between 2 email addresses
// pass a CommonFriendListRequest model as parameter
// return an array of Relationship model and an error type
func (repo *repository) FindCommonFriendsByEmails(request models.CommonFriendListRequest) ([]models.Relationship, error) {
	if err := pkg.CheckValidEmails(request.Friends); err != nil {
		return []models.Relationship{}, err
	}

	var dollarSignParams string
	var arg []interface{}
	for i, friend := range request.Friends {
		dollarSignParams += "$" + strconv.Itoa(i+1) + ","
		arg = append(arg, friend)
	}
	dollarSignParams = dollarSignParams[:len(dollarSignParams)-1]

	sqlStatement := "SELECT target,count(*) FROM public.relationship where requestor in (" + dollarSignParams + ") and target not in (" + dollarSignParams + ") group by target having count(*)>1 union SELECT requestor ,count(*) FROM public.relationship where target in (" + dollarSignParams + ") and requestor not in(" + dollarSignParams + ") group by requestor having count(*)>1"
	rows, err := repo.db.Query(sqlStatement, arg...)
	if err != nil {
		return []models.Relationship{}, err
	}
	defer rows.Close()

	var relationships []models.Relationship
	for rows.Next() {
		var relationship models.Relationship
		var count int
		if err := rows.Scan(&relationship.Target, &count); err != nil {
			return []models.Relationship{}, err
		}
		relationships = append(relationships, relationship)
	}
	return relationships, nil
}

// SubscribeFromEmail function used to insert a new subscribe connection into relationship table
// pass a SubscribeRequest model as parameter
// return a Relationship model and an error type
func (repo *repository) SubscribeFromEmail(req models.SubscribeRequest) (models.Relationship, error) {
	if err := pkg.CheckValidEmails([]string{req.Requestor, req.Target}); err != nil {
		return models.Relationship{}, err
	}
	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		return models.Relationship{}, err
	}
	_, err = tx.Exec("INSERT INTO public.relationship(requestor, target, subscribed) VALUES ($1,$2,true) ON CONFLICT (requestor,target) DO UPDATE SET subscribed = EXCLUDED.subscribed", req.Requestor, req.Target)
	if err != nil {
		tx.Rollback()
		return models.Relationship{}, err
	}
	tx.Commit()

	return models.Relationship{Requestor: req.Requestor, Target: req.Target, Subscribed: true}, nil
}

// BlockSubscribeByEmail function used to update data in relationship table to block a subscribe connection
// pass a BlockSubscribeRequest model as parameter
// return a Relationship model and an error type
func (repo *repository) BlockSubscribeByEmail(req models.BlockSubscribeRequest) (models.Relationship, error) {
	if err := pkg.CheckValidEmails([]string{req.Requestor, req.Target}); err != nil {
		return models.Relationship{}, err
	}
	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		return models.Relationship{}, err
	}

	//suppose A block B:
	//if A and B are friend, A no longer receive notify from B
	_, err = tx.Exec("INSERT INTO public.relationship(requestor,target,subscribe_blocked) VALUES ($1,$2,true) ON CONFLICT (requestor,target) DO UPDATE SET subscribe_blocked = EXCLUDED.subscribe_blocked", req.Requestor, req.Target)
	if err != nil {
		tx.Rollback()
		return models.Relationship{}, err
	}

	tx.Commit()

	return models.Relationship{Requestor: req.Requestor, Target: req.Target, FriendBlocked: true}, nil
}

// GetSubscribingEmailListByEmail function used to update data in relationship table to block a subscribe connection
// pass a GetSubscribingEmailListRequest model as parameter
// return an array of Relationship model and an error type
func (repo *repository) GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) ([]models.Relationship, error) {
	if err := pkg.CheckValidEmail(req.Sender); err != nil {
		return []models.Relationship{}, err
	}

	var relationships []models.Relationship

	//has a friend connection
	rows, err := repo.db.Query("SELECT requestor, target, is_friend, friend_blocked, subscribed, subscribe_blocked FROM public.relationship rs WHERE rs.requestor=$1 OR rs.target=$1 AND is_friend=true AND friend_blocked=false", req.Sender)
	if err != nil {
		return []models.Relationship{}, err
	}

	defer rows.Close()
	var friends []string
	for rows.Next() {
		var relationshipTmp models.Relationship
		if err := rows.Scan(&relationshipTmp.Requestor, &relationshipTmp.Target, &relationshipTmp.IsFriend, &relationshipTmp.FriendBlocked, &relationshipTmp.Subscribed, &relationshipTmp.SubscribeBlock); err != nil {
			return []models.Relationship{}, err
		}
		friends = append(friends, relationshipTmp.Requestor, relationshipTmp.Target)
	}

	//if has a friend connection, but blocked in subscribers tables
	rows, err = repo.db.Query("SELECT requestor, target, is_friend, friend_blocked, subscribed, subscribe_blocked FROM public.relationship rs WHERE rs.requestor=$1 OR rs.target=$1 AND is_friend=true AND subscribed=true AND friend_blocked=false AND subscribe_blocked=false", req.Sender)
	if err != nil {
		return []models.Relationship{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var relationshipTmp models.Relationship
		if err := rows.Scan(&relationshipTmp.Requestor, &relationshipTmp.Target, &relationshipTmp.IsFriend, &relationshipTmp.FriendBlocked, &relationshipTmp.Subscribed, &relationshipTmp.SubscribeBlock); err != nil {
			return []models.Relationship{}, err
		}
		friends = append(friends, relationshipTmp.Requestor, relationshipTmp.Target)
	}

	//if subscribed to updates
	rows, err = repo.db.Query("SELECT requestor, target, is_friend, friend_blocked, subscribed, subscribe_blocked FROM public.relationship rs WHERE rs.target=$1 AND subscribed=true AND subscribe_blocked=false", req.Sender)
	if err != nil {
		return []models.Relationship{}, err
	}

	defer rows.Close()
	for rows.Next() {
		var relationshipTmp models.Relationship
		if err := rows.Scan(&relationshipTmp.Requestor, &relationshipTmp.Target, &relationshipTmp.IsFriend, &relationshipTmp.FriendBlocked, &relationshipTmp.Subscribed, &relationshipTmp.SubscribeBlock); err != nil {
			return []models.Relationship{}, err
		}
		friends = append(friends, relationshipTmp.Requestor)
	}

	//if being mentioned in the update
	textArr := strings.Split(req.Text, " ")
	for _, text := range textArr {
		if pkg.IsEmailValid(text) {
			relationships = append(relationships, models.Relationship{Target: text})
		}
	}

	//remove duplicated items
	friends = pkg.RemoveDuplicatedItems(friends)
	//remove requestor itself
	friends = pkg.RemoveItemInArray(friends, req.Sender)

	for _, item := range friends {
		relationships = append(relationships, models.Relationship{Target: item})
	}

	return relationships, nil
}
