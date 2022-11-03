package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockedRepo struct {
	products []Product
}

func (m *mockedRepo) GetAllBySeller(sellerId string) ([]Product, error) {

	var allProds []Product

	for _, prod := range m.products {

		if prod.SellerID == sellerId {
			allProds = append(allProds, prod)

		}
	}

	if sellerId == "" {
		return []Product{}, errors.New("sellerId is required")
	}

	if len(m.products) == 0 {
		return []Product{}, errors.New("prods is empty")
	}

	return allProds, nil

}

func TestGetAllBySeller(t *testing.T) {

	sellId := "FEX112AC"

	allProds := []Product{{ID: "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55},
		{
			ID:          "stub",
			SellerID:    "FEX112AC",
			Description: "generics product",
			Price:       123232.55}}

	repo := mockedRepo{products: allProds}

	serv := NewService(&repo)

	prods, err := serv.GetAllBySeller(sellId)

	assert.Nil(t, err)
	assert.Equal(t, allProds, prods)

}

func TestGetAllBySellerFail(t *testing.T) {

	var expected []Product
	repo := NewRepository()
	serv := NewService(repo)

	prods, err := serv.GetAllBySeller("")

	assert.NotNil(t, err)
	assert.Equal(t, expected, prods)

}
