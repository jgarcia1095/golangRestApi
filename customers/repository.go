package customers

import (
	"database/sql"

	"github.com/golangRestApi/helper"
)

type Repository interface {
	GetCustomers(params *getCustomersRequest) ([]*Customer, error)
	GetTotalCustomers() (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetCustomers(params *getCustomersRequest) ([]*Customer, error) {
	const sql = `SELECT id,first_name ,last_name ,address,business_phone ,city ,company 
	FROM customers 
	LIMIT ? OFFSET ?`

	results, err := repo.db.Query(sql, params.Limit, params.Offset)

	helper.Catch(err)

	var customers []*Customer
	for results.Next() {
		customer := &Customer{}
		err = results.Scan(&customer.ID, &customer.FirstName, &customer.LastName, &customer.Address, &customer.BusinessPhone, &customer.City, &customer.Company)
		if err != nil {
			panic(err)
		}

		customers = append(customers, customer)
	}
	return customers, err
}

func (repo *repository) GetTotalCustomers() (int, error) {
	const sql = `SELECT COUNT(*) FROM customers`

	var total int
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)

	return total, nil
}
