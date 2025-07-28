package models

import "gorm.io/gorm"

// todo model struct untuk komunikasi database
type Todo struct {
	gorm.Model
	Name        string
	Description string
}
