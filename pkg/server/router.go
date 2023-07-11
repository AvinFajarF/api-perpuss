package server

import (
	"github.com/AvinFajarF/pkg/server/http"
	"github.com/gin-gonic/gin"
)

func NewRouterUser (UserHandler *http.UserHandler, BookHandler *http.BookHandler) *gin.Engine{

	router := gin.Default()

	v1 := router.Group("api/v1")

	v1.POST("/register", UserHandler.CreateUser)
	v1.POST("/login", UserHandler.Login)
	


	return router

}