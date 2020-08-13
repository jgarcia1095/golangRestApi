package product

import "github.com/golangRestApi/helper"

type Service interface {
	GetProductByID(param *getProductByIDRequest) (*Product, error)
	GetProducts(params *getProductsRequest) (*ProductList, error)
	InsertProduct(params *getAddProductsRequest) (int64, error)
	UpdateProduct(params *updateProductsRequest) (int64, error)
	DeleteProduct(params *deleteProductsRequest) (int64, error)
	GetBestSellers() (*ProductTopResponse, error)
	InsertEmployee(params *addEmployeesRequest) (int64, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
func (s *service) GetProductByID(param *getProductByIDRequest) (*Product, error) {
	return s.repo.GetProductById(param.ProductID)
}

func (s *service) GetProducts(params *getProductsRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(params)

	helper.Catch(err)
	totalProducts, err := s.repo.GetTotalProducts()

	helper.Catch(err)
	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}

func (s *service) InsertProduct(params *getAddProductsRequest) (int64, error) {
	return s.repo.InsertProduct(params)
}

func (s *service) UpdateProduct(params *updateProductsRequest) (int64, error) {
	return s.repo.UpdateProduct(params)
}

func (s *service) DeleteProduct(params *deleteProductsRequest) (int64, error) {
	return s.repo.DeleteProduct(params)
}

func (s *service) GetBestSellers() (*ProductTopResponse, error) {
	products, err := s.repo.GetBestSellers()
	helper.Catch(err)
	totalVentas, err := s.repo.GetTotalVentas()

	helper.Catch(err)
	return &ProductTopResponse{Data: products, TotalVentas: totalVentas}, nil
}

// func (s *service) GetBestsEmployee() (*BestEmployee, error) {
// 	return s.repo.GetBestsEmployee()
// }

func (s *service) InsertEmployee(params *addEmployeesRequest) (int64, error) {
	return s.repo.InsertEmployee(params)
}
