package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"todo/internal/controller"
	"todo/internal/orm"
)

func main() {
	db, err := orm.NewDB()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	tasks_controller := controller.TasksController{DB: db}

	e.GET("/tasks", tasks_controller.Index)

	e.Logger.Fatal(e.Start(":1323"))
}
