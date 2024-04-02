package models

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Nickname  string `json:"nickname"`
	Bio       string `json:"bio"`
	CreatedAt string `json:"created_at"`
}
