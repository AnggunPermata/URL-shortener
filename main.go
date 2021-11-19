package main

import (
	"fmt"
	"github.com/anggunpermata/url-shortener/config"
	"github.com/anggunpermata/url-shortener/helper/storage"
	"github.com/anggunpermata/url-shortener/server/route"
	"github.com/labstack/echo"
)

func main(){
	config.InitPort()
	storage.InitializeStore()
	e := echo.New()

	//register routes
	route.InitRoute(e)
	Port := fmt.Sprintf(":%d", config.PORT)
	if err := e.Start(Port); err != nil {
		e.Logger.Fatal("Cannot start server")
	}
}
