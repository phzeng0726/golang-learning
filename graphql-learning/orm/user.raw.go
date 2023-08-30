package orm

type UserData struct {
	UserId     string             `gorm:"column:user_id;" json:"userId"`
	Name       string             `gorm:"column:name;" json:"name"`
	Email      string             `gorm:"column:email;" json:"email"`
	TotalPrice int                `gorm:"column:total_price;" json:"totalPrice"`
	Purchases  []UserPurchaseData `gorm:"-" json:"purchases"`
}

type UserPurchaseData struct {
	UserId   string `gorm:"column:user_id;"`
	Quantity int    `gorm:"column:quantity;" json:"quantity"`
	Category string `gorm:"column:category;" json:"category"`
	Name     string `gorm:"column:name;" json:"name"`
	Price    int    `gorm:"column:price;" json:"price"`
	SubTotal int    `gorm:"column:sub_total;" json:"subTotal"`
}
