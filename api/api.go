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

	// teachers
	router.POST("/create_teacher", h.CreateTeacher)
	router.GET("/teacher_by_id/:id", h.GetTeacherById)
	router.GET("/teachers_list", h.GetTeachersList)

	router.Run(":8080")
}
