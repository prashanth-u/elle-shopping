package models

type Product struct {
    Id   int
    Name string
	Price float64
	Image string
	Discount float32
	Quantity int
	Description string
	CategoryId int
}