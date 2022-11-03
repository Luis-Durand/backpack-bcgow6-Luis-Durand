package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockService struct {
	product       Products
	sliceProds    []Products
	errWrite      error
	errRead       error
	ReadWasCalled bool
}

func (m *MockService) Read(data interface{}) error {
	if m.errRead != nil {
		return m.errRead
	}
	m.ReadWasCalled = true
	prods, ok := data.(*[]Products)
	if !ok {
		return errors.New("read failed")
	}

	*prods = append(*prods, m.sliceProds...)

	return nil

}

func (m *MockService) Write(data interface{}) error {
	if m.errWrite != nil {
		return m.errWrite
	}
	produs, ok := data.([]Products)
	if !ok {
		return errors.New("read failed")
	}

	m.sliceProds = produs
	return nil

}

func TestUpdatee(t *testing.T) {

	oneProd := Products{
		Id:     2,
		Name:   "Magdalenas",
		Type:   "Arcor",
		Price:  310,
		Amount: 1,
	}

	newProd := []Products{{
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

	repo := NewRepository(&mockDb)
	serv := NewService(repo)

	//act
	prods, err := serv.Update(2, "Magdalenas", "Arcor", 310, 1)

	//assert

	assert.Nil(t, err)
	assert.Equal(t, oneProd, prods)
	assert.True(t, mockDb.ReadWasCalled)

}

func TestUpdateFail(t *testing.T) {

	mockDB := MockDB{sliceProds: []Products{}}
	repo := NewRepository(&mockDB)
	serv := NewService(repo)

	prod, err := serv.Update(2, "rumba", "galletita", 300, 1)

	assert.NotNil(t, err)
	assert.Equal(t, Products{}, prod)
}

func TestDelete(t *testing.T) {

	tee := []Products{
		{
			Id:     2,
			Name:   "Leche",
			Type:   "Serenisima",
			Price:  230,
			Amount: 1,
		},
	}

	oneProd := Products{
		Id:     1,
		Name:   "Lays",
		Type:   "Snacks",
		Price:  210,
		Amount: 1,
	}

	newProd := []Products{{
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

	repo := NewRepository(&mockDb)
	serv := NewService(repo)

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
