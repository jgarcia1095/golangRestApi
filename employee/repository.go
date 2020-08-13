package employee

import (
	"database/sql"

	"github.com/golangRestApi/helper"
)

type Repository interface {
	GetEmployees(params *getEmployeeRequest) ([]*Employee, error)
	GetTotalEmployees() (int, error)
	GetEmployeeByID(param *getEmployeeByIDRequest) (*Employee, error)
	GetBestEmployee() (*BestEmployee, error)
	InsertEmployee(params *addEmployeeRequest) (int64, error)
	UpdateEmployee(params *updateEmployeeRequest) (int64, error)
	DeleteEmployee(params *deleteEmployeeRequest) (int64, error)
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

func (repo *repository) GetBestEmployee() (*BestEmployee, error) {
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

	helper.Catch(err)

	return employee, nil
}

func (repo *repository) InsertEmployee(params *addEmployeeRequest) (int64, error) {
	const sql = `INSERT INTO employees (last_name ,first_name ,company 
		,email_address ,job_title,business_phone 
		,home_phone ,mobile_phone,fax_number 
		,address)
		VALUES(?,?,?,?,?,?,?,?,?,?)`
	result, err := repo.db.Exec(sql, params.LastName, params.FirstName, params.Company,
		params.EmailAddress, params.JobTitle, params.BusinessPhone, params.HomePhone, params.MobilePhone,
		params.FaxNumber, params.Address,
	)
	helper.Catch(err)
	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) UpdateEmployee(params *updateEmployeeRequest) (int64, error) {
	const sql = `Update employees
				SET last_name = ?,
				first_name = ?,
				company = ?,
				email_address = ?,
				job_title = ?,
				business_phone = ?,
				home_phone = ?,
				mobile_phone = ?,
				fax_number = ?,
				address = ?
				WHERE id = ?`
	_, err := repo.db.Exec(sql, params.LastName, params.FirstName, params.Company, params.EmailAddress,
		params.JobTitle, params.BusinessPhone, params.HomePhone, params.MobilePhone, params.FaxNumber, params.Address, params.ID)
	helper.Catch(err)
	return params.ID, nil
}

func (repo *repository) DeleteEmployee(params *deleteEmployeeRequest) (int64, error) {
	const sql = `DELETE from employees
				WHERE id = ?`
	result, err := repo.db.Exec(sql, params.EmployeeID)
	helper.Catch(err)
	count, err := result.RowsAffected()
	helper.Catch(err)
	return count, nil
}
