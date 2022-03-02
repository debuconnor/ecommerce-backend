package product

import "fmt"

type product struct{
	id			int
	name 		string
	price 		float32
	description	string
	color		string
	createdAt	string
	enabled		bool
	category	[]int
	image		string
}

func Test(){
	product := product{
		id: 1,
		name: "testProduct",
		price: 10,
		description: "This is a test product",
		color: "black",
		createdAt: "2022-03-02",
		enabled: true,
		category: []int{1,2},
		image: "none",
	}
	fmt.Println("product test")
	fmt.Println(product)
}
