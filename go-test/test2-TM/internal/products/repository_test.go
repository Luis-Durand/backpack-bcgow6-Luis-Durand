package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubDB struct {
	newProds []Products
}

type MockDB struct {
	ReadWasCalled bool
	BeforeUpdate  Products
	AfterUpdate   Products
	sliceProds    []Products
}

func (d StubDB) Read(data interface{}) error {

	prods, ok := data.(*[]Products)
	if !ok {
		return errors.New("read failed")
	}

	*prods = append(*prods, d.newProds...)

	return nil

}

func (d StubDB) Write(data interface{}) error {

	return nil

}

func TestRead(t *testing.T) {

	// arrange
	prods := []Products{
		{
			Id:     1,
			Name:   "Yogurt",
			Type:   "Lacteo",
			Price:  210,
			Amount: 1,
		},
		{
			Id:     2,
			Name:   "Jorgito",
			Type:   "Alfajor",
			Price:  150,
			Amount: 1,
		},
	}

	expectedProds := []Products{
		{
			Id:     1,
			Name:   "Yogurt",
			Type:   "Lacteo",
			Price:  210,
			Amount: 1,
		},
		{
			Id:     2,
			Name:   "Jorgito",
			Type:   "Alfajor",
			Price:  150,
			Amount: 1,
		},
	}

	stubDb := StubDB{prods}
	repo := NewRepository(stubDb)
	serv := NewService(repo)

	//Act

	allProds, err := serv.GetAll()

	// Asserts
	assert.Nil(t, err)
	assert.Equal(t, expectedProds, allProds)
}

func (m *MockDB) Read(data interface{}) error {
	m.ReadWasCalled = true
	arrProds, ok := data.(*[]Products)
	if !ok {
		return errors.New("Hubo un fallo en read")
	}

	*arrProds = append(*arrProds, m.sliceProds...)

	return nil

}

func (m *MockDB) Write(data interface{}) error {

	arrProds, ok := data.([]Products)
	if !ok {
		return errors.New("Hubo un fallo en read")
	}

	arrProds = append(arrProds, m.sliceProds...)

	return nil
}

func TestUpdate(t *testing.T) {

	//arrange

	firstProds := Products{
		1, "Pepsi", "Gaseosa", 200, 1,
	}

	updatedFirstProd := Products{
		1, "Coca col4", "Gaseosa", 300, 2,
	}

	mockDb := MockDB{
		ReadWasCalled: false,
		BeforeUpdate:  firstProds,
		AfterUpdate:   updatedFirstProd,
		sliceProds:    []Products{firstProds}}

	repo := NewRepository(&mockDb)
	serv := NewService(repo)

	//act
	prod, err := serv.Update(updatedFirstProd.Id, updatedFirstProd.Name, updatedFirstProd.Type, updatedFirstProd.Price, updatedFirstProd.Amount)

	//assertion

	assert.Nil(t, err)
	assert.Equal(t, updatedFirstProd, prod)
	assert.True(t, mockDb.ReadWasCalled)
}
