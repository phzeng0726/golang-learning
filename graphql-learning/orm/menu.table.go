package orm

type Menu struct {
	Id       int64  `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	Category string `gorm:"column:category;not null;" json:"category"`
	Name     string `gorm:"column:name;not null;" json:"name"`
	Price    int    `gorm:"column:price;not null;" json:"price"`
}

func (u Menu) TableName() string {
	return "menu"
}
