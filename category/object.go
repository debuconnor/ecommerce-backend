package category

type category struct{
	id			int
	Code		string	`json:"code"`
	Name		string	`json:"name"`
	Order		int		`json:"order"`
	Enabled		bool	`json:"enabled"`
	CreatedAt	string	`json:"createdAt"`
	UpdatedAt	string	`json:"updatedAt"`
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