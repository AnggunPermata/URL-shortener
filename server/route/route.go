package route

import (
	"github.com/anggunpermata/url-shortener/server/controller"
	"github.com/labstack/echo"
)

func InitRoute(e *echo.Echo){
	e.GET("/", controller.Homepage)
	//e.POST("/url_shortener", controller.UrlShortener)
	e.GET("/url_shortener", controller.RouteSubmitPost)
	e.POST("/url_shortener", controller.RouteSubmitPost)
	e.GET("/:url", controller.AccessURL)
}
