package product

//Attribute IDs
const (
	ENABLED_ID		=	"1"
	NAME_ID			=	"2"
	PRICE_ID		=	"3"
	DESCRIPTION_ID	=	"4"
	COLOR_ID		=	"5"
	CREATEDAT_ID	=	"6"
)

//Store Procedures in the database
const (
	CreateProduct	=	"CreateProduct"
	GetProduct 		=	"GetProduct"
	UpdateSku		=	"UpdateProductSku"
	SetEnabled		=	"SetProductEnabled"
	SetName			=	"SetProductName"
	SetPrice		=	"SetProductPrice"
	SetDescription	=	"SetProductDescription"
	SetColor		=	"SetProductColor"
	SetCreatedAt	=	"SetProductCreatedAt"
	SetUpdatedAt	=	"SetProductUpdtedAt"
)