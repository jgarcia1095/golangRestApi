package main

import (
	"database/sql"
	"encoding/json"
	"golangRestApi/database"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

var dataBaseConnection *sql.DB

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dataBaseConnection = database.InitDb()

	r := chi.NewRouter()
	// r.Use(middleware.Logger)
	r.Get("/products", AllProductos)
	r.Post("/products", CreateProductos)
	r.Put("/products/{id}", UpdateProducto)
	r.Delete("/products/{id}", DeleteProducto)
	http.ListenAndServe(":3000", r)
	defer dataBaseConnection.Close()
}

func DeleteProducto(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := dataBaseConnection.Prepare("delete from products where id=?")
	catch(err)

	_, er := query.Exec(id)
	catch(er)
	defer query.Close()

	responseWithJson(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func UpdateProducto(w http.ResponseWriter, r *http.Request) {
	var producto Product
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&producto)

	query, err := dataBaseConnection.Prepare("Update products SET product_code=?, description=? where id=?")
	catch(err)

	_, er := query.Exec(producto.Product_Code, producto.Description, id)
	catch(er)
	defer query.Close()

	responseWithJson(w, http.StatusOK, map[string]string{"message": "successfully updated"})
}

func CreateProductos(w http.ResponseWriter, r *http.Request) {
	var producto Product

	json.NewDecoder(r.Body).Decode(&producto)
	query, err := dataBaseConnection.Prepare("Insert products SET product_code=?, description=?")
	catch(err)

	_, er := query.Exec(producto.Product_Code, producto.Description)
	catch(er)
	defer query.Close()

	responseWithJson(w, http.StatusCreated, map[string]string{"message": "ssuccessfully created"})
}

func AllProductos(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id,product_code,COALESCE(description,'')
				FROM products`
	results, err := dataBaseConnection.Query(sql)
	catch(err)
	var products []*Product

	for results.Next() {
		product := &Product{}
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)

		catch(err)
		products = append(products, product)
	}
	responseWithJson(w, http.StatusOK, products)
}

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
