package model

type User struct {
	UUID     string `json:"uuid" gorm:"index;comment:the uuid of user"`
	UserName string `json:"userName" gorm:"userName;comment: the name of user"`
	Password string `json:"password" gorm:"password; comment: the password of user"`
	NickName string `json:"nickName" gorm:"nickName; comment: the nickName of user"`
	Theme    string `json:"theme" gorm:"theme; comment: the theme of user"`
	Avatar   string `json:"avatar" gorm:"theme; comment: the avator of user"`
	Role     string `json:"role" gorm:"role; comment: the role of user"`
	RoleID   string `json:"roleID" gorm:"roleID; comment: the roleID of user"`
	Phone    string `json:"phone" gorm:"phone; comment: the phone of user"`
	Email    string `json:"email" gorm:"email; comment: the email of user"`
	Status   int    `json:"status" gorm:"email; comment: the status of user"`
}

func (User) TableName() string {
	return "users"
}
