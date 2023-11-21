package main

import (
	"verbme-api/controller"
	"verbme-api/db"
	"verbme-api/repository"
	"verbme-api/router"
	"verbme-api/usecase"
)

func main() {
	//DB起動
	db := db.NewDB()
	//バリデーターの初期化
	// userValidator := validator.NewUserValidator()
	// taskValidator := validator.NewTaskValidator()
	userRepository := repository.NewUserRepository(db)
	// taskRepository := repository.NewTaskRepository(db)
	userUsecase := usecase.NewUserUseCase(userRepository)
	// taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	userController := controller.NewUserController(userUsecase)
	// taskController := controller.NewTaskController(taskUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
