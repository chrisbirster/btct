package server

import (
	"fmt"
	"net/http"
	"strconv"

	app "btct/app"
	"github.com/labstack/echo/v4"
)

func FuncTaskIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Task Queue API")
	}
}

func FuncTaskId(appInstance *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid task id")
		}
		task, err := appInstance.GetTask(taskId)
		if err != nil {
			return c.String(http.StatusNotFound, fmt.Sprintf("Err: %v", err))
		}
		return c.JSON(http.StatusOK, task)
	}
}

func FuncTaskAdd(appInstance *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		description := c.FormValue("description")
		if description == "" {
			return c.String(http.StatusBadRequest, "Task description is required")
		}
		task := appInstance.AddTask(description)
		return c.JSON(http.StatusCreated, task)
	}
}

func FuncTaskMarkComplete(appInstance *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid task ID")
		}
		err = appInstance.MarkTaskComplete(id)
		if err != nil {
			return c.String(http.StatusNotFound, err.Error())
		}
		return c.String(http.StatusOK, fmt.Sprintf("Task %d marked as complete", id))
	}
}

// FuncTaskFromNFC handles NFC requests to add tasks based on query params
func FuncTaskFromNFC(appInstance *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Extract query parameters
		taskDescription := c.QueryParam("description")
		if taskDescription == "" {
			return c.String(http.StatusBadRequest, "Task description is required")
		}

		// Add task to the app's task queue
		task := appInstance.AddTask(taskDescription)

		// Return the created task as a JSON response
		return c.JSON(http.StatusCreated, task)
	}
}
