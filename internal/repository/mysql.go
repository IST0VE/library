package repository

import (
	"database/sql"

	"github.com/IST0VE/library/config"
)

type Repository struct {
	db *sql.DB
}

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func NewRepository(cfg *config.Config) (*Repository, error) {
	db, err := sql.Open("mysql", cfg.DBUser+":"+cfg.DBPassword+"@tcp("+cfg.DBHost+")/"+cfg.DBName)
	if err != nil {
		return nil, err
	}
	// Инициализация схемы БД, если нужно
	return &Repository{db: db}, nil
}

func (r *Repository) AddBook(title, author string, year int) error {
	_, err := r.db.Exec("INSERT INTO books (title, author, year_published) VALUES (?, ?, ?)", title, author, year)
	return err
}

// Получение всех книг
func (r *Repository) GetAllBooks() ([]Book, error) {
	var books []Book
	rows, err := r.db.Query("SELECT id, title, author, year_published FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.YearPublished); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

// Поиск книги по названию или автору
func (r *Repository) SearchBooks(query string) ([]Book, error) {
	var books []Book
	rows, err := r.db.Query("SELECT id, title, author, year_published FROM books WHERE title LIKE ? OR author LIKE ?", "%"+query+"%", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.YearPublished); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

// Здесь добавьте функции для добавления, поиска и получения книг
