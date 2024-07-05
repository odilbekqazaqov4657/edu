package models

import (
	"time"

	"github.com/google/uuid"
)

type TeacherCreateReq struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Teacher represents a teacher in the system
type Teacher struct {
	TeacherID uuid.UUID `json:"teacher_id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Course represents a course in the system
type Course struct {
	CourseID   string    `json:"course_id"`
	CourseName string    `json:"course_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Group represents a group in the system
type Group struct {
	GroupID   string    `json:"group_id"`
	GroupName string    `json:"group_name"`
	CourseID  string    `json:"course_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Subject represents a subject in the system
type Subject struct {
	SubjectID   string    `json:"subject_id"`
	SubjectName string    `json:"subject_name"`
	GroupID     string    `json:"group_id"`
	TeacherID   string    `json:"teacher_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// Student represents a student in the system
type Student struct {
	StudentID string    `json:"student_id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	GroupID   string    `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Grade represents a grade in the system
type Grade struct {
	GradeID    string    `json:"grade_id"`
	GradeName  string    `json:"grade_name"`
	GradeValue int       `json:"grade_value"`
	GradeDate  time.Time `json:"grade_date"`
	SubjectID  string    `json:"subject_id"`
	GroupID    string    `json:"group_id"`
	StudentID  string    `json:"student_id"`
	TeacherID  string    `json:"teacher_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type GetListReq struct {
	Page  int
	Limit int
}

type GetTeachersListResp struct {
	Teachers []*Teacher
	Count    int
}
