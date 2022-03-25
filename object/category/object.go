package category

type category struct{
	id				int
	Code			string	`json:"Code"`
	Name			string	`json:"Name"`
	Order			int		`json:"Order"`
	ParentCategory	int		`json:"ParentId"`
	Enabled			bool	`json:"Enabled"`
	CreatedAt		string	`json:"CreatedAt"`
	UpdatedAt		string	`json:"UpdatedAt"`
}

func New() category{
	item := new(category)
	return *item
}

func (item *category) setId(id int) {
	item.id = id
}

func (item *category) GetId() int{
	return item.id
}