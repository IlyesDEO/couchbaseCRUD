package api

import (
	contactHandler "github.com/IlyesDEO/goCrud/pkg/handlers/contact"
	"github.com/gin-gonic/gin"
)

func Start() {

	// Run gin server
	router := gin.Default()

	router.GET("/contact", contactHandler.GetAll)
	router.GET("/contact/:id", contactHandler.GetById)
	router.POST("/contact", contactHandler.Create)
	router.PUT("/contact/:id", contactHandler.Update)
	router.DELETE("/contact/:id", contactHandler.Delete)

	router.Run(":8080")
}
