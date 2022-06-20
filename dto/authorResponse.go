package dto

type AuthorResponse struct {
	Id         string `db:"author_id"`
	Name       string `db:"full_name"`
	Birthplace string `db:"birthplace"`
	Movement   string `db:"movement"`
}
