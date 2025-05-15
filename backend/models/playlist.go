package models

import "gorm.io/gorm"

type Playlist struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Tracks      []Track `gorm:"many2many:playlist_tracks;"`
}
