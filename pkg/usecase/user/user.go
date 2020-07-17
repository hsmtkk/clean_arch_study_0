package user

import "github.com/hsmtkk/clean_arch_study_0/pkg/domain/user"

type UserUseCase interface {
	Add(user user.User) error
	Users() ([]user.User, error)
	UserByID(id int) (user.User, error)
}

type userUseCaseImpl struct {
	userRepo UserRepository
}

func New(userRepo UserRepository) UserUseCase {
	return &userUseCaseImpl{}
}

func (userUseCase *userUseCaseImpl) Add(user user.User) error {
	_, err := userUseCase.userRepo.Store(user)
	return err
}

func (userUseCase *userUseCaseImpl) Users() ([]user.User, error) {
	return userUseCase.userRepo.FindAll()
}

func (userUseCase *userUseCaseImpl) UserByID(id int) (user.User, error) {
	return userUseCase.userRepo.FindByID(id)
}

type UserRepository interface {
	Store(user.User) (int, error)
	FindByID(int) (user.User, error)
	FindAll() ([]user.User, error)
}
