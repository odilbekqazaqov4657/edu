package api

import (
	v1 "app/api/handlers/v1"
	"app/storage"

	"github.com/gin-gonic/gin"
)

func Api(storage storage.StorageI) {
	router := gin.Default()

	h := v1.NewHandler(storage)

	router.GET("/ping", h.Ping)
	router.POST("/teacher", h.CreateTeacher)

	router.Run(":8080")
}
