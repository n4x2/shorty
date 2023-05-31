package model

import "gorm.io/gorm"

type Url struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
	gorm.Model
}

func (Url) TableName() string { return "urls" }
