package server

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"

	app "btct/app"
)

var (
	MODE = os.Getenv("MODE")
	PORT = os.Getenv("PORT")
)

func StartServer(appInstance *app.App, staticFiles embed.FS) {

	goth.UseProviders(
		google.New(
			os.Getenv("GOOGLE_CLIENT_ID"),
			os.Getenv("GOOGLE_CLIENT_SECRET"),
			os.Getenv("GOOGLE_CALLBACK"),
			"email", "profile",
		),
	)

	e := echo.New()

	// --- Middleware ---
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	secret := os.Getenv("BTCT_SECRET")
	if secret == "" {
		log.Fatal("BTCT_SECRET environment variable is not set")
	}
	store := sessions.NewCookieStore([]byte(secret))
	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   86400 * 7, // 7 days
		Secure:   secure(),
		SameSite: http.SameSiteLaxMode,
	}
	e.Use(session.Middleware(store))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:  "dist",
		HTML5: true,
	}))

	// Serve static files from the embedded `dist` directory
	e.StaticFS("/", echo.MustSubFS(staticFiles, "dist"))

	// --- OAuth routes ---
	e.GET("/auth/:provider", FuncGoogleLogin())
	e.GET("/auth/:provider/callback", FuncGoogleLoginCallback())

	// --- API routes ---
	api := e.Group("/api", requireAuth)
	api.GET("/", FuncTaskIndex())
	api.GET("/me", FuncMe())
	api.GET("/tasks/:id", FuncTaskId(appInstance))
	api.GET("/tasks", FuncTaskList(appInstance))
	api.POST("/tasks/create", FuncTaskAdd(appInstance))
	api.PUT("/tasks/:id/complete", FuncTaskMarkComplete(appInstance))
	api.POST("/nfc", FuncTaskFromNFC(appInstance))

	// start app
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", PORT)))

}

func secure() bool {
	if MODE == "dev" {
		return false
	}
	return true
}
