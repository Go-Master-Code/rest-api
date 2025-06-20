package book

import (
	"gorm.io/gorm"
)

//semua operasi database ditampng di repository.go

type Repository interface {
	//daftar func yg berkaitan dgn tbl book
	FindAll() ([]Book, error) //return value: slice of Book, error
	Create(book Book) (Book, error)
	FindById(ID int) (Book, error)
	Update(book Book) (Book, error)
	Delete(book Book) (Book, error)
}

// implementasi interface Repository
type repository struct {
	//properti
	db *gorm.DB
}

// func ini dipanggil setiap sebelum melakukan operasi database
func NewRepository(db *gorm.DB) *repository {
	return &repository{db} //return struct repository dengan atribut db yang dikirim dari param
}

func (r *repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error

	return books, err
}

// method FindById hampir sama dengan FindAll, hanya bedanya single instance
func (r *repository) FindById(ID int) (Book, error) {
	var book Book                     //single instance
	err := r.db.Find(&book, ID).Error //masukkan parameter ID pada klausa where setelah &book

	return book, err
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error //operasi database
	return book, err
}

func (r *repository) Delete(book Book) (Book, error) {
	err := r.db.Delete(&book).Error //operasi database
	return book, err
}
