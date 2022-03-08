package customer

type customer struct{
	id			int
	Email		string	`json:"email"`
	Firstname	string	`json:"firstname"`
	Middlename	string	`json:"middlename"`
	Lastname	string	`json:"lastname"`
	Username	string	`json:"username"`
	Phone		string	`json:"phone"`
	Address1	string	`json:"address1"`
	Address2	string	`json:"address2"`
	Address3	string	`json:"address3"`
	CreatedAt	string	`json:"createdAt"`
	UpdatedAt	string	`json:"updatedAt"`
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