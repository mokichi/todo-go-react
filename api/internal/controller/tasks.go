package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"todo/internal/entity"
)

type TasksController struct {
	DB *gorm.DB
}

func (ctrl *TasksController) Index(c echo.Context) error {
	var tasks []entity.Task
	ctrl.DB.Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}
