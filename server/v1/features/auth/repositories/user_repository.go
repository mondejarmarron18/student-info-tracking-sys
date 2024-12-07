package repositories

import (
	"database/sql"
	"log"
	"server/v1/utils"
	"time"

	"server/v1/features/users/domains"
)

type UserRepo struct {
	dbConn *sql.DB
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		dbConn: utils.GetConn(),
	}
}

func (r *UserRepo) CreateUser(user *domains.User) (*domains.User, error) {
	currentTime := time.Now()

	row := r.dbConn.QueryRow("INSERT INTO user_account (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, email, password, created_at, updated_at", user.Email, user.Password, currentTime, currentTime)

	err := row.Scan(&user.Id, &user.Email, nil, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Println("Error creating user:", err)

		return nil, err
	}

	log.Println("User created successfully")
	return user, nil
}

func (r *UserRepo) GetUers(filter utils.Filter) (*[]domains.User, error) {
	users := &[]domains.User{}

	if filter.Limit == 0 {
		filter.Limit = 10
	}

	if filter.Offset == 0 {
		filter.Offset = 0
	}

	query := "SELECT * FROM user_account LIMIT $1 OFFSET $2"
	rows, errRows := r.dbConn.Query(query, filter.Limit, filter.Offset)

	if errRows != nil {
		log.Println("Error getting users:", errRows)
		return nil, errRows
	}

	for rows.Next() {
		user := &domains.User{}
		userExclude := domains.User{}
		err := rows.Scan(&user.Id, &user.Email, &userExclude.Password, &user.VerifiedAt, &user.DeletedAt, &user.CreatedAt, &user.UpdatedAt)

		if err != nil {
			log.Println("Error scanning user:", err)
			return nil, err
		}

		*users = append(*users, *user)
	}

	return users, nil
}

func (r *UserRepo) GetUserById(id string) (*domains.User, error) {
	user := &domains.User{}
	userExclude := domains.User{}
	query := "SELECT * FROM user_account WHERE id = $1"
	row := r.dbConn.QueryRow(query, id)

	err := row.Scan(&user.Id, &user.Email, &userExclude.Password, &user.VerifiedAt, &user.DeletedAt, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Println("Error getting user:", err)
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) GetUserByEmail(email string) (*domains.User, error) {
	user := &domains.User{}
	userExclude := domains.User{}
	query := "SELECT * FROM user_account WHERE email = $1"
	row := r.dbConn.QueryRow(query, email)

	err := row.Scan(&user.Id, &user.Email, &userExclude.Password, &user.VerifiedAt, &user.DeletedAt, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Println("Error getting user:", err)
		return nil, err
	}

	return user, nil
}

func (r *UserRepo) UpdateUser(user *domains.User) error {
	query := "UPDATE user_account SET email = $1, password = $2, updated_at = $3 WHERE id = $4"

	_, err := r.dbConn.Exec(query, user.Email, user.Password, time.Now(), user.Id)

	if err != nil {
		log.Println("Error updating user:", err)
		return err
	}

	return nil
}

func (r *UserRepo) DeleteUser(id string, isHardDelete bool) error {

	var err error

	if isHardDelete {
		query := "DELETE FROM user_account WHERE id = $1"
		_, err = r.dbConn.Exec(query, id)
	} else {
		query := "UPDATE user_account SET deleted_at = $1 WHERE id = $1"
		_, err = r.dbConn.Exec(query, time.Now(), id)
	}

	if err != nil {
		log.Println("Error deleting user:", err)
		return err
	}

	return nil
}

func (r *UserRepo) IsEmailExist(email string) (bool, error) {
	query := "SELECT EXISTS(SELECT 1 FROM user_account WHERE email = $1)"

	row := r.dbConn.QueryRow(query, email)

	var isExist bool
	err := row.Scan(&isExist)

	if err != nil {
		log.Println("Error checking email:", err)
		return false, err
	}

	return isExist, nil
}
