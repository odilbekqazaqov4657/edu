package v1

import (
	"app/models"
	"app/storage"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	storage storage.StorageI
}

func NewHandler(storage storage.StorageI) handler {
	return handler{storage: storage}
}

func (h *handler) Ping(ctx *gin.Context) {
	ctx.JSON(200, map[string]string{"message": "pong"})
}

func (h *handler) CreateTeacher(ctx *gin.Context) {

	var reqBody models.TeacherCreateReq

	ctx.BindJSON(&reqBody)

	var teacher = &models.Teacher{}

	ParseDate[models.TeacherCreateReq, *models.Teacher](reqBody, teacher)

	teacher.TeacherID = uuid.New()
	teacher.CreatedAt = time.Now()
	teacher.UpdatedAt = time.Now()

	h.storage.TeacherRepo().CreateTeacher(ctx, teacher)

}

func ParseDate[T1 any, T2 any](data1 T1, data2 T2) error {
	byteData, err := json.Marshal(data1)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteData, data2)
	if err != nil {
		return err
	}
	return nil
}
