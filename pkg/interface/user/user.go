package user

import (
	"github.com/hsmtkk/clean_arch_study_0/pkg/interface/database"
	"github.com/hsmtkk/clean_arch_study_0/pkg/usecase/user"
	"github.com/labstack/echo"
)

type UserController interface {
	Create(ctx echo.Context) error
	Index(ctx echo.Context) error
	Show(ctx echo.Context) error
}

type userControllerImpl struct {
	userUseCase user.UserUseCase
	sqlHandler database.SQLHandler
}

func New(sqlHandler database.SQLHandler) UserController {
	repo := 
	userUseCase := user.New()
	return &userControllerImpl{sqlHandler: sqlHandler}
}

func (userController *userControllerImpl) Create(ctx echo.Context) error {
	return nil
}

func (userController *userControllerImpl) Index(ctx echo.Context) error {
	return nil
}

func (userController *userControllerImpl) Show(ctx echo.Context) error {
	return nil
}
