package services

import "github.com/newit-hieutm/go-backend/internal/repo"

type UserService struct{
	userRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repo.NewUserRepo(),
	}
}


func (us *UserService) GetUserInfo() string {
	return us.userRepo.GetInfo()
}