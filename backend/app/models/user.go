package models

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Nickname  string `json:"nickname" gorm:"type:varchar(50)"`
	Bio       string `json:"bio" gorm:"type:varchar(255)"`
	CreatedAt string `json:"created_at" gorm:"type:timestamp"`
}
