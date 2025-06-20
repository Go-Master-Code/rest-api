package book

type Service interface {
	//daftar func yg berkaitan dgn tbl book
	FindAll() ([]Book, error) //return value: slice of Book, error
	FindById(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(ID int, book BookRequest) (Book, error)
	Delete(ID int) (Book, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository} //return struct service dengan atribut repository yang dikirim dari param
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err
	//return s.repository.FindAll() -> ini cara yang lebih singkat
}

func (s *service) FindById(ID int) (Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	//parsing price
	price, _ := bookRequest.Price.Int64()

	//rating, _ := bookRequest.Rating.Int64()
	//discount, _ := bookRequest.Discount.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      bookRequest.Rating,   //tidak diparsing kerena di request type nya int, bukan json.Number
		Discount:    bookRequest.Discount, //tidak diparsing kerena di request type nya int, bukan json.Number
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(ID int, bookRequest BookRequest) (Book, error) {
	//cari dulu data barang dengan id tertentu
	book, _ := s.repository.FindById(ID)

	//parsing price
	price, _ := bookRequest.Price.Int64()

	//rating, _ := bookRequest.Rating.Int64()
	//discount, _ := bookRequest.Discount.Int64()

	//set value dari masing-masing atribut entity buku bersasarkan isi book request (request body)
	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.Price = int(price)
	book.Discount = bookRequest.Discount
	book.Rating = bookRequest.Rating

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) { //layer service perlu parameter ID yang ditangkap oleh handler
	//cari dulu data barang dengan id tertentu
	book, _ := s.repository.FindById(ID)

	deletedBook, err := s.repository.Delete(book)
	return deletedBook, err
}
