package model

// Url represents a URL entity.
type Url struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	LongUrl  string `json:"long_url"`
	ShortUrl string `json:"short_url"`
}

// TableName specifies the table name for the Url model in the database.
func (Url) TableName() string { return "urls" }
