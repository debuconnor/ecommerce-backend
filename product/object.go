package product

type product struct{
	id			int
	Sku			string	`json:"sku"`
	Name 		string	`json:"name"`
	Price 		float64	`json:"price"`
	Description	string	`json:"description"`
	CreatedAt	string	`json:"createdAt"`
	UpdatedAt	string	`json:"updatedAt"`
	Enabled		bool	`json:"enabled"`
}

func New() product{
	item := new(product)
	return *item
}

func (item *product) setId(id int) {
	item.id = id
}

func (item *product) GetId() int{
	return item.id
}