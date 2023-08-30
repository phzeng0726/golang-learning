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

// type User struct {
// 	UserId       string     `gorm:"column:user_id;not null;primaryKey;"`
// 	Location     string     `gorm:"column:location;not null;"`
// 	Language     string     `gorm:"column:language;not null;"`
// 	JoinDate     time.Time  `gorm:"column:join_date;not null;"`
// 	ViewTutorial string     `gorm:"column:view_tutorial;not null;"`
// 	CompanyId    *string    `gorm:"column:company_id;"` // 可能為null就要用point
// 	Tag          *string    `gorm:"column:tag;"`
// 	CreatedAt    time.Time  `gorm:"column:created_at;not null;"`
// 	UpdatedAt    *time.Time `gorm:"column:updated_at;default:null;"`
// }

// func (u User) TableName() string {
// 	return "user"
// }

// type UserAccount struct {
// 	UserId      string     `gorm:"column:user_id;not null;primaryKey;"`
// 	Name        string     `gorm:"column:name;not null;"`
// 	Email       string     `gorm:"column:email;not null;"`
// 	LastLoginAt *time.Time `gorm:"column:last_login_at;default:null;"`
// 	CreatedAt   time.Time  `gorm:"column:created_at;not null;"`
// 	UpdatedAt   *time.Time `gorm:"column:updated_at;default:null;"`
// 	DeletedAt   *time.Time `gorm:"column:deleted_at;default:null;"`
// }

// func (u UserAccount) TableName() string {
// 	return "user_account"
// }
