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
	"github.com/stretchr/testify/mock"
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

// 1. Create user
func (repo *repository) CreateUser(request models.CreatingUserRequest) (models.User, error) {
	if valid, err := pkg.CheckValidEmail(request.Email); !valid {
		return models.User{}, err
	}

	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		return models.User{}, err
	}
	_, err = tx.Exec("INSERT INTO public.user_account(user_email) VALUES($1)", request.Email)

	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			panic(err1)
		}
		return models.User{}, err
	}

	tx.Commit()
	return models.User{Email: request.Email}, nil
}

// 1.
func (repo *repository) CreateFriendConnection(friendConnectionRequest models.FriendConnectionRequest) (models.Relationship, error) {

	//check empty or invalid email format
	if len(friendConnectionRequest.Friends) != 2 {
		return models.Relationship{}, errors.New("invalid request")
	}
	if valid, err := pkg.CheckValidEmails([]string{friendConnectionRequest.Friends[0], friendConnectionRequest.Friends[1]}); !valid {
		return models.Relationship{}, err
	}

	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		return models.Relationship{}, err
	}
	_, err = tx.Exec("INSERT INTO public.relationship(requestor, target, is_friend) VALUES($1,$2,true),($2,$1,true) ON CONFLICT (requestor,target) DO UPDATE SET is_friend = EXCLUDED.is_friend", friendConnectionRequest.Friends[0], friendConnectionRequest.Friends[1])

	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			panic(err1)
		}
		return models.Relationship{}, err
	}

	tx.Commit()
	return models.Relationship{Requestor: friendConnectionRequest.Friends[0], Target: friendConnectionRequest.Friends[1], Is_friend: true}, nil
}

// 2.
func (repo *repository) FindFriendsByEmail(request models.FriendListRequest) ([]models.Relationship, error) {
	if valid, err := pkg.CheckValidEmail(request.Email); !valid {
		return []models.Relationship{}, err
	}

	rows, err := repo.db.Query("SELECT requestor FROM public.relationship WHERE target=$1 and is_friend=true AND FRIEND_BLOCKED=false UNION SELECT target FROM public.relationship WHERE requestor=$1 and is_friend=true AND FRIEND_BLOCKED=false", request.Email)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var relationships []models.Relationship
	for rows.Next() {
		relationshipTmp := models.Relationship{Requestor: request.Email}
		if err := rows.Scan(&relationshipTmp.Target); err != nil {
			panic(err)
		}
		relationships = append(relationships, relationshipTmp)
	}
	return relationships, nil
}

// 3.
func (repo *repository) FindCommonFriendsByEmails(request models.CommonFriendListRequest) ([]models.Relationship, error) {
	if valid, err := pkg.CheckValidEmails(request.Friends); !valid {
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
		panic(err)
	}
	defer rows.Close()

	var relationships []models.Relationship
	for rows.Next() {
		var relationship models.Relationship
		var count int
		if err := rows.Scan(&relationship.Target, &count); err != nil {
			panic(err)
		}
		relationships = append(relationships, relationship)
	}
	return relationships, nil
}

// 4.
func (repo *repository) SubscribeFromEmail(req models.SubscribeRequest) (models.Relationship, error) {
	if valid, err := pkg.CheckValidEmails([]string{req.Requestor, req.Target}); !valid {
		panic(err)
	}
	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		panic(err)
	}
	_, err = tx.Exec("INSERT INTO public.relationship(requestor, target, subscribed) VALUES ($1,$2,true) ON CONFLICT (requestor,target) DO UPDATE SET subscribed = EXCLUDED.subscribed", req.Requestor, req.Target)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}
	tx.Commit()

	return models.Relationship{Requestor: req.Requestor, Target: req.Target, Subscribed: true}, nil
}

// 5.
func (repo *repository) BlockSubscribeByEmail(req models.BlockSubscribeRequest) (models.Relationship, error) {
	if valid, err := pkg.CheckValidEmails([]string{req.Requestor, req.Target}); !valid {
		panic(err)
	}
	tx, err := repo.db.BeginTx(repo.ctx, nil)
	if err != nil {
		panic(err)
	}

	//suppose A block B:
	//if A and B are friend, A no longer receive notify from B
	_, err = tx.Exec("INSERT INTO public.relationship(requestor,target,subscribe_blocked) VALUES ($1,$2,true) ON CONFLICT (requestor,target) DO UPDATE SET subscribe_blocked = EXCLUDED.subscribe_blocked", req.Requestor, req.Target)
	if err != nil {
		err1 := tx.Rollback()
		if err1 != nil {
			panic(err1)
		}
		panic(err)
	}

	tx.Commit()

	return models.Relationship{Requestor: req.Requestor, Target: req.Target, Friend_blocked: true}, nil
}

