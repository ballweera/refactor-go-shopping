package model

// Product is model of product infomration
type Product struct {
	SKU   string  `json:"sku"`          // sku of product
	Name  string  `json:"product_name"` // name of product
	Price float64 `json:"price"`        // price of product
	Unit  int64   `json:"unit"`         // unit of product
}
