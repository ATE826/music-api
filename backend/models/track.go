package models

import "gorm.io/gorm"

type Track struct {
	gorm.Model
	Title    string `json:"title"`
	Artist   string `json:"artist"`
	Genre    string `json:"genre"`
	Duration int    `json:"duration"` // Duration in seconds
}
