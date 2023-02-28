package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Stogage interface {
	CreateUser(*User) error
	UpdateUser(*User) error
	getUserByPN(int) (*User, error)
	getUsers() ([]*User, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=bonik sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) CreateUser(u *User) error {

	var dbname string
	if err := s.db.QueryRow("select name from users where pers_number = $1", u.PersonalNumber).Scan(&dbname); err != nil {
		if err == sql.ErrNoRows {
			q := `insert into users (name, pers_number, user_schema, create_at)	values ($1, $2, $3, $4)`
			_, err = s.db.Query(q, u.Name, u.PersonalNumber, u.UserSchema, u.CreatedAt)
			if err != nil {
				return err
			}
		}
		return err
	}

	return nil
}

func (s *PostgresStore) UpdateUser(*User) error {
	return nil
}

func (s *PostgresStore) getUserByPN(pn int) (*User, error) {
	user := new(User)
	if err := s.db.QueryRow("select id, name, pers_number, user_schema, create_at from users where pers_number = $1", pn).Scan(
		&user.ID,
		&user.Name,
		&user.PersonalNumber,
		&user.UserSchema,
		&user.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *PostgresStore) getUsers() ([]*User, error) {
	rows, err := s.db.Query(`select * from users`)
	if err != nil {
		return nil, err
	}
	users := []*User{}
	for rows.Next() {
		user := new(User)
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.PersonalNumber,
			&user.UserSchema,
			&user.CreatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func (s *PostgresStore) Setup() error {
	return s.CreateTableUser()
}

func (s *PostgresStore) CreateTableUser() error {
	q := `create table if not exists users (
		id serial primary key,
		name varchar(255),
		pers_number integer,
		user_schema varchar(2),
		create_at timestamp
	)`
	_, err := s.db.Exec(q)
	return err
}
