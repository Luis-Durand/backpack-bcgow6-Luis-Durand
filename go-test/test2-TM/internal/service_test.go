package products

import (
	"errors"
	"testing"

	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-test/test2-TM/internal/products"
	"github.com/stretchr/testify/assert"
)

type MockService struct {
	product       products.Products
	sliceProds    []products.Products
	ReadWasCalled bool
}

func (m *MockService) Read(data interface{}) error {
	m.ReadWasCalled = true
	prods, ok := data.(*[]products.Products)
	if !ok {
		return errors.New("read failed")
	}

	*prods = append(*prods, m.sliceProds...)

	return nil

}

func (m *MockService) Write(data interface{}) error {

	produs, ok := data.([]products.Products)
	if !ok {
		return errors.New("read failed")
	}

	m.sliceProds = produs
	return nil

}

func TestUpdate(t *testing.T) {

	oneProd := products.Products{
		Id:     2,
		Name:   "Magdalenas",
		Type:   "Arcor",
		Price:  310,
		Amount: 1,
	}

	newProd := []products.Products{{
		Id:     1,
		Name:   "Lays",
		Type:   "Snacks",
		Price:  210,
		Amount: 1,
	}, {
		Id:     2,
		Name:   "Pepitos",
		Type:   "Galletitas",
		Price:  150,
		Amount: 1,
	}}

	mockDb := MockService{product: oneProd, sliceProds: newProd}

	repo := products.NewRepository(&mockDb)
	serv := products.NewService(repo)

	//act
	prods, err := serv.Update(2, "Magdalenas", "Arcor", 310, 1)

	//assert

	assert.Nil(t, err)
	assert.Equal(t, oneProd, prods)
	assert.True(t, mockDb.ReadWasCalled)
}

func TestDelete(t *testing.T) {

	tee := []products.Products{
		{
			Id:     2,
			Name:   "Leche",
			Type:   "Serenisima",
			Price:  230,
			Amount: 1,
		},
	}

	oneProd := products.Products{
		Id:     1,
		Name:   "Lays",
		Type:   "Snacks",
		Price:  210,
		Amount: 1,
	}

	newProd := []products.Products{{
		Id:     1,
		Name:   "Lays",
		Type:   "Snacks",
		Price:  210,
		Amount: 1,
	}, {
		Id:     2,
		Name:   "Leche",
		Type:   "Serenisima",
		Price:  230,
		Amount: 1,
	}}

	mockDb := MockService{product: oneProd, sliceProds: newProd}

	repo := products.NewRepository(&mockDb)
	serv := products.NewService(repo)

	//act
	err := serv.Delete(1)

	//assert
	assert.Nil(t, err)
	assert.True(t, mockDb.ReadWasCalled)
	assert.Equal(t, tee, mockDb.sliceProds)

	errs := serv.Delete(5)

	assert.NotNil(t, errs)
	assert.True(t, mockDb.ReadWasCalled)
	assert.Equal(t, tee, mockDb.sliceProds)

}
