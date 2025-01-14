package main

import (
	"database/sql"
	"html/template"
)

type Problem struct {
	Name      string
	Title     string
	Statement template.HTML
}

type Submission struct {
	ID        int
	Problem   string
	Lang      string
	Status    string
	Source    string
	MaxTime   int
	MaxMemory int
	UserName  sql.NullString
	User      User `gorm:"foreignkey:UserName"`
}

type Task struct {
	Submission int
}

type SubmissionTestcaseResult struct {
	Submission int
	Testcase   string
	Status     string
	Time       int
	Memory     int
}

type User struct {
	Name     string
	Passhash string
	Admin    bool
}

func (u User) getName() sql.NullString {
	return sql.NullString{u.Name, u.Name != ""}
}
