package products

type MockService struct {
	Prods        []Product
	ErrorService error
}

func (m *MockService) GetAllBySeller(sellerID string) ([]Product, error) {

	if m.ErrorService != nil {
		return []Product{}, m.ErrorService
	}

	return m.Prods, nil
}
