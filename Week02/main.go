package main

import (
	"fmt"
	"log"

	"Week02/internal/biz"
	"Week02/internal/dao"
	"Week02/internal/service"
)

type App struct {
	service.Service
}

func InitApp() *App {
	db := dao.NewDB()
	dao := dao.New(db)
	biz := biz.New(dao)
	svc := service.New(biz)
	return &App{svc}
}

func main() {
	app := InitApp()
	name, err := app.GetUserName()
	if err != nil {
		log.Fatalf("failed to get user name: %+v", err)
	}
	fmt.Println(name)
}
