package product

type product struct{
	id			int
	Sku			string	`json:"SKU"`
	Name 		string	`json:"Name"`
	Price 		float64	`json:"Price"`
	Description	string	`json:"Description"`
	CreatedAt	string	`json:"CreatedAt"`
	UpdatedAt	string	`json:"UpdatedAt"`
	Enabled		bool	`json:"Enabled"`
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