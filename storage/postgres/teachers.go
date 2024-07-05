package postgres

import (
	"app/models"
	repoi "app/storage/repoI"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type teacherRepo struct {
	db *pgx.Conn
}

func NewTeacherRepo(db *pgx.Conn) repoi.TeacherRepoI {
	return &teacherRepo{db: db}
}

func (t *teacherRepo) CreateTeacher(ctx context.Context, teacher *models.Teacher) error {

	query := `INSERT INTO
				teachers(
					teacher_id,
					name,
					surname,
					email,
					password,
					created_at,
					updated_at)
				VALUES(
					$1, $2, $3, $4, $5, $6, $7
				)`

	_, err := t.db.Exec(ctx, query, teacher.TeacherID, teacher.Name, teacher.Surname, teacher.Email, teacher.Password, teacher.CreatedAt, teacher.UpdatedAt)

	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println("New teacher inserted successfully")

	return nil

}

func (t *teacherRepo) GetTeachersList(ctx context.Context, req models.GetListReq) (models.GetTeachersListResp, error) {
	res := models.GetTeachersListResp{}

	return res, nil
}
