package repo

type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}


func (us *UserRepo) GetInfo() string {
	return "Hieu"
}