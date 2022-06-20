package domain

import (
	"Golang-Book-Author-API/errs"
	"Golang-Book-Author-API/logger"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type AuthorRepositoryDb struct {
	client *sql.DB
}

func (d AuthorRepositoryDb) FindAll() ([]Author, error) {
	findAllSql := "select author_id, name, birthplace, movement from authors"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		logger.Error("Error while querying author table " + err.Error())
		return nil, err
	}

	authors := make([]Author, 0)
	for rows.Next() {
		var a Author
		err := rows.Scan(&a.Id, &a.Name, &a.Birthplace, &a.Movement)
		if err != nil {
			logger.Error("Error while scanning authors " + err.Error())
			return nil, err
		}
		authors = append(authors, a)
	}
	return authors, nil
}

func (d AuthorRepositoryDb) ById(id string) (*Author, *errs.AppError) {
	customerSql := "select author_id, name, birthplace, movement from authors where author_id = ?"

	row := d.client.QueryRow(customerSql, id)
	var a Author
	err := row.Scan(&a.Id, &a.Name, &a.Birthplace, &a.Movement)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("author not found")
		} else {
			logger.Error("Error while scanning author " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &a, nil
}

func NewAuthorRepositoryDb() AuthorRepositoryDb {
	client, err := sql.Open("mysql", "root:Bb080695@tcp(localhost:3306)/reading")
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return AuthorRepositoryDb{client}
}
