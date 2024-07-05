package storage

import (
	"app/storage/postgres"
	repoi "app/storage/repoI"

	"github.com/jackc/pgx/v5"
)

type StorageI interface {
	TeacherRepo() repoi.TeacherRepoI
}

type storage struct {
	teacherRepo repoi.TeacherRepoI
}

func NewStorage(db *pgx.Conn) StorageI {
	return &storage{
		teacherRepo: postgres.NewTeacherRepo(db),
	}
}

func (s *storage) TeacherRepo() repoi.TeacherRepoI {
	return s.teacherRepo
}
