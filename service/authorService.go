package service

import (
	"Golang-Book-Author-API/domain"
	"Golang-Book-Author-API/errs"
)

type AuthorService interface {
	GetAllAuthors() ([]domain.Author, error)
	GetAuthor(string) (*domain.Author, *errs.AppError)
}

type DefaultAuthorService struct {
	repo domain.AuthorRepository
}

func (s DefaultAuthorService) GetAllAuthors() ([]domain.Author, error) {
	return s.repo.FindAll()
}

func (s DefaultAuthorService) GetAuthor(id string) (*domain.Author, *errs.AppError) {
	return s.repo.ById(id)
}

func NewAuthorService(repository domain.AuthorRepository) DefaultAuthorService {
	return DefaultAuthorService{repository}
}
