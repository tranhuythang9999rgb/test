package entities

type Order struct {
	ProductID int64 `form:"product_id"`
	Quantity  int   `form:"quantity"`
}
