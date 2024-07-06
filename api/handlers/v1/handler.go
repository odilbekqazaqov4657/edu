package v1

import (
	"app/models"
	"app/storage"
	"context"
	"encoding/json"
	"log"
	"strconv"
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

func (h *handler) GetTeacherById(ctx *gin.Context) {

	id := ctx.Param("id")

	teacher, err := h.storage.TeacherRepo().GetTeacherByID(context.Background(), id)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, teacher)

}

func (h *handler) GetTeachersList(ctx *gin.Context) {

	var req models.GetListReq

	Limit := ctx.Query("limit")
	Page := ctx.Query("page")

	req.Limit, _ = strconv.Atoi(Limit)
	req.Page, _ = strconv.Atoi(Page)

	teacher, err := h.storage.TeacherRepo().GetTeachersList(context.Background(), req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(200, teacher)

}

func (h *handler) UpdateTeacher(ctx *gin.Context) {

	var req *models.Teacher

	ctx.BindJSON(&req)

	err := h.storage.TeacherRepo().UpdateTeacher(context.Background(), req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, "updated succesfully !")
}

func (h *handler) DeleteTeacher(ctx *gin.Context) {

	id := ctx.Param("id")

	err := h.storage.TeacherRepo().DeleteTeacher(context.Background(), id)

	if err != nil {
		log.Println("error on deleting teacher", err)
		return
	}

	ctx.JSON(201, "deleted successfully")
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
