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

func (t *teacherRepo) GetTeacherByID(ctx context.Context, id string) (*models.Teacher, error) {

	query := `SELECT 
				teacher_id,
					name,
					surname,
					email,
					password,
					created_at,
					updated_at
				FROM 
					users
				WHERE
					teacher_id = $1`

	row := t.db.QueryRow(ctx, query, id)

	var teacher models.Teacher

	err := row.Scan(
		&teacher.TeacherID,
		&teacher.Name,
		&teacher.Surname,
		&teacher.Email,
		&teacher.Password,
		&teacher.CreatedAt,
		&teacher.UpdatedAt,
	)

	if err != nil {
		log.Println("Error on scanning teacher !", err)
		return nil, err
	}

	return &models.Teacher{}, nil
}

func (t *teacherRepo) GetTeachersList(ctx context.Context, req models.GetListReq) (models.GetTeachersListResp, error) {

	var teacher models.Teacher
	resp := models.GetTeachersListResp{}

	Limit := req.Limit
	offset := (req.Page - 1) * req.Limit

	query := `
        SELECT 
            teacher_id,
            name,
            surname,
            email,
            created_at,
            updated_at
        FROM 
            teachers
        LIMIT $1 OFFSET $2
    `

	row, err := t.db.Query(ctx, query, Limit, offset)

	if err != nil {
		log.Println("Error on scanning teachers schedule !", err)
		return resp, err
	}

	for row.Next() {

		row.Scan(
			&teacher.TeacherID,
			&teacher.Name,
			&teacher.Surname,
			&teacher.Email,
			&teacher.CreatedAt,
			&teacher.UpdatedAt,
		)

		resp.Teachers = append(resp.Teachers, &models.Teacher{})
	}

	return resp, nil

}

func (t *teacherRepo) UpdateTeacher(ctx context.Context, req *models.Teacher) error {

	query := `
		UPDATE
			teachers
		SET
			name = $1,
			surname = $2,
			email = $3,
			password = $4,
			created_at = $5,
			updated_at = $6
		WHERE
			teacher_id = $7
		`

	_, err := t.db.Exec(
		ctx, query,
		req.Name,
		req.Surname,
		req.Email,
		req.Password,
		req.CreatedAt,
		req.UpdatedAt,
		req.TeacherID,
	)

	if err != nil {
		log.Println("error on updating information of teacher:", err)
		return err
	}

	fmt.Println("Informations updated successfully")

	return nil
}

func (t *teacherRepo) DeleteTeacher(ctx context.Context, id string) error {

	query := `
		DELETE 
		FROM
			teachers
		WHERE
			teacher_id=$1;
	`

	_, err := t.db.Exec(ctx, query, id)

	if err != nil {

		log.Println("error on deleting teacher ", err)
		return err
	}

	fmt.Println("deleted successfully")
	return nil
}
