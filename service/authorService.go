package service

import (
	"Golang-Book-Author-API/domain"
	"Golang-Book-Author-API/dto"
	"Golang-Book-Author-API/errs"
)

type AuthorService interface {
	GetAllAuthors() ([]domain.Author, error)
	GetAuthor(string) (*dto.AuthorResponse, *errs.AppError)
}

type DefaultAuthorService struct {
	repo domain.AuthorRepository
}

func (s DefaultAuthorService) GetAllAuthors() ([]domain.Author, error) {
	return s.repo.FindAll()
}

func (s DefaultAuthorService) GetAuthor(id string) (*dto.AuthorResponse, *errs.AppError) {
	a, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := a.ToDto()
	return &response, nil
}

func NewAuthorService(repository domain.AuthorRepository) DefaultAuthorService {
	return DefaultAuthorService{repository}
}
