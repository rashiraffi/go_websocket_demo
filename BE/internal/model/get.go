package model

import "wsserver/internal/entities"

func (m *model) GetUserByEmail(emailID string) *entities.User {
	user, ok := m.userList[emailID]
	if !ok {
		return nil
	}
	return user
}
