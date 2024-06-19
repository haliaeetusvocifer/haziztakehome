package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name    string
	Logo    string
	Sport   string
	League  string
	Founded string
	Stadium string
	Coach   string
}
