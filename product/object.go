package product

type product struct{
	id			int
	sku			string
	name 		string
	price 		float64
	description	string
	createdAt	string
	updatedAt	string
	enabled		bool
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

func (item *product) SetSku(sku string){
	item.sku = sku
}

func (item *product) GetSku() string{
	return item.sku
}

func (item *product) SetName(name string){
	item.name = name
}

func (item *product) GetName() string{
	return item.name
}

func (item *product) SetPrice(price float64){
	item.price = price
}

func (item *product) GetPrice() float64{
	return item.price
}

func (item *product) SetDescription(desciption string){
	item.description = desciption
}

func (item *product) GetDescription() string{
	return item.description
}

func (item *product) setCreatedAt(createdAt string) {
	item.createdAt = createdAt
}

func (item *product) GetCreatedAt() string{
	return item.createdAt
}

func (item *product) setUpdatedAt(updatedAt string){
	item.updatedAt = updatedAt
}

func (item *product) GetUpdatedAt() string{
	return item.updatedAt
}

func (item *product) SetEnabled(enabled bool){
	item.enabled = enabled
}

func (item *product) GetEnabled() bool{
	return item.enabled
}