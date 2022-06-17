package domain

type Author struct {
	Id         string
	Name       string
	Birthplace string
	Movement   string
}

type AuthorRepository interface {
	FindAll() ([]Author, error)
}
