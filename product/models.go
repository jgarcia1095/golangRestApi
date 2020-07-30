package product

type Product struct {
	Id           int     `json:"id"`
	ProductCode  string  `json:"productCode"`
	ProductName  string  `json:"productName"`
	Description  string  `json:"description"`
	StandardCost float64 `json:"standardCost"`
	ListPrice    float64 `json:"listPrice"`
	Category     string  `json:"category"`
}

type ProductList struct {
	Data         []*Product `json:"data"`
	TotalRecords int        `json:"totalRecords"`
}

type BestEmployee struct {
	ID          int    `json:"id"`
	TotalVentas int    `json:"totalVentas"`
	LastName    string `json:"lastName"`
	FirstName   string `json:"FirstName"`
}
