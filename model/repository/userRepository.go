package repository

import (
	"github.com/go-kickstart-orm/db"
	"github.com/go-kickstart-orm/model/entity"
)

type UserRepository struct{}

func (rep UserRepository) Save(user *entity.User) {
	con := db.GetConnection()
	defer con.Close()

	if con.NewRecord(user) {
		con.Create(&user)
	} else {
		con.Model(&user).Updates(user)
	}

}

func (rep UserRepository) GetAll() (users []entity.User) {
	con := db.GetConnection()
	defer con.Close()

	con.Find(&users)
	return users
}

func (rep UserRepository) FindById(id uint) (user entity.User) {
	con := db.GetConnection()
	defer con.Close()
	con.First(&user, id)
	return user
}

func (rep UserRepository) FindByLogin(login string) (user entity.User) {
	con := db.GetConnection()
	defer con.Close()
	con.Where("login=?", login).First(&user)
	return user
}
