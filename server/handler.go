package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	app "btct/app"
	database "btct/database"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

const SESSION = "btct_session"

func FuncTaskIndex() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Task Queue API")
	}
}

func FuncTaskId(appInstance *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId, err := strconv.ParseInt(c.Param("id"), 10, 64)

		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid task id")
		}
		task, err := appInstance.Queries.GetTask(c.Request().Context(), taskId)
		if err != nil {
			return c.String(http.StatusNotFound, fmt.Sprintf("Err: %v", err))
		}
		return c.JSON(http.StatusOK, task)
	}
}
func FuncTaskList(appInstance *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		tasks, err := appInstance.Queries.ListTasks(c.Request().Context())
		if err != nil {
			return c.String(http.StatusNotFound, fmt.Sprintf("Err: %v", err))
		}
		return c.JSON(http.StatusOK, tasks)
	}
}

func FuncTaskAdd(appInstance *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		description := c.FormValue("description")
		if description == "" {
			return c.String(http.StatusBadRequest, "Task description is required")
		}
		task, err := appInstance.Queries.CreateTask(c.Request().Context(), database.CreateTaskParams{
			Description: description,
			Complete:    false,
		})
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to add task: %v", err))
		}
		return c.JSON(http.StatusCreated, task)
	}
}

func FuncTaskMarkComplete(appInstance *app.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid task ID")
		}
		err = appInstance.Queries.UpdateTaskComplete(c.Request().Context(), database.UpdateTaskCompleteParams{
			Complete: true,
			ID:       id,
		})
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
		description := c.QueryParam("description")
		if description == "" {
			return c.String(http.StatusBadRequest, "Task description is required")
		}

		// Add task to the app's task queue
		task, err := appInstance.Queries.CreateTask(c.Request().Context(), database.CreateTaskParams{
			Description: description,
			Complete:    false,
		})

		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("Failed to add task: %v", err))
		}
		return c.JSON(http.StatusCreated, task)

	}
}

// FuncGoogleLogin handles google auth
func FuncGoogleLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		gothic.GetProviderName = func(req *http.Request) (string, error) {
			provider := c.Param("provider")
			if provider != "" {
				return provider, nil
			}
			return "", c.String(http.StatusBadRequest, "unable to get provider")
		}

		gothic.BeginAuthHandler(c.Response().Writer, c.Request())
		return nil
	}
}

// FuncGoogleLogin handles google auth
func FuncGoogleLoginCallback() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := gothic.CompleteUserAuth(c.Response().Writer, c.Request())
		if err != nil {
			return c.String(http.StatusUnauthorized, err.Error())
		}
		log.Printf("User: %+v\n", user)

		sess, err := session.Get(SESSION, c)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		sess.Values["user_id"] = user.UserID
		sess.Values["avatar_url"] = user.AvatarURL
		sess.Values["email"] = user.Email

		log.Printf("Session values before save: %+v", sess.Values)
		err = sess.Save(c.Request(), c.Response().Writer)
		if err != nil {
			log.Printf("Error saving session: %v", err)
			return c.String(http.StatusBadRequest, err.Error())
		}

		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

// requireAuth protects API routes from unauthorized use
func requireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get(SESSION, c)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		userID, ok := sess.Values["user_id"].(string)
		if !ok || userID == "" {
			// not logged in
			c.Redirect(http.StatusSeeOther, "/auth/google")
		}

		// user logged in
		return next(c)
	}
}

func FuncMe() echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get(SESSION, c)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		userID, ok := sess.Values["user_id"].(string)
		if !ok || userID == "" {
			// Not logged in
			return c.NoContent(http.StatusUnauthorized)
		}

		email, _ := sess.Values["email"].(string)
		avatarURL, _ := sess.Values["avatar_url"].(string)

		// Return user info (or just "logged_in": true)
		return c.JSON(http.StatusOK, map[string]interface{}{
			"user_id":    userID,
			"email":      email,
			"avatar_url": avatarURL,
		})
	}
}
