package service

import "Golang-Book-Author-API/domain"

type AuthorService interface {
	GetAllAuthors() ([]domain.Author, error)
}

type DefaultAuthorService struct {
	repo domain.AuthorRepository
}

func (s DefaultAuthorService) GetAllAuthors() ([]domain.Author, error) {
	return s.repo.FindAll()
}

func NewAuthorService(repository domain.AuthorRepository) DefaultAuthorService {
	return DefaultAuthorService{repository}
}
