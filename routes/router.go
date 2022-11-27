package routes

import (
	ctrl "admin/controller"
	"admin/controller/auth"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GetRoutes() *echo.Echo {
	e := echo.New()
	useMiddlewares(e)

	// CRUD User
	e.POST("v1/admin/user/create", ctrl.CreateRegisteredUser)
	e.GET("v1/admin/user", ctrl.GetRegisteredUser)
	e.POST("v1/admin/user", ctrl.UpdateRegisteredUser)
	e.DELETE("v1/admin/user", ctrl.DeleteRegisteredUser)

	// CRUD User Points
	e.GET("v1/admin/user/points", ctrl.GetUserPoints)
	e.POST("v1/admin/article/points", ctrl.UpdateArticlePointById)
	e.POST("v1/admin/user/points", ctrl.ResetUserPoints)

	// CRUD Articles
	e.GET("v1/admin/article", ctrl.GetArticleByID)
	e.POST("v1/admin/article", ctrl.UpdateArticleByID)
	e.POST("v1/admin/article/create", ctrl.InsertArticle)
	e.DELETE("v1/admin/article", ctrl.DeleteArticleByID)

	// Login and Register
	e.POST("v1/admin/login", auth.Login)
	e.POST("v1/admin/register", auth.Register)
	return e
}

func useMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
}
