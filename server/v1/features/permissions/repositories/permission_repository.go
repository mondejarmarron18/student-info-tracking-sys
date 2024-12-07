package repositories

import (
	"database/sql"
	"errors"
	"log"
	"server/v1/utils"

	"server/v1/features/users/constants"
	"server/v1/features/users/domains"
)

type PermissionRepo struct {
	dbConn       *sql.DB
	errorMessage utils.ErrorMessage
}

func NewPermissionRepo() *PermissionRepo {
	return &PermissionRepo{
		dbConn:       utils.GetConn(),
		errorMessage: utils.NewErrorMessage(),
	}
}

func (r *PermissionRepo) CreatePermission(user *domains.User) (*domains.User, error) {
	userExclude := domains.User{}

	row := r.dbConn.QueryRow("INSERT INTO permission () VALUES ($1, $2, $3, $4) RETURNING id, email, password, created_at, updated_at", user.Email, user.Password)

	err := row.Scan(&user.Id, &user.Email, &userExclude.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Println("Error creating user:", err)

		return nil, err
	}

	log.Println("User created successfully")
	return user, nil
}

func (r *PermissionRepo) GetPermissions(filter utils.Filter) (*[]domains.User, error) {
	users := &[]domains.User{}

	if !utils.IsValueInList(filter.SortBy, constants.UserTableColumns) {
		return nil, errors.New(r.errorMessage.InvalidRequest)
	}

	query := "SELECT * FROM user_account WHERE email ILIKE '%' || $3 || '%' ORDER BY " + filter.SortBy + " " + filter.SortOrder + " LIMIT $1 OFFSET $2"

	rows, errRows := r.dbConn.Query(query, filter.Limit, filter.Offset, filter.Q)

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
			return nil, errors.New(r.errorMessage.InternalServerError)
		}

		*users = append(*users, *user)
	}

	return users, nil
}

func (r *PermissionRepo) GetPermissionById(id string) (*domains.User, error) {
	user := &domains.User{}
	userExclude := domains.User{}
	query := "SELECT * FROM user_account WHERE id = $1"
	row := r.dbConn.QueryRow(query, id)

	err := row.Scan(&user.Id, &user.Email, &userExclude.Password, &user.VerifiedAt, &user.DeletedAt, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		log.Println("Error getting user:", err)
		return nil, errors.New(r.errorMessage.NotFound)
	}

	return user, nil
}
