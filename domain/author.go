package domain

import "Golang-Book-Author-API/errs"

type Author struct {
	Id         string `db:"author_id"`
	Name       string
	Birthplace string
	Movement   string
}

type AuthorRepository interface {
	FindAll() ([]Author, error)
	ById(string) (*Author, *errs.AppError)
}
