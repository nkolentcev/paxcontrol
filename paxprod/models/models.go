package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name           string `json:"name" gorm:"text;not null;default:null"`
	PersonalNumber int    `json:"personal_number" gorm:"int; unique; not null;default:null"`
	ZoneType       int    `json:"zone_type" gorm:"int:not null;default:0"` // 0 - общая 1 - чистая 2 - обе
	Boarding       int    `json:"boarding" gorm:"int:not null;default:0"`  // 0 - нет 1 - да - фича возможность выбора рейса регистрации посадки
}

type BoardingPass struct {
	gorm.Model
	Name       string `json:"name" gorm:"text;not null;default:null"`
	Booking    string `json:"booking" gorm:"text;not null;default:null"`
	FlightSpec string `json:"fligt_spec" gorm:"text;not null;default:null"`
	FlightNum  string `json:"flight_number" gorm:"text;not null;default:null"`
	JDate      string `json:"jdate" gorm:"text;not null;default:null"`
	TypePass   string `json:"type_pass" gorm:"text;not null;default:null"`
	Zone       int    `json:"zone" gorm:"int:not null;default:0"`
	Check1     int    `json:"check1" gorm:"int:not null;default:0"`
	Check2     int    `json:"check2" gorm:"int:not null;default:0"`
}
