package api

import (
	"net/http"

	"github.com/NuttapolCha/simple-covid-grouping/app"
	"github.com/gin-gonic/gin"
)

type API struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

var (
	apis    []API
	service *app.App
)

// Init initialize a simple API HTTP router
func Init(application *app.App) *gin.Engine {
	if service == nil {
		service = application
	}

	router := gin.New()

	// middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	for _, api := range apis {
		switch api.Method {
		case http.MethodGet:
			router.GET(api.Path, api.HandlerFunc)
		case http.MethodPost:
			router.POST(api.Path, api.HandlerFunc)
		case http.MethodPut:
			router.PUT(api.Path, api.HandlerFunc)
		case http.MethodDelete:
			router.DELETE(api.Path, api.HandlerFunc)
		case http.MethodPatch:
			router.PATCH(api.Path, api.HandlerFunc)
		case http.MethodOptions:
			router.OPTIONS(api.Path, api.HandlerFunc)
		case http.MethodHead:
			router.HEAD(api.Path, api.HandlerFunc)
		default:
			router.Any(api.Path, api.HandlerFunc)
		}
	}
	return router
}
