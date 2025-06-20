package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Go-Master-Code/rest-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type bookHandler struct {
	bookService book.Service
}

// =====HANDLER FINAL=====
func NewBookHander(bookService book.Service) *bookHandler {
	return &bookHandler{bookService}
}

// =====HANDLER FINAL=====
func (h *bookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Eko Kurniawan",
		"bio":  "Software Engineer",
	})
}

func (h *bookHandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"content":  "Hello world",
		"subtitle": "Belajar golang REST API",
	})
}

func (h *bookHandler) BooksHandler(c *gin.Context) {
	//contoh URL: localhost:8080/books/1
	//contoh URL > 1 var: localhost:8080/books/1/Andra
	//URL parameter / disebut juga path variable
	id := c.Param("id")       //tangkap var "id" (PATH variable)
	title := c.Param("title") //tangkap var "title" (PATH variable)

	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (h *bookHandler) BukuHandler(c *gin.Context) {
	//contoh URL: localhost:8080/v2/books/1
	//contoh URL > 1 var: localhost:8080/v2/books/1/Andra
	//URL parameter / disebut juga path variable
	kode := c.Param("kode")   //tangkap var "id" (PATH variable)
	judul := c.Param("judul") //tangkap var "title" (PATH variable)

	c.JSON(http.StatusOK, gin.H{
		"kode":  kode,
		"judul": judul,
	})
}

func (h *bookHandler) QueryHandler(c *gin.Context) {
	//contoh URL: localhost:8080/query?title=bumi manusia
	//contoh URL 2 query: localhost:8080/query?title=Judul&price=50000
	//disebut dengan query string
	title := c.Query("title")
	price := c.Query("price")

	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"price": price,
	})
}

func (h *bookHandler) BukuQueryHandler(c *gin.Context) {
	//contoh URL: localhost:8080/v2/query?judul=bumi manusia
	//contoh URL 2 query: localhost:8080/v2/query?judul=manusia&harga=50000
	//disebut dengan query string
	judul := c.Query("judul")
	harga := c.Query("harga")

	c.JSON(http.StatusOK, gin.H{
		"judul": judul,
		"harga": harga,
	})
}

// =====Handler CRUD=====
func (h *bookHandler) CreateBookHandler(c *gin.Context) {

	var bookRequest book.BookRequest //var baru dari struct BookRequest -> nanti akan didapat dari body json di postman

	err := c.ShouldBindJSON(&bookRequest) //atribut pada struct akan berisi data request body yang diinput postman

	if err != nil {
		log.Fatal(err) //-> server akan berhenti ketika error
		//jika error, ganti dengan response json saja

		errorMessages := []string{}
		//looping untuk setiap error validasi yang terjadi
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage) //menambahkan pesan error ke dalam
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages, //semua pesan error dimasukkan ke key "errors"
		})
		return
	}

	book, err := h.bookService.Create(bookRequest) //handler akses service akses repository akses db dan entity book, return value nya berupa newBook (data yang dimasukkan pada body json)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return //agar proses berhenti
	}

	//jika sukses
	c.JSON(http.StatusOK, gin.H{
		// "title":       book.Title,
		// "description": book.Description,
		// "price":       book.Price,
		// "rating":      book.Rating,
		// "discount":    book.Discount,
		"data": book,
	})
}

func (h *bookHandler) UpdateBookHandler(c *gin.Context) {

	/* CONTOH REQUEST BODY JSON MASUKKAN DI POSTMAN
	{
		"title" : "PHP",
		"price" : "150000",
		"rating" : 4,
		"discount" : 8,
		"description" : "PHP Encyclopedia"
	}
	*/

	var bookRequest book.BookRequest
	err := c.ShouldBindJSON(&bookRequest) //atribut pada struct akan berisi data request body yang diinput postman

	if err != nil {
		//log.Fatal(err) -> server akan berhenti ketika error
		//jika error, ganti dengan response json saja

		errorMessages := []string{}
		//looping untuk setiap error validasi yang terjadi
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage) //menambahkan pesan error ke dalam
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages, //semua pesan error dimasukkan ke key "errors"
		})
		return
	}

	id, _ := strconv.Atoi(c.Param("id")) //ambil parameter ID dari URL
	book, err := h.bookService.Update(id, bookRequest)

	bookResponse := convertToBookResponse(book) //convert ke format json

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
		return //agar proses berhenti
	}

	//jika sukses
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse,
	})
}

func (h *bookHandler) GetBooksHandler(c *gin.Context) {
	books, err := h.bookService.FindAll() //var books ini masih dalam bentuk entity book, belum dalam format .json
	if err != nil {                       //jika error tampilkan pesan json
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err, //semua pesan error dimasukkan ke key "errors"
		})
		return
	}

	var booksResponse []book.BookResponse //inisiasi slice baru dari struct BookResponse

	for _, b := range books {
		bookResponse := convertToBookResponse(b)            //private func
		booksResponse = append(booksResponse, bookResponse) //tiap 1 objek book dimasukkan ke dalam slice booksResponse yang sudah diinisiasi di atas
	}

	//jika berhasil, tampilkan data books
	c.JSON(http.StatusOK, gin.H{
		"data": booksResponse, //kirim slice bookResponse sebagai data
	})

}

func (h *bookHandler) GetBooksByIdHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	//tangkap var "id" (PATH variable), konversi ke int
	b, err := h.bookService.FindById((id)) //pakai var b biar ga bentrok dengan book.BookResponse{} di bawah

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err, //semua pesan error dimasukkan ke key "errors"
		})
		return
	}

	//panggil private func untuk mengconvert data struct ke dalam format response json
	bookResponse := convertToBookResponse(b)

	//jika berhasil, tampilkan data books
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse, //kirim slice bookResponse sebagai data
	})

}

func (h *bookHandler) DeleteHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	//tangkap var "id" (PATH variable), konversi ke int
	b, err := h.bookService.Delete(id) //pakai var b biar ga bentrok dengan book.BookResponse{} di bawah

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err, //semua pesan error dimasukkan ke key "errors"
		})
		return
	}

	//panggil private func untuk mengconvert data struct ke dalam format response json
	bookResponse := convertToBookResponse(b)

	//jika berhasil, tampilkan data books
	c.JSON(http.StatusOK, gin.H{
		"data": bookResponse, //kirim slice bookResponse sebagai data
	})

}

// private func untuk mengubah response entity book menjadi response json
func convertToBookResponse(b book.Book) book.BookResponse { //penamaan parameter harus berbeda dengan book.BookResponse
	return book.BookResponse{
		ID:          b.ID,
		Title:       b.Title,
		Price:       b.Price,
		Description: b.Description,
		Rating:      b.Rating,
		Discount:    b.Discount,
	}
}
