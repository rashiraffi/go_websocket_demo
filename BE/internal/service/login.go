package service

import (
	"errors"
	"wsserver/pkg/auth"

	"go.uber.org/zap"
)

func (s *service) Login(email, pass string) (string, error) {
	user := s.model.GetUserByEmail(email)
	if user == nil {
		return "", errors.New("CLIENT_ERROR::User details not found")
	}

	if user.Password != pass {
		return "", errors.New("CLIENT_ERROR::Incorrect password")
	}

	// Generate JWT token
	token, err := auth.GenerateJWTToken(map[string]any{
		"name":  user.Name,
		"email": user.EMail,
	})
	if err != nil {
		zap.S().Warnw("Error generating JWT token", err.Error())
	}

	return token, nil
}
