package repoi

import (
	"app/models"
	"context"
)

type TeacherRepoI interface {
	CreateTeacher(ctx context.Context, req *models.Teacher) error
	GetTeacherByID(ctx context.Context, id string) (*models.Teacher, error)
	GetTeachersList(ctx context.Context, req models.GetListReq) (models.GetTeachersListResp, error)
	UpdateTeacher(ctx context.Context, req *models.Teacher) error
	DeleteTeacher(ctx context.Context, id string) error
}
