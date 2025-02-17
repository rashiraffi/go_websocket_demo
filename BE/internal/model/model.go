package model

import "wsserver/internal/entities"

type Model interface {
	GetUserByEmail(emailID string) *entities.User
}

type model struct {
	userList map[string]*entities.User
}

func New() Model {
	return &model{
		userList: userList,
	}
}
