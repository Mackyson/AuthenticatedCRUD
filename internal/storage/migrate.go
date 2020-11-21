package storage

import (
	"AuthenticatedCRUD/model"
	"AuthenticatedCRUD/pkg/DButil"
)

func Migrate() {

	db := DButil.GetClient()

	db.LogMode(true)
	db.SingularTable(true)

	db.AutoMigrate(&model.User{}, &model.Task{})
}
