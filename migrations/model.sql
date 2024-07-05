CREATE DATABASE misis_grades;

-- create teachers table
CREATE TABLE teachers(
    teacher_id  UUID        PRIMARY KEY     NOT NULL,  
    name        VARCHAR(60)                 NOT NULL,
    surname     VARCHAR(60)                 NOT NULL,
    email       VARCHAR(60) UNIQUE          NOT NULL,
    password    VARCHAR(255)                NOT NULL,
    created_at  TIMESTAMP   DEFAULT         CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP   DEFAULT         CURRENT_TIMESTAMP
);

-- create courses table
CREATE TABLE courses(
    course_id   UUID        PRIMARY KEY     NOT NULL,  
    course_name VARCHAR(60)                 NOT NULL,
    created_at  TIMESTAMP   DEFAULT         CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP   DEFAULT         CURRENT_TIMESTAMP
);

-- Create groups table
CREATE TABLE groups(
    group_id   UUID        PRIMARY KEY     NOT NULL,  
    group_name VARCHAR(60)                 NOT NULL,
    created_at  TIMESTAMP   DEFAULT         CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP   DEFAULT         CURRENT_TIMESTAMP
)