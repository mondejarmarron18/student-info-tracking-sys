package repositories

import (
	"database/sql"
	"errors"
	"log"
	"server/v1/utils"

	"server/v1/features/permissions/constants"
	"server/v1/features/permissions/domains"
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

func (r *PermissionRepo) CreatePermission(permission *domains.Permission) (*domains.Permission, error) {
	permissionExclude := domains.Permission{}

	row := r.dbConn.QueryRow("INSERT INTO permission (name, description) VALUES ($1, $2) RETURNING id, name, description, created_at, updated_at", permission.Name, permission.Description)

	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.CreatedAt, &permission.UpdatedAt, &permissionExclude.UpdatedAt)

	if err != nil {
		log.Println("Error creating permission:", err)

		return nil, err
	}

	log.Println("Permission created successfully")
	return permission, nil
}

func (r *PermissionRepo) GetPermissions(filter utils.Filter) (*[]domains.Permission, error) {
	permissions := &[]domains.Permission{}

	if !utils.IsValueInList(filter.SortBy, constants.PermissionTableColumns) {
		return nil, errors.New(r.errorMessage.InvalidRequest)
	}

	query := "SELECT * FROM permission WHERE deleted_at IS NULL AND (name ILIKE '%' || $3 || '%' OR description ILIKE '%' || $3 || '%') ORDER BY " + filter.SortBy + " " + filter.SortOrder + " LIMIT $1 OFFSET $2"

	rows, errRows := r.dbConn.Query(query, filter.Limit, filter.Offset, filter.Q)

	if errRows != nil {
		log.Println("Error getting permissions:", errRows)
		return nil, errRows
	}

	for rows.Next() {
		permission := &domains.Permission{}

		err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.CreatedAt, &permission.UpdatedAt)

		if err != nil {
			log.Println("Error scanning permission:", err)
			return nil, errors.New(r.errorMessage.InternalServerError)
		}

		*permissions = append(*permissions, *permission)
	}

	return permissions, nil
}

func (r *PermissionRepo) GetPermissionById(id string) (*domains.Permission, error) {
	permission := &domains.Permission{}

	query := "SELECT * FROM user_account WHERE id = $1 AND deleted_at IS NULL"
	row := r.dbConn.QueryRow(query, id)

	err := row.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.CreatedAt, &permission.UpdatedAt)

	if err != nil {
		log.Println("Error getting permission:", err)
		return nil, errors.New(r.errorMessage.NotFound)
	}

	return permission, nil
}
