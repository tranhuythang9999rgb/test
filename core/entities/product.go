package entities

type Product struct {
	CreatorID       int64   `form:"creator_id"`
	Name            string  `form:"name"`
	Price           float64 `form:"price"`
	Quantity        int     `form:"quantity"`
	Description     string  `form:"description"`
	DiscountPercent float64 `form:"discount_percent"`
	StatusSell      bool    `form:"status_sell"`
}
