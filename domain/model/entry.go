package model

import (
	"time"
)

type Archivo struct {
	ID        int        `gorm:"primary_key" json:"id"`
	Userid    int	     `gorm:"not null" json:"userid" binding:"required"`
	Name     string      `gorm:"not null" json:"name" binding:"required"`
	Link     string      `gorm:"not null" json:"link" binding:"required"`
	Chatid     string    `gorm:"not null" json:"chatid" binding:"required"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	// When you get an Entry Model with gorm, you can get associated tags using junction table automatically.
	// When you write or update an Entry Model, you have to include　tags id. (API Requests don't have tags id.)
	Tags []*Tag `gorm:"many2many:entries_tags;association_autoupdate:false;association_autocreate:false;" json:"tags"`
}
