package employee

type Employee struct {
	ID            int    `json:"id"`
	LastName      string `json:"lastName"`
	FirstName     string `json:"FirstName"`
	Company       string `json:"company"`
	EmailAddress  string `json:"emailAddress"`
	JobTitle      string `json:"jobTitle"`
	BusinessPhone string `json:"businessPhone"`
	HomePhone     string `json:"homePhone"`
	MobilePhone   string `json:"mobilePhone"`
	FaxNumber     string `json:"faxNumber"`
	Address       string `json:"address"`
}

type EmployeeList struct {
	Data         []*Employee `json:"data"`
	TotalRecords int         `json:"totalRecords"`
}

// type Product struct {
// 	ID           int     `json:"id"`
// 	ProductCode  string  `json:"productCode"`
// 	ProductName  string  `json:"productName"`
// 	Description  string  `json:"description"`
// 	StandardCost float64 `json:"standardCost"`
// 	ListPrice    float64 `json:"listPrice"`
// 	Category     string  `json:"category"`
// }

// type ProductList struct {
// 	Data         []*Product `json:"data"`
// 	TotalRecords int        `json:"totalRecords"`
// }

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
