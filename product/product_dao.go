package product

// ProductDAO is DAO of product information
type ProductDAO struct{}

// NewProductDAO instantiates ProductDAO
func NewProductDAO() *ProductDAO {
	return &ProductDAO{}
}

// GetAllProducts returns list of products
func (dao *ProductDAO) GetAllProducts() []Product {
	products := []Product{
		{SKU: "12345", Name: "iPhone", Price: 100.00, Unit: 100},
		{SKU: "22345", Name: "Pixel 4", Price: 70.00, Unit: 200},
	}
	return products
}
