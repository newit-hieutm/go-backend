package repo

type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}


func (us *UserRepo) GetInfo() string {
	a := "132"
	b := "sfads"
	return a + b
}