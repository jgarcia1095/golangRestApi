package customers

type Customer struct {
	ID            int    `json:"id"`
	FirstName     string `json:"FirstName"`
	LastName      string `json:"lastName"`
	Address       string `json:"address"`
	BusinessPhone string `json:"businessPhone"`
	City          string `json:"city"`
	Company       string `json:"company"`
}

type CustomerList struct {
	Data         []*Customer `json:"data"`
	TotalRecords int         `json:"totalRecords"`
}

// type BestEmployee struct {
// 	ID          int    `json:"id"`
// 	TotalVentas int    `json:"totalVentas"`
// 	LastName    string `json:"lastName"`
// 	FirstName   string `json:"FirstName"`
// }

// type ProductTop struct {
// 	ID          int     `json:"id"`
// 	ProductName string  `json:"product_name"`
// 	Vendidos    float64 `json:"vendidos"`
// }
// type ProductTopResponse struct {
// 	Data        []*ProductTop `json:"data"`
// 	TotalVentas float64       `json:"totalVentas"`
// }
