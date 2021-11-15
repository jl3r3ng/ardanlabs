package book

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log"
	"time"
)
// User manages the set of API's for user access.
type Book struct {
	log *log.Logger
	db  *sqlx.DB
}

// New constructs a User for api access.
func New(log *log.Logger, db *sqlx.DB) Book {
	return Book{
		log: log,
		db:  db,
	}
}
func (u Book)GetBook(bookID int) (Books, error) {
	//Retrieve
	res := Book{}

	var id int
	var name string
	var author string
	var pages int
	var publicationDate pq.NullTime

	err := u.db.QueryRow(`SELECT id, name, author, pages, publication_date FROM books where id = $1`, bookID).Scan(&id, &name, &author, &pages, &publicationDate)
	if err == nil {
		res = Book{ID: id, Name: name, Author: author, Pages: pages, PublicationDate: publicationDate.Time}
	}

	return res, err
}

func (u Book) AllBooks() ([]Books, error) {
	//Retrieve
	books := []Books{}

	rows, err := u.db.Query(`SELECT id, name, author, pages, publication_date FROM books order by id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var author string
		var pages int
		var publicationDate pq.NullTime

		err = rows.Scan(&id, &name, &author, &pages, &publicationDate)
		if err != nil {
			return books, err
		}

		currentBook := Books{ID: id, Name: name, Author: author, Pages: pages}
		if publicationDate.Valid {
			currentBook.PublicationDate = publicationDate.Time
		}

		books = append(books, currentBook)
	}

	return books, err
}

func InsertBook(name, author string, pages int, publicationDate time.Time) (int, error) {
	//Create
	var bookID int
	err := db.QueryRow(`INSERT INTO books(name, author, pages, publication_date) VALUES($1, $2, $3, $4) RETURNING id`, name, author, pages, publicationDate).Scan(&bookID)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Last inserted ID: %v\n", bookID)
	return bookID, err
}

func UpdateBook(id int, name, author string, pages int, publicationDate time.Time) (int, error) {
	//Create
	res, err := db.Exec(`UPDATE books set name=$1, author=$2, pages=$3, publication_date=$4 where id=$5 RETURNING id`, name, author, pages, publicationDate, id)
	if err != nil {
		return 0, err
	}

	rowsUpdated, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsUpdated), err
}

func RemoveBook(bookID int) (int, error) {
	//Delete
	res, err := db.Exec(`delete from books where id = $1`, bookID)
	if err != nil {
		return 0, err
	}

	rowsDeleted, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsDeleted), nil
}
