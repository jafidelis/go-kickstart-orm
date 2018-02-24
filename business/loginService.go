package business

import (
	"errors"

	"github.com/go-kickstart-orm/auth"
	"github.com/go-kickstart-orm/model/entity"
	"github.com/go-kickstart-orm/model/repository"
	"github.com/go-kickstart-orm/util"
)

type LoginService struct {
	userRepository repository.UserRepository
}

func (ls LoginService) Login(user *entity.User) (string, error) {
	if user.Login == "" {
		return "", errors.New("Login not informed")
	}

	if user.Password == "" {
		return "", errors.New("Password not informed")
	}

	userDb := ls.userRepository.FindByLogin(user.Login)

	if (entity.User{}) == userDb {
		return "", errors.New("User not registered")
	}

	var err error
	user.Password, err = util.EncriptyMd5(user.Password)
	if err != nil {
		util.CheckErr(err)
		return "", errors.New("Error validating login")
	}

	if userDb.Password != user.Password {
		return "", errors.New("Password does not match")
	}

	token := auth.GenerateJWT(*user)

	return token, nil
}
