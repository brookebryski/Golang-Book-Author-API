package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type AuthorRepositoryDb struct {
	client *sql.DB
}

func (d AuthorRepositoryDb) FindAll() ([]Author, error) {
	findAllSql := "select author_id, name, birthplace, movement from authors"

	rows, err := d.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying author table " + err.Error())
		return nil, err
	}

	authors := make([]Author, 0)
	for rows.Next() {
		var a Author
		err := rows.Scan(&a.Id, &a.Name, &a.Birthplace, &a.Movement)
		if err != nil {
			log.Println("Error while scanning authors " + err.Error())
			return nil, err
		}
		authors = append(authors, a)
	}
	return authors, nil
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
