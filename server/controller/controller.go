package controller

import (
	"github.com/anggunpermata/url-shortener/helper/models"
	"github.com/anggunpermata/url-shortener/helper/storage"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

func Homepage(c echo.Context)error{
	type M map[string]interface{}
	data := M{"message": "welcome to homepage !!"}
	return c.Render(http.StatusOK, "home.page.html", data)
}

func RouteSubmitPost(c echo.Context) error{
	if c.Request().Method == "POST" {
		originalUrl := c.FormValue("original_url")
		shortenedUrl := c.FormValue("shortened_url")
		if _, err := storage.RetrieveInitialUrl(shortenedUrl); err != nil {
			if err2 := storage.SaveUrlMapping(shortenedUrl, originalUrl, "guest"); err2 != nil{
				log.Println(err2)
				return c.JSON(http.StatusBadRequest, err2)
			}

			var data = map[string]interface{}{
				"message": "Short URL Created",
				"created": true,
				"original_url": originalUrl,
				"short_url": "https://anggunpermata-us.herokuapp.com/" + shortenedUrl,
			}
			return c.Render(http.StatusOK, "urlshortener.page.html", data)
		}

		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "url already exists",
		})
	}
	return c.Render(http.StatusOK, "urlshortener.page.html", nil)
}

func UrlShortener(c echo.Context) error {
	req := new(models.UrlShortener_Payload)
	if err := c.Bind(&req); err != nil{
		log.Println(err)
		return c.JSON(http.StatusBadRequest, "error taking payload")
	}
	originalUrl := req.OriginalURL
	shortenedUrl := req.ShortenedURL

	if _, err := storage.RetrieveInitialUrl(shortenedUrl); err != nil {
		if err2 := storage.SaveUrlMapping(shortenedUrl, originalUrl, "guest"); err2 != nil{
			log.Println(err2)
			return c.JSON(http.StatusBadRequest, err2)
		}
		return c.JSON(200, map[string]interface{}{
			"message": "successfully creating short url",
			"short url": "http://localhost:8080/" + shortenedUrl,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "url already exists",
	})
}

func AccessURL(c echo.Context) error {
	shortURL := c.Param("url")
	initialURL, err := storage.RetrieveInitialUrl(shortURL)
	if err !=nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "url not exists",
		})
	}
	c.Redirect(302, initialURL)
	return c.JSON(http.StatusOK, "success")
}