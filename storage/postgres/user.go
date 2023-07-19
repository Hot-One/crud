package postgres

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	"user/models"
)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(req *models.UserCreate) (*models.UserPrimaryKey, error) {
	var (
		id    = uuid.New().String()
		query = `
			INSERT INTO users(id, full_name, login, password)
			VALUES($1, $2, $3, $4);
		`
	)

	_, err := r.db.Exec(query, id, req.FullName, req.Login, req.Password)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &models.UserPrimaryKey{
		Id: id,
	}, nil
}

func (r *userRepo) GetById(req *models.UserPrimaryKey) (*models.User, error) {
	var (
		resp  models.User
		query string
	)

	query = `
		SELECT
			id,
			full_name,
			login,
			password
		FROM users
		WHERE id = $1
		`
	err := r.db.QueryRow(query, req.Id).Scan(
		&resp.Id,
		&resp.FullName,
		&resp.Login,
		&resp.Password,
	)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *userRepo) GetList(req *models.UserGetListRequest) (*models.UserGetListResponse, error) {

	var (
		resp   = &models.UserGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			full_name,
			login,
			password
		FROM users
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND title ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var (
			user models.User
		)
		err := rows.Scan(
			&resp.Count,
			&user.Id,
			&user.FullName,
			&user.Login,
			&user.Password,
		)

		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &user)
	}

	return resp, nil
}

func (r *userRepo) Update(req *models.UserUpdate) (*models.UserPrimaryKey, error) {
	var (
		query = `
			UPDATE users
			SET 
			full_name = $1, 
			login = $2, 
			password= $3 
			WHERE id = $4
		`
	)

	_, err := r.db.Exec(query, req.FullName, req.Login, req.Password, req.Id)
	if err != nil {
		return nil, err
	}
	return &models.UserPrimaryKey{Id: req.Id}, nil
}

func (r *userRepo) Delete(req *models.UserPrimaryKey) error {
	var (
		query = `
			DELETE FROM users WHERE id = $1
		`
	)

	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return err
	}
	return nil
}
