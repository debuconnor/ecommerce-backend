package customer

type customer struct{
	id			int
	email		string
	firstname	string
	middlename	string
	lastname	string
	username	string
	phone		string
	address1	string
	address2	string
	address3	string
	createdAt	string
	updatedAt	string
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

func (user *customer) SetEmail(email string){
	user.email = email
}

func (user *customer) GetEmail() string{
	return user.email
}

func (user *customer) SetFirstname(firstname string){
	user.firstname = firstname
}

func (user *customer) GetFirstname() string{
	return user.firstname
}

func (user *customer) SetMiddlename(middlename string){
	user.middlename = middlename
}

func (user *customer) GetMiddlename() string{
	return user.middlename
}

func (user *customer) SetLastname(lastname string){
	user.lastname = lastname
}

func (user *customer) GetLastname() string{
	return user.lastname
}

func (user *customer) SetUsername(username string){
	user.username = username
}

func (user *customer) GetUsername() string{
	return user.username
}

func (user *customer) SetPhone(phone string){
	user.phone = phone
}

func (user *customer) GetPhone() string{
	return user.phone
}

func (user *customer) SetAddress1(address string){
	user.address1 = address
}

func (user *customer) GetAddress1() string{
	return user.address1
}

func (user *customer) SetAddress2(address string){
	user.address2 = address
}

func (user *customer) GetAddress2() string{
	return user.address2
}

func (user *customer) SetAddress3(address string){
	user.address3 = address
}

func (user *customer) GetAddress3() string{
	return user.address3
}

func (user *customer) setCreatedAt(createdAt string){
	user.createdAt = createdAt
}

func (user *customer) GetCreatedAt() string{
	return user.createdAt
}

func (user *customer) setUpdatedAt(updatedAt string){
	user.updatedAt = updatedAt
}

func (user *customer) GetUpdatedAt() string{
	return user.updatedAt
}