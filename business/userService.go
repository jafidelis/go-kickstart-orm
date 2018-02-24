package business

import (
	"fmt"

	"github.com/go-kickstart-orm/model/entity"
	"github.com/go-kickstart-orm/model/repository"
	"github.com/go-kickstart-orm/util"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (us UserService) Save(user *entity.User) {
	if user.Password != "" {
		var err error
		user.Password, err = util.EncriptyMd5(user.Password)
		if err != nil {
			util.CheckErr(err)
		}
	}
	us.userRepository.Save(user)
}

func (us UserService) GetAll() (users []entity.User) {
	return us.userRepository.GetAll()
}

func (us UserService) CreateDefaultUser() {
	users := us.userRepository.GetAll()
	if len(users) == 0 {
		user := entity.User{
			FirstName: "Ayrton",
			LastName:  "Senna",
			Login:     "admin",
			Password:  "admin",
		}
		us.Save(&user)
		fmt.Println("User default Create", user)
	}
}
