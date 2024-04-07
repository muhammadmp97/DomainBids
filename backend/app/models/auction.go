package models

type Auction struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	SLD           string `json:"sld"`
	TLD           string `json:"tld"`
	UserID        uint   `json:"user_id" gorm:"index"`
	StartingPrice uint   `json:"starting_price"`
	Description   string `json:"description"`
	Status        string `json:"status"`
	EndsAt        string `json:"ends_at" gorm:"type:timestamp"`
	CreatedAt     string `json:"created_at" gorm:"type:timestamp"`

	Bids []Bid `json:"bids" gorm:"foreignKey:AuctionID"`
}
