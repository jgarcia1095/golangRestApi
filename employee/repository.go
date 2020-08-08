package employee

import (
	"database/sql"

	"github.com/golangRestApi/helper"
)

type Repository interface {
	GetEmployees(params *getEmployeeRequest) ([]*Employee, error)
	GetTotalEmployees() (int, error)
	GetEmployeeByID(param *getEmployeeByIDRequest) (*Employee, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetEmployees(params *getEmployeeRequest) ([]*Employee, error) {
	const sql = `SELECT id,last_name ,first_name ,company 
	,email_address ,job_title,business_phone 
	,home_phone ,COALESCE(mobile_phone,'') ,fax_number 
	,address 
	FROM employees
	LIMIT ? OFFSET ?`

	results, err := repo.db.Query(sql, params.Limit, params.Offset)

	helper.Catch(err)

	var employees []*Employee
	for results.Next() {
		employee := &Employee{}
		err = results.Scan(&employee.ID, &employee.LastName, &employee.FirstName,
			&employee.Company, &employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone,
			&employee.HomePhone, &employee.MobilePhone, &employee.FaxNumber, &employee.Address,
		)
		if err != nil {
			panic(err)
		}

		employees = append(employees, employee)
	}
	return employees, err
}

func (repo *repository) GetTotalEmployees() (int, error) {
	const sql = `SELECT COUNT(*) FROM employees`

	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)

	return total, nil
}

func (repo *repository) GetEmployeeByID(param *getEmployeeByIDRequest) (*Employee, error) {
	const sql = `SELECT id,last_name ,first_name ,company 
	,email_address ,job_title,business_phone 
	,home_phone ,COALESCE(mobile_phone,'') ,fax_number 
	,address 
	FROM employees
	WHERE id=?`
	row := repo.db.QueryRow(sql, param.EmployeeID)

	employee := &Employee{}

	err := row.Scan(&employee.ID, &employee.LastName, &employee.FirstName,
		&employee.Company, &employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone,
		&employee.HomePhone, &employee.MobilePhone, &employee.FaxNumber, &employee.Address,
	)
	helper.Catch(err)
	return employee, err
}

// func (repo *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
// 	const sql = `SELECT id,product_code,product_name ,COALESCE(description,''),standard_cost ,list_price ,category
// 	FROM products
// 	LIMIT ? OFFSET ?`

// 	results, err := repo.db.Query(sql, params.Limit, params.Offset)

// 	helper.Catch(err)

// 	var products []*Product
// 	for results.Next() {
// 		product := &Product{}
// 		err = results.Scan(&product.ID, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost, &product.ListPrice, &product.Category)
// 		if err != nil {
// 			panic(err)
// 		}

// 		products = append(products, product)
// 	}
// 	return products, err
// }

// func (repo *repository) GetTotalProducts() (int, error) {
// 	const sql = `SELECT COUNT(*) FROM products`

// 	var total int
// 	row := repo.db.QueryRow(sql)
// 	err := row.Scan(&total)
// 	helper.Catch(err)

// 	return total, nil
// }

// func (repo *repository) InsertProduct(params *getAddProductsRequest) (int64, error) {
// 	const sql = `INSERT INTO products
// 			 (product_code,product_name,category,description,list_price,standard_cost)
// 				VALUES(?,?,?,?,?,?)`
// 	result, err := repo.db.Exec(sql, params.ProductCode, params.ProductName, params.Category, params.Description, params.ListPrice,
// 		params.StandardCost)
// 	helper.Catch(err)
// 	id, _ := result.LastInsertId()
// 	return id, nil
// }

// func (repo *repository) UpdateProduct(params *updateProductsRequest) (int64, error) {
// 	const sql = `Update products
// 				SET product_code = ?,
// 				product_name = ?,
// 				category = ?,
// 				description = ?,
// 				list_price = ?,
// 				standard_cost = ?
// 				WHERE id = ?`
// 	_, err := repo.db.Exec(sql, params.ProductCode, params.ProductName, params.Category, params.Description,
// 		params.ListPrice, params.StandardCost, params.ID)
// 	helper.Catch(err)
// 	return params.ID, nil
// }

// func (repo *repository) DeleteProduct(params *deleteProductsRequest) (int64, error) {
// 	const sql = `DELETE from products
// 				WHERE id = ?`
// 	result, err := repo.db.Exec(sql, params.ProductID)
// 	helper.Catch(err)
// 	count, err := result.RowsAffected()
// 	helper.Catch(err)
// 	return count, nil
// }

// func (repo *repository) GetBestSellers() ([]*ProductTop, error) {
// 	const sql = `SELECT
// 	od.product_id ,	p.product_name ,SUM(od.quantity *od.unit_price) vendido
// 	FROM order_details od
// 	INNER JOIN products  p ON od.product_id = p.id
// 	GROUP BY od.product_id
// 	ORDER BY vendido DESC
// 	LIMIT 10`

// 	results, err := repo.db.Query(sql)
// 	helper.Catch(err)

// 	var products []*ProductTop

// 	for results.Next() {
// 		product := &ProductTop{}
// 		err = results.Scan(&product.ID, &product.ProductName, &product.Vendidos)
// 		if err != nil {
// 			panic(err)
// 		}
// 		products = append(products, product)
// 	}
// 	return products, nil
// }

// func (repo *repository) GetTotalVentas() (float64, error) {
// 	const sql = `SELECT SUM(od.quantity*od.unit_price) vendido
// 					FROM order_details od`
// 	var total float64
// 	row := repo.db.QueryRow(sql)
// 	err := row.Scan(&total)
// 	helper.Catch(err)

// 	return total, nil
// }

// func (repo *repository) GetBestsEmployee() (*BestEmployee, error) {
// 	const sql = `SELECT e.id ,
// 	count(e.id ) as totalVentas,
// 	e.first_name ,e.last_name
// 	FROM orders o
// 	INNER JOIN employees e ON o.employee_id = e.id
// 	GROUP BY o.employee_id
// 	ORDER BY totalVentas DESC
// 	LIMIT 1`

// 	row := repo.db.QueryRow(sql)
// 	employee := &BestEmployee{}

// 	err := row.Scan(&employee.ID, &employee.TotalVentas, &employee.FirstName, &employee.LastName)

// 	helper.Catch(err)

// 	return employee, nil
// }

// func (repo *repository) InsertEmployee(params *addEmployeesRequest) (int64, error) {
// 	const sql = `INSERT INTO employees
// 			 (first_name ,last_name ,company ,address ,business_phone,email_address ,
// 			fax_number ,home_phone ,job_title ,mobile_phone)
// 				VALUES(?,?,?,?,?,?,?,?,?,?)`
// 	result, err := repo.db.Exec(sql, params.FirstName, params.LasttName, params.Company,
// 		params.BusinessPhone, params.EmailAddress, params.FaxNumber,
// 		params.HomePhone, params.JobTitle, params.MobilPhone)
// 	helper.Catch(err)
// 	id, _ := result.LastInsertId()
// 	helper.Catch(err)
// 	return id, nil
// }
