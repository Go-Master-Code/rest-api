package book

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
	Discount    int    `json:"discount"`
	//CreatedAt   time.Time tidak perlu ada di response
	//UpdatedAt   time.Time
}
