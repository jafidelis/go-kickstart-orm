package db

import (
	"github.com/go-kickstart-orm/util"
	"github.com/jinzhu/gorm"

	//Driver postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
)

//GetConnection create a new connection with database
func GetConnection() *gorm.DB {
	config, err := util.GetConfiguration()
	util.CheckErr(err)

	dataSource := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.DB.Server, config.DB.User, config.DB.Database, config.DB.Password)

	db, err := gorm.Open(config.DB.Engine, dataSource)
	util.CheckErr(err)

	return db
}
