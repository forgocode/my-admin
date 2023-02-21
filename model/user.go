package model

type User struct {
	UUID     string `json:"uuid" gorm:"index;comment:the uuid of user"`
	UserName string
	Password string
	NickName string
	Theme    string ``
	Avatar   string
	Role     string
	RoleID   string
	Phone    string
	Email    string
	Status   int
}

func (User) TableName() string {
	return "users"
}
