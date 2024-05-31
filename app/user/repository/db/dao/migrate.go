package dao

import "TodoList_demo/app/user/repository/db/model"

func migrate() {
	db.AutoMigrate(&model.User{})
}
