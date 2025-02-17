package model

import "wsserver/internal/entities"

var (
	// Users is the list of users
	userList = map[string]*entities.User{
		"user1@test.com": {
			ID:       1,
			Name:     "User 1",
			EMail:    "user1@test.com",
			Password: "user1psd",
		},
		"user2@test.com": {
			ID:       2,
			Name:     "User 2",
			EMail:    "user2@test.com",
			Password: "user2psd",
		},
	}
)