// 6.
func (repo *repository) GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) ([]models.Relationship, error) {
	if valid, err := pkg.CheckValidEmail(req.Sender); !valid {
		panic(err)
	}

	var relationships []models.Relationship

	//has a friend connection
	rows, err := repo.db.Query("select requestor, target, is_friend, friend_blocked, subscribed, subscribe_blocked from public.relationship rs where rs.requestor=$1 or rs.target=$1 and is_friend=true and friend_blocked=false", req.Sender)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var friends []string
	for rows.Next() {
		var relationshipTmp models.Relationship
		if err := rows.Scan(&relationshipTmp.Requestor, &relationshipTmp.Target, &relationshipTmp.Is_friend, &relationshipTmp.Friend_blocked, &relationshipTmp.Subscribed, &relationshipTmp.Subscribe_block); err != nil {
			panic(err)
		}
		friends = append(friends, relationshipTmp.Requestor, relationshipTmp.Target)
	}

	//if has a friend connection, but blocked in subscribers tables
	rows, err = repo.db.Query("select requestor, target, is_friend, friend_blocked, subscribed, subscribe_blocked from public.relationship rs where rs.requestor=$1 or rs.target=$1 and is_friend=true and subscribed=true and friend_blocked=false and subscribe_blocked=true", req.Sender)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var relationshipTmp models.Relationship
		if err := rows.Scan(&relationshipTmp.Requestor, &relationshipTmp.Target, &relationshipTmp.Is_friend, &relationshipTmp.Friend_blocked, &relationshipTmp.Subscribed, &relationshipTmp.Subscribe_block); err != nil {
			panic(err)
		}
		friends = append(friends, relationshipTmp.Requestor, relationshipTmp.Target)
	}

	//if subscribed to updates
	rows, err = repo.db.Query("select requestor, target, is_friend, friend_blocked, subscribed, subscribe_blocked from public.relationship rs where rs.target=$1 and subscribed=true and subscribe_blocked=false", req.Sender)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var relationshipTmp models.Relationship
		if err := rows.Scan(&relationshipTmp.Requestor, &relationshipTmp.Target, &relationshipTmp.Is_friend, &relationshipTmp.Friend_blocked, &relationshipTmp.Subscribed, &relationshipTmp.Subscribe_block); err != nil {
			panic(err)
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
	for _, item := range friends {
		relationships = append(relationships, models.Relationship{Target: item})
	}

	return relationships, nil
}

type FriendConnectionRepoMock struct {
	mock.Mock
}

func (f *FriendConnectionRepoMock) CreateUser(request models.CreatingUserRequest) (models.User, error) {
	if valid, err := pkg.CheckValidEmail(request.Email); !valid || err != nil {
		return models.User{}, err
	}
	return models.User{Email: request.Email}, nil
}

func (f *FriendConnectionRepoMock) FindFriendsByEmail(models.FriendListRequest) ([]models.Relationship, error) {
	return []models.Relationship{}, nil
}
func (f *FriendConnectionRepoMock) FindCommonFriendsByEmails(request models.CommonFriendListRequest) ([]models.Relationship, error) {
	return []models.Relationship{}, nil
}
func (f *FriendConnectionRepoMock) CreateFriendConnection(request models.FriendConnectionRequest) (models.Relationship, error) {
	if valid, err := pkg.CheckValidEmails(request.Friends); !valid || err != nil {
		return models.Relationship{}, nil
	}
	return models.Relationship{}, nil
}
func (f *FriendConnectionRepoMock) SubscribeFromEmail(req models.SubscribeRequest) (models.Relationship, error) {
	if len(req.Requestor) > 0 && len(req.Target) > 0 {
		return models.Relationship{}, nil
	}
	return models.Relationship{}, nil
}
func (f *FriendConnectionRepoMock) BlockSubscribeByEmail(req models.BlockSubscribeRequest) (models.Relationship, error) {
	if len(req.Requestor) > 0 && len(req.Target) > 0 {
		return models.Relationship{}, nil
	}
	return models.Relationship{}, nil
}
func (f *FriendConnectionRepoMock) GetSubscribingEmailListByEmail(req models.GetSubscribingEmailListRequest) ([]models.Relationship, error) {
	return []models.Relationship{}, nil
}
