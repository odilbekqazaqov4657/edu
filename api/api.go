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
	router.GET("/teachers_list: id", h.GetTeachersList)
	router.PUT("/update_teachers: id", h.UpdateTeacher)
	router.DELETE("/delete_teacher: id", h.DeleteTeacher)

	router.Run(":8080")
}
