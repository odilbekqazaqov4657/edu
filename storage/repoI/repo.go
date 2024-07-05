package repoi

import (
	"app/models"
	"context"
)

type TeacherRepoI interface {
	CreateTeacher(ctx context.Context, req *models.Teacher) error
	GetTeachersList(ctx context.Context, req models.GetListReq) (models.GetTeachersListResp, error)
}
