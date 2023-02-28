package main

import (
	"time"
)

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	PersonalNumber int       `json:"personalNumber"`
	UserSchema     string    `json:"userSchema"`
	CreatedAt      time.Time `json:"createdAt"`
}

type CreateUserRequest struct {
	Name           string `json:"name"`
	UserSchema     string `json:"userSchema"`
	PersonalNumber int    `json:"personalNumber"`
}

func NewUser(name, uschema string, pn int) *User {
	if uschema == "" {
		uschema = "33"
	}
	return &User{
		//	ID:             rand.Intn(999999),
		Name:           name,
		UserSchema:     uschema,
		PersonalNumber: pn,
		CreatedAt:      time.Now().UTC(),
	}
}

type BoardingPass struct {
	Fmt      string `json:"fmt"`
	Name     string `json:"name"`
	Booking  string `json:"booking"`
	JDate    string `json:"jdate"`
	TypePass string `json:"tp"`   //Y- эконом J-бизнес F-первый
	Zone     int    `json:"zone"` //0 - общая. 1-чистая
	Checkin1 int    `json:"ch1"`
	Checkin2 int    `json:"ch2"`
}

func NewBoardingPass(name, booking, jdate, tp string) *BoardingPass {
	return &BoardingPass{
		Fmt:      "M1",
		Name:     name,
		Booking:  booking,
		JDate:    jdate,
		TypePass: tp,
		Zone:     0,
		Checkin1: 0,
		Checkin2: 0,
	}
}
