package customer

type customer struct{
	id			int
	Email		string	`json:"Email"`
	Firstname	string	`json:"Firstname"`
	Middlename	string	`json:"Middlename"`
	Lastname	string	`json:"Lastname"`
	Username	string	`json:"Username"`
	Phone		string	`json:"Phone"`
	Address1	string	`json:"Address1"`
	Address2	string	`json:"Address2"`
	Address3	string	`json:"Address3"`
	CreatedAt	string	`json:"CreatedAt"`
	UpdatedAt	string	`json:"UpdatedAt"`
}

func New() customer{
	item := new(customer)
	return *item
}

func (user *customer) setId(id int) {
	user.id = id
}

func (user *customer) GetId() int{
	return user.id
}