package service

import (
	"strconv"
	"GoWebapitest/internal/core/port"
)

type ResetPasswordService struct {
	repo port.ResetPasswordRepository
}

func NewResetPasswordService(r port.ResetPasswordRepository) *ResetPasswordService {
	return &ResetPasswordService{repo: r}
}


func (s *ResetPasswordService) ResetPassword(useridStr string, password string) error {
	userid, err := strconv.Atoi(useridStr)
	if err != nil {
		return err
	}

	return s.repo.ResetPassword(userid, password)
}