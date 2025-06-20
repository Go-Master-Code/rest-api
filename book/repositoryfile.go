package book

import (
	"fmt"
)

//repository ini dibuat untuk menyimpan data ke dalam file .txt
//file ini coba mengimplementasikan interface Repository
//artinya semua function yang ada pada interface / kontrak tsb harus dimiliki oleh
//struct fileRepository di bawah ini

type fileRepository struct {
	//struct kosong
}

// =====implementasi interface Repository=====
func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (r *fileRepository) FindAll() ([]Book, error) {
	var books []Book
	fmt.Println("Find All")

	return books, nil //error = nil
}

// method FindById hampir sama dengan FindAll, hanya bedanya single instance
func (r *fileRepository) FindById(ID int) (Book, error) {
	var book Book //single instance
	fmt.Println("FInd By ID")

	return book, nil
}

func (r *fileRepository) Create(book Book) (Book, error) {
	fmt.Println("Create A Book")
	return book, nil
}
