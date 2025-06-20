package main

import (
	"log"

	"github.com/Go-Master-Code/rest-api/book"
	"github.com/Go-Master-Code/rest-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//koneksi database
	dsn := "root:root@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB connection error")
	}

	//auto migrate : membuat table berdasarkan struct yang didefinisikan di Golang, bisa di run lagi script di bawah apabila ada perubahan pada struktur tabel, misal mau tambah atribut discount
	db.AutoMigrate(&book.Book{})

	//=====REPOSITORY=====
	//FindAll()
	// bookRepository := book.NewRepository(db)

	// books, err := bookRepository.FindAll()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// for _, book := range books {
	// 	fmt.Println("Title: ", book.Title)
	// 	fmt.Println("Description: ", book.Description)
	// 	fmt.Println("Price: ", book.Price)
	// 	fmt.Println("Rating: ", book.Rating)
	// }

	//FindByID()
	// bookRepository := book.NewRepository(db)
	// book, err := bookRepository.FindById(3)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("ID: ", book.ID)
	// fmt.Println("Title: ", book.Title)
	// fmt.Println("Price: ", book.Price)
	// fmt.Println("Rating: ", book.Rating)

	//Nanti pakai dependency injection disini
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHander(bookService)

	//contoh penggunaan bookFileRepository (baca dan tulis data buku ke dalam file .txt)
	//contoh handler fileRepository yang mengimplementasikan interface Repository
	// bookFileRepository := book.NewFileRepository()
	// bookService := book.NewService(bookFileRepository) //ini bisa dilakukan karena pada NewService, yang dibutuhkan adalah repository
	// //bisa bookRepository, bisa bookFileRepository atau apa saja yang memenuhi kontrak interface Repository
	// bookHandler := handler.NewBookHander(bookService)

	//CRUD
	// //CREATE DATA
	// book := book.Book{}
	// book.Title = "Perjuangan Bangsa"
	// book.Price = 240000
	// book.Discount = 7
	// book.Rating = 7
	// book.Description = "Kisah nyata"

	// err = db.Create(&book).Error

	// if err != nil {
	// 	fmt.Println("========================")
	// 	fmt.Println("Error creating book data")
	// 	fmt.Println("========================")
	// }

	//READ DATA
	// var book book.Book

	// err = db.Debug().First(&book, 2).Error //untuk memunculkan sql, pakai func Debug() sebelum First()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("Title: ", book.Title)
	// fmt.Printf("Book object %v", book)

	// read data > 1 record
	// var books []book.Book

	// err = db.Debug().Find(&books).Error

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _, b := range books {
	// 	fmt.Println(b.Title)
	// 	fmt.Printf("Book object %v", b)
	// }

	//Find buku berdasarkan nama field, misalnya title
	// var books []book.Book
	// err = db.Debug().Where("rating = ?", "7").Find(&books).Error
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for _, b := range books {
	// 	fmt.Println(b.Title)
	// 	fmt.Printf("Book object %v", b)
	// }

	//UPDATE DATA
	// var book book.Book //data single

	// err = db.Debug().Where("id = ?", 1).First(&book).Error

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// book.Title = "Perjuangan Bangsa"
	// book.Description = "Kisah perjuangan bangsa"
	// book.Price = 150000
	// book.Rating = 9

	// err = db.Save(&book).Error //update data berdasarkan perubahan atribut di atas
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//DELETE DATA
	// var book book.Book //data single

	// err = db.Debug().Where("id = ?", 1).First(&book).Error

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// err = db.Delete(book).Error

	// if err != nil {
	// 	fmt.Println(err)
	// }

	//new router
	router := gin.Default()

	//API Versioning jika ada api baru, biasanya ditambahkan path baru misalnya v1
	v1 := router.Group("/v1")
	v2 := router.Group("/v2")

	/*router default
	//router GET
	router.GET("/", rootHandler)       //default path
	router.GET("/hello", helloHandler) //path hello
	//router.GET("/books/:id", booksHandler) //:id adalah sebuah PATH variable
	router.GET("/books/:id/:title", booksHandler) //untuk tangkap >1 PATH variable
	router.GET("/query", queryHandler)
	*/

	//router baru dengan v1
	v1.GET("/", bookHandler.RootHandler)
	v1.GET("/hello", bookHandler.HelloHandler)
	v1.GET("/books/:id/:title", bookHandler.BooksHandler)
	v1.GET("/query", bookHandler.QueryHandler)

	v1.POST("/books", bookHandler.CreateBookHandler)
	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.GET("/book/:id", bookHandler.GetBooksByIdHandler)
	v1.PUT("/book/:id", bookHandler.UpdateBookHandler)
	v1.DELETE("/book/:id", bookHandler.DeleteHandler)

	//api versioning dengan path v2/
	v2.GET("/buku/:kode/:judul", bookHandler.BukuHandler)
	v2.GET("/query", bookHandler.BukuQueryHandler)

	router.Run("localhost:8080") //custom port sertakan localhost di depannya agar tidak perlu permission

	//struktur
	//main
	//handler
	//service
	//repository
	//db
	//mysql
}
