package models

type Bid struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	UserID    uint   `json:"user_id"`
	AuctionID uint   `json:"auction_id" gorm:"index"`
	Price     uint   `json:"price"`
	CreatedAt string `json:"created_at" gorm:"type:timestamp"`

	User User `json:"user" gorm:"foreignKey:user_id"`
}
