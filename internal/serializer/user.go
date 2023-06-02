package serializer

import (
	"gin_api_frame/internal/model"
)

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	Avatar    string `json:"avatar"`
	LastLogin int64
	Address   string      `json:"address"`
	Biography string      `json:"biography"`
	Phone     string      `json:"phone"`
	Location  model.Point `json:"location"`
	Extra     string      `json:"extra"`
}

// BuildUser 序列化用户
func BuildUser(user *model.User) *User {
	u := &User{
		ID:        user.ID,
		UserName:  user.UserName,
		Email:     user.Email,
		Status:    user.Status,
		LastLogin: user.LastLogin,
		Avatar:    user.Avatar,
		Address:   user.Address,
		Biography: user.Address,
		Phone:     user.Phone,
		Location:  user.Location,
		Extra:     user.Extra,
	}

	return u
}

func BuildUsers(items []*model.User) (users []*User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
