package main

import (
	"gorm.io/gorm"
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
	setRoutes(e, db)
	e.Logger.Fatal(e.Start(":1323"))
}

func setRoutes(e *echo.Echo, db *gorm.DB) {
	tasks_controller := controller.TasksController{DB: db}
	e.GET("/tasks", tasks_controller.Index)
	e.GET("/tasks/completed", tasks_controller.Completed)
	e.POST("/tasks", tasks_controller.Create)
	e.PATCH("/tasks/:id", tasks_controller.Update)
	e.DELETE("/tasks/:id", tasks_controller.Delete)
}
