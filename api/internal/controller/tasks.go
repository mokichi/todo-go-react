package controller

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"todo/internal/entity"
	"todo/internal/usecase"
)

type TasksController struct {
	DB *gorm.DB
}

func (controller *TasksController) Index(c echo.Context) error {
	u := usecase.TasksUsecase{DB: controller.DB}
	tasks := u.FindByCompleted(false)
	return c.JSON(http.StatusOK, tasks)
}

func (controller *TasksController) Completed(c echo.Context) error {
	u := usecase.TasksUsecase{DB: controller.DB}
	tasks := u.FindByCompleted(true)
	return c.JSON(http.StatusOK, tasks)
}

type taskParams struct {
	Title     *string
	Completed *bool
}

func setParamsToTask(task *entity.Task, params taskParams) {
	if params.Title != nil {
		task.Title = *params.Title
	}
	if params.Completed != nil {
		task.Completed = *params.Completed
	}
}

func (controller *TasksController) Create(c echo.Context) error {
	params := taskParams{}
	c.Bind(&params)
	task := entity.Task{}
	setParamsToTask(&task, params)

	u := usecase.TasksUsecase{DB: controller.DB}
	if err := u.Create(&task); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, task)
}

func (controller *TasksController) Update(c echo.Context) error {
	id := c.Param("id")
	u := usecase.TasksUsecase{DB: controller.DB}
	task, err := u.First(id)
	if err != nil {
		return err
	}

	params := taskParams{}
	c.Bind(&params)
	setParamsToTask(&task, params)

	if err = u.Update(&task); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, task)
}

func (controller *TasksController) Delete(c echo.Context) error {
	id := c.Param("id")
	u := usecase.TasksUsecase{DB: controller.DB}
	task, err := u.First(id)
	if err != nil {
		return err
	}

	if err = u.Delete(&task); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
