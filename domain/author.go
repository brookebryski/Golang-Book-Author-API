package domain

import (
	"Golang-Book-Author-API/dto"
	"Golang-Book-Author-API/errs"
)

type Author struct {
	Id         string `db:"author_id"`
	Name       string
	Birthplace string
	Movement   string
}

func (a Author) ToDto() dto.AuthorResponse {
	return dto.AuthorResponse{
		Id:         a.Id,
		Name:       a.Name,
		Birthplace: a.Birthplace,
		Movement:   a.Movement,
	}
}

type AuthorRepository interface {
	FindAll() ([]Author, error)
	ById(string) (*Author, *errs.AppError)
}
