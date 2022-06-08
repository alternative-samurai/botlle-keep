package entities

import (
	"example/src/infrastructures/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getRepo() *gorm.DB {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgresp@localhost:5432/bottle"))
	if err != nil {
		panic(err)
	}
	return db
}

func Insert(args *model.Bottle) *gorm.DB {
	return getRepo().Create(args)

}
