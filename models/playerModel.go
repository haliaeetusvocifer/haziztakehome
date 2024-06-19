package models

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	FirstName      string
	LastName       string
	ProfilePicture string
	// TeamId uint `gorm: "foreignKey: Id"` // to review later as teamID should be linked to teams model
	Team        string
	TeamID      string
	Position    string
	Nationality string
	DateOfBirth time.Time
	Height      float32
	MarketValue float32
}
