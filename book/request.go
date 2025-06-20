package book

import "encoding/json"

// buat struct untuk menangkap input user di method POST
// tambahkan tag json dan binding (validasi input) untuk setiap field, lengkapnya baca di Github
type BookRequest struct {
	Title       string      `json:"title" binding:"required"`           //harus diisi
	Price       json.Number `json:"price" binding:"required,number"`    //harus diisi dan berupa number
	Rating      int         `json:"rating" binding:"required,number"`   //harus diisi dan berupa number
	Discount    int         `json:"discount" binding:"required,number"` //harus diisi dan berupa number
	Description string      `json:"description" binding:"required"`     //var SubTitle dipakai untuk menangkap json yang namanya sub_title
	//SubTitle    string      `json:"sub_title" binding:"required"`
	//untuk menangkap field sub_title pada requestBody json, pakai json tag pada atribut SubTitle
}
