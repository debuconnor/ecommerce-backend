package product

type product struct{
	id			int
	sku			string
	name 		string
	price 		float64
	description	string
	color		string
	createdAt	string
	enabled		bool
	category	[]int
	image		string
}

func New() product{
	item := new(product)
	return *item
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

func (item *product) SetColor(color string){
	item.color = color
}

func (item *product) GetColor() string{
	return item.color
}

func (item *product) SetCreatedAt(createdAt string){
	item.createdAt = createdAt
}

func (item *product) GetCreatedAt() string{
	return item.createdAt
}

func (item *product) SetEnabled(enabled bool){
	item.enabled = enabled
}

func (item *product) GetEnabled() bool{
	return item.enabled
}

func (item *product) SetCategory(category []int){
	item.category = category
}

func (item *product) GetCategory() []int{
	return item.category
}

func (item *product) SetImage(image string){
	item.image = image
}

func (item *product) GetImage() string{
	return item.image
}