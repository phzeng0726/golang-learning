package orm

import "time"

type UserPurchase struct {
	Id        int64     `gorm:"column:id;not null;primaryKey;autoIncrement;"`
	UserId    string    `gorm:"column:user_id;not null;"`
	MenuId    int64     `gorm:"column:menu_id;not null;"`
	Menu      Menu      `gorm:"foreignKey:MenuId"`
	Quantity  int       `gorm:"column:quantity;not null;"`
	CreatedAt time.Time `gorm:"column:created_at;not null;"`
}

func (u UserPurchase) TableName() string {
	return "user_purchase"
}
