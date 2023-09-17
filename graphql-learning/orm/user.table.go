package orm

import "time"

type User struct {
	UserId    string    `gorm:"column:user_id;not null;primaryKey;"`
	Name      string    `gorm:"column:name;not null;"`
	Email     string    `gorm:"column:email;not null;"`
	CreatedAt time.Time `gorm:"column:created_at;not null;"`
}

func (u User) TableName() string {
	return "user"
}
