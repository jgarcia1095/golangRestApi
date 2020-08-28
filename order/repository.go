package order

import (
	"database/sql"
	"fmt"

	"github.com/golangRestApi/helper"
)

type Repository interface {
	GetOrderByID(params *getOrderByIDRequest) (*OrderItem, error)
	GetOrders(params *getOrdersRequest) ([]*OrderItem, error)
	GetTotalOrders(params *getOrdersRequest) (int64, error)
	InsertOrder(params *addOrderRequest) (int64, error)
	InsertOrderDetail(params *addOrderDetailRequest) (int64, error)
	UpdateOrder(params *addOrderRequest) (int64, error)
	UpdateOrderDetail(params *addOrderDetailRequest) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(databaseConnection *sql.DB) Repository {
	return &repository{db: databaseConnection}
}

func (repo *repository) GetOrderByID(params *getOrderByIDRequest) (*OrderItem, error) {
	const sql = `SELECT o.id ,o.customer_id ,o.order_date ,o.status_id ,
	os.status_name ,CONCAT(c.first_name,' ',c.last_name) as customer_name,
	c.company ,
	c.address ,
	c.business_phone,
	c.city 
	FROM orders o 
	INNER JOIN orders_status os 	
	ON o.status_id = os.id 
	INNER JOIN customers c 
	ON o.customer_id  = c.id 
	WHERE o.id =?`

	order := &OrderItem{}

	row := repo.db.QueryRow(sql, params.OrderID)

	err := row.Scan(&order.ID, &order.CustomerID, &order.OrderDate, &order.StatusID, &order.StatusName,
		&order.Customer, &order.Company, &order.Address, &order.Phone, &order.City)
	helper.Catch(err)

	orderDetail, err := GetOrderDetail(repo, &params.OrderID)
	helper.Catch(err)

	order.Data = orderDetail

	return order, nil
}

func GetOrderDetail(repo *repository, orderID *int64) ([]*OrderDetailItem, error) {
	const sql = `SELECT order_id ,od.id ,quantity ,unit_price ,
	p.product_name ,product_id 
	FROM order_details od 
	INNER JOIN products p
	ON od.product_id = p.id 
	WHERE od.order_id =?`

	results, err := repo.db.Query(sql, orderID)
	helper.Catch(err)

	var orders []*OrderDetailItem

	for results.Next() {
		order := &OrderDetailItem{}
		err = results.Scan(&order.OrderID, &order.ID, &order.Quantity, &order.UnitPrice,
			&order.ProductName, &order.ProductID)
		helper.Catch(err)

		orders = append(orders, order)
	}

	return orders, nil
}

func (repo *repository) GetOrders(params *getOrdersRequest) ([]*OrderItem, error) {
	var filter string

	if params.Status != nil {
		filter += fmt.Sprintf(" AND o.status_id = %v ", params.Status.(float64))
	}

	if params.DateFrom != nil && params.DateTo == nil {
		filter += fmt.Sprintf(" AND o.order_date >= '%v' ", params.DateFrom.(string))
	}

	if params.DateFrom == nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date <= '%v' ", params.DateTo.(string))
	}

	if params.DateFrom != nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date between '%v' and '%v' ", params.DateFrom.(string), params.DateTo.(string))
	}

	var sql = `SELECT o.id ,
	o.customer_id ,
	o.order_date ,
	o.status_id ,
	os.status_name ,
	CONCAT(c.first_name , ' ', c.last_name ) as customer_name
	FROM orders o 
	INNER JOIN orders_status os ON o.status_id  = os.id 
	INNER JOIN customers c ON o.customer_id = c.id 
	WHERE 1=1`

	sql = sql + filter + " LIMIT ? OFFSET ?"

	results, err := repo.db.Query(sql, params.Limit, params.Offset)

	helper.Catch(err)

	var orders []*OrderItem

	for results.Next() {
		order := &OrderItem{}
		err = results.Scan(&order.ID, &order.CustomerID, &order.OrderDate, &order.StatusID,
			&order.StatusName, &order.Customer,
		)
		helper.Catch(err)

		orderDetail, err := GetOrderDetail(repo, &order.ID)
		helper.Catch(err)

		order.Data = orderDetail
		orders = append(orders, order)
	}
	return orders, err

}

func (repo *repository) GetTotalOrders(params *getOrdersRequest) (int64, error) {
	var filter string

	if params.Status != nil {
		filter += fmt.Sprintf(" AND o.status_id = %v ", params.Status.(float64))
	}

	if params.DateFrom != nil && params.DateTo == nil {
		filter += fmt.Sprintf(" AND o.order_date >= '%v' ", params.DateFrom.(string))
	}

	if params.DateFrom == nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date <= '%v' ", params.DateTo.(string))
	}

	if params.DateFrom != nil && params.DateTo != nil {
		filter += fmt.Sprintf(" AND o.order_date between '%v' and '%v'", params.DateFrom.(string), params.DateTo.(string))
	}

	var sql = `SELECT COUNT(*)
	FROM orders o
	WHERE 1=1` + filter

	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)

	return total, nil
}

func (repo *repository) InsertOrder(params *addOrderRequest) (int64, error) {

	var sql = `INSERT INTO orders
	(customer_id, order_date)
	VALUES(?,?)`

	result, err := repo.db.Exec(sql, params.CustomerID, params.OrderDate)
	helper.Catch(err)
	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) InsertOrderDetail(params *addOrderDetailRequest) (int64, error) {

	var sql = `INSERT INTO order_details
	(order_id, product_id, quantity, unit_price)
	VALUES(?,?,?,?)`

	result, err := repo.db.Exec(sql, params.OrderID, params.ProductID, params.Quantity, params.UnitPrice)
	helper.Catch(err)
	id, _ := result.LastInsertId()
	return id, nil
}

func (repo *repository) UpdateOrder(params *addOrderRequest) (int64, error) {
	var sql = `UPDATE orders
	SET customer_id = ?
	WHERE id =?`

	_, err := repo.db.Exec(sql, params.CustomerID, params.ID)
	helper.Catch(err)
	return params.ID, nil
}

func (repo *repository) UpdateOrderDetail(params *addOrderDetailRequest) (int64, error) {
	var sql = `UPDATE order_details
	SET quantity = ?,
	unit_price = ?
	WHERE id =?`

	_, err := repo.db.Exec(sql, params.Quantity, params.UnitPrice, params.ID)
	helper.Catch(err)
	return params.ID, nil
}
