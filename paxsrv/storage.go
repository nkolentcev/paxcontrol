package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type Stogage interface {
	CreateUser(*User) error
	UpdateUser(*User) error
	getUserByPN(int) (*User, error)
	getUsers() ([]*User, error)
	getBoardings() ([]*BoardingPass, error)
	CreateBoardinPass(*BoardingPass) error
	getBoardingPassByBooking(string) (*BoardingPass, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "port=6599 user=postgres dbname=postgres password=bonik sslmode=disable"
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

func (s *PostgresStore) CreateBoardinPass(bp *BoardingPass) error {
	var dbbooking string
	if err := s.db.QueryRow("select booking from boardpass where booking like $1", bp.Booking).Scan(&dbbooking); err != nil {
		if err == sql.ErrNoRows {

			q := `insert into boardpass (fmt, name, booking, jdate, type_pass, zone, create_at, check1, check2) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
			_, err := s.db.Query(q, "M1", bp.Name, bp.Booking, bp.JDate, bp.TypePass, bp.Zone, time.Now().UTC(), 0, 0)
			if err != nil {
				return err
			}
		}
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

func (s *PostgresStore) getBoardingPassByBooking(bk string) (*BoardingPass, error) {

	bpass := new(BoardingPass)
	if err := s.db.QueryRow("select id, fmt, name, booking, jdate, type_pass, zone, check1, check2, create_at from boardpass where booking like $1", bk).Scan(
		&bpass.ID,
		&bpass.Fmt,
		&bpass.Name,
		&bpass.Booking,
		&bpass.JDate,
		&bpass.TypePass,
		&bpass.Zone,
		&bpass.Check1,
		&bpass.Check2,
		&bpass.CreatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return bpass, err
		}
		if err != nil {
			return nil, err
		}
	}
	return bpass, nil

}

func (s *PostgresStore) getBoardings() ([]*BoardingPass, error) {
	boardings := []*BoardingPass{}
	return boardings, nil
}

func (s *PostgresStore) Setup() error {
	err := s.CreateTableUser()
	if err != nil {
		return err
	}
	err = s.CreateTableBoardPass()
	if err != nil {
		return err
	}
	return nil
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

func (s *PostgresStore) CreateTableBoardPass() error {
	q := `create table if not exists boardpass (
		id serial primary key,
		fmt varchar(2),
		name varchar(255),
		booking varchar(7),
		jdate varchar(3),
		type_pass varchar(1),
		zone integer,
		check1 integer,
		check2 integer,
		create_at timestamp
	)`
	_, err := s.db.Exec(q)
	return err
}
