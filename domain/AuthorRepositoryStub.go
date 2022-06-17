package domain

type AuthorRepositoryStub struct {
	authors []Author
}

func (s AuthorRepositoryStub) FindAll() ([]Author, error) {
	return s.authors, nil
}

func NewAuthorRepositoryStub() AuthorRepositoryStub {
	authors := []Author{
		{"1", "Virginia Woolf", "London, England", "Modernist"},
		{"2", "Sylvia Plath", "Boston, Massachusetts", "Confessional"},
	}
	return AuthorRepositoryStub{authors}
}
