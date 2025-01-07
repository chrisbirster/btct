package server

import (
	"embed"
	"fmt"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"

	app "btct/app"
)

const PORT = "42069"

func StartServer(appInstance *app.App, staticFiles embed.FS) {
	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			"http://localhost:42069/auth/google/callback",
			"email", "profile",
		),
	)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	secret := os.Getenv("BTCT_SECRET")
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(secret))))

	// Serve static files from the embedded `dist` directory
	e.StaticFS("/", echo.MustSubFS(staticFiles, "dist"))

	// oauth routes
	// --- OAuth routes ---
	e.GET("/auth/:provider", FuncGoogleLogin())
	e.GET("/auth/:provider/callback", FuncGoogleLoginCallback())

	// api routes
	api := e.Group("/api", requireAuth)
	api.GET("/", FuncTaskIndex())
	api.GET("/me", FuncMe())
	api.GET("/task/:id", FuncTaskId(appInstance))
	api.POST("/tasks/create", FuncTaskAdd(appInstance))
	api.PUT("/tasks/:id/complete", FuncTaskMarkComplete(appInstance))
	api.POST("/nfc", FuncTaskFromNFC(appInstance))

	// start app
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", PORT)))

}
