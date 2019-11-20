package product

// DAO is DAO of product information
type DAO struct{}

// NewDAO instantiates ProductDAO
func NewDAO() *DAO {
	return &DAO{}
}

// GetAllProducts returns list of products
func (dao *DAO) GetAllProducts() []Product {
	products := []Product{
		{SKU: "12345", Name: "iPhone", Price: 100.00, Unit: 100},
		{SKU: "22345", Name: "Pixel 4", Price: 70.00, Unit: 200},
	}
	return products
}
