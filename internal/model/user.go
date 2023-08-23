package model

type User struct {
	Id       uint   `gorm:"primarykey"`
	Uid      string `gorm:"unique;not null"`
	Username string `gorm:"unique;not null"`
	Nickname string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"not null"`
	// IsDeleted string `gorm:"not null"'`
	// CreatedAt time.Time
	// UpdatedAt time.Time
}

func (u *User) TableName() string {
	return "users"
}
