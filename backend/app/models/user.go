package models

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"type:varchar(255)"`
	Password  string `json:"-" gorm:"type:varchar(255)"`
	Nickname  string `json:"nickname" gorm:"type:varchar(50)"`
	Bio       string `json:"bio" gorm:"type:varchar(255)"`
	Token     string `json:"-" gorm:"type:varchar(255)"`
	CreatedAt string `json:"created_at" gorm:"type:timestamp"`
}
