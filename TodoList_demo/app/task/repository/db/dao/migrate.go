package dao

import "TodoList_demo/app/task/repository/db/model"

func migrate() {
	db.AutoMigrate(&model.Task{})
}
