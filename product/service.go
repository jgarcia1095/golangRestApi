package product

type Service interface {
	GetProductById(param *getProductByIdRequest) (*Product, error)
	GetProducts(params *getProductsRequest) (*ProductList, error)
	InsertProduct(params *getAddProductsRequest) (int64, error)
	GetBestsEmployee() (*BestEmployee, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
func (s *service) GetProductById(param *getProductByIdRequest) (*Product, error) {
	return s.repo.GetProductById(param.ProductId)
}

func (s *service) GetProducts(params *getProductsRequest) (*ProductList, error) {
	products, err := s.repo.GetProducts(params)

	if err != nil {
		panic(err)
	}
	totalProducts, err := s.repo.GetTotalProducts()

	if err != nil {
		panic(err)
	}
	return &ProductList{Data: products, TotalRecords: totalProducts}, nil
}

func (s *service) InsertProduct(params *getAddProductsRequest) (int64, error) {
	return s.repo.InsertProduct(params)
}

func (s *service) GetBestsEmployee() (*BestEmployee, error) {
	return s.repo.GetBestsEmployee()
}
