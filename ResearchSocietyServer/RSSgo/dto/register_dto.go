package dto

import "hkc.ink/rss/model"

type RegisterModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rm RegisterModel) ToUserModel() *model.User {
	return &model.User{Email: rm.Email, Password: rm.Password}
}
