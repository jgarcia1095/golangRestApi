package product

import "database/sql"

type Repository interface {
	GetProductById(product int) (*Product, error)
	GetProducts(params *getProductsRequest) ([]*Product, error)
	GetTotalProducts() (int, error)
	InsertProduct(params *getAddProductsRequest) (int64, error)
	UpdateProduct(params *updateProductsRequest) (int64, error)
	DeleteProduct(params *deleteProductsRequest) (int64, error)
	GetBestsEmployee() (*BestEmployee, error)
	InsertEmployee(params *addEmployeesRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetProductById(productId int) (*Product, error) {
	const sql = `SELECT id,product_code,product_name,COALESCE(description,''),standard_cost,list_price,category
		FROM products
		WHERE id=?`
	row := repo.db.QueryRow(sql, productId)

	product := &Product{}

	err := row.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost, &product.ListPrice, &product.Category)
	if err != nil {
		panic(err)
	}
	return product, err
}

func (repo *repository) GetProducts(params *getProductsRequest) ([]*Product, error) {
	const sql = `SELECT id,product_code,product_name ,COALESCE(description,''),standard_cost ,list_price ,category 
	FROM products
	LIMIT ? OFFSET ?`

	results, err := repo.db.Query(sql, params.Limit, params.Offset)

	if err != nil {
		panic(err)
	}

	var products []*Product
	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.Id, &product.ProductCode, &product.ProductName, &product.Description, &product.StandardCost, &product.ListPrice, &product.Category)
		if err != nil {
			panic(err)
		}

		products = append(products, product)
	}
	return products, err
}

func (repo *repository) GetTotalProducts() (int, error) {
	const sql = `SELECT COUNT(*) FROM products`

	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	if err != nil {
		panic(err)
	}

	return total, nil
}

func (repo *repository) InsertProduct(params *getAddProductsRequest) (int64, error) {
	const sql = `INSERT INTO products
			 (product_code,product_name,category,description,list_price,standard_cost)
				VALUES(?,?,?,?,?,?)`
	result, err := repo.db.Exec(sql, params.ProductCode, params.ProductName, params.Category, params.Description, params.ListPrice,
		params.StandardCost)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) UpdateProduct(params *updateProductsRequest) (int64, error) {
	const sql = `Update products
				SET product_code = ?,
				product_name = ?,
				category = ?,
				description = ?,
				list_price = ?,
				standard_cost = ?
				WHERE id = ?`
	_, err := repo.db.Exec(sql, params.ProductCode, params.ProductName, params.Category, params.Description,
		params.ListPrice, params.StandardCost, params.ID)
	if err != nil {
		panic(err)
	}
	return params.ID, nil
}

func (repo *repository) DeleteProduct(params *deleteProductsRequest) (int64, error) {
	const sql = `DELETE from products
				WHERE id = ?`
	result, err := repo.db.Exec(sql, params.ProductID)
	if err != nil {
		panic(err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return count, nil
}

func (repo *repository) GetBestsEmployee() (*BestEmployee, error) {
	const sql = `SELECT e.id ,
	count(e.id ) as totalVentas,
	e.first_name ,e.last_name 
	FROM orders o
	INNER JOIN employees e ON o.employee_id = e.id 
	GROUP BY o.employee_id 
	ORDER BY totalVentas DESC 
	LIMIT 1`

	row := repo.db.QueryRow(sql)
	employee := &BestEmployee{}

	err := row.Scan(&employee.ID, &employee.TotalVentas, &employee.FirstName, &employee.LastName)

	if err != nil {
		panic(err)
	}

	return employee, nil
}

func (repo *repository) InsertEmployee(params *addEmployeesRequest) (int64, error) {
	const sql = `INSERT INTO employees
			 (first_name ,last_name ,company ,address ,business_phone,email_address ,
			fax_number ,home_phone ,job_title ,mobile_phone)
				VALUES(?,?,?,?,?,?,?,?,?,?)`
	result, err := repo.db.Exec(sql, params.FirstName, params.LasttName, params.Company,
		params.BusinessPhone, params.EmailAddress, params.FaxNumber,
		params.HomePhone, params.JobTitle, params.MobilPhone)
	if err != nil {
		panic(err)
	}
	id, _ := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return id, nil
}
