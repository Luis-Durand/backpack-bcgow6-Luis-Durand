package products

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubDB struct {
	newProds []Products
	newErr   error
}

type MockDB struct {
	ReadWasCalled bool
	BeforeUpdate  Products
	AfterUpdate   Products
	sliceProds    []Products
	errRead       error
	errWrite      error
}

func (d StubDB) Read(data interface{}) error {
	if d.newErr != nil {
		return d.newErr
	}
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

func TestGetAll(t *testing.T) {

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

	stubDb := StubDB{prods, nil}
	repo := NewRepository(stubDb)

	//Act

	allProds, err := repo.GetAll()

	// Asserts
	assert.Nil(t, err)
	assert.Equal(t, expectedProds, allProds)
}

func TestGetAllFail(t *testing.T) {
	//arrange
	var prodEmpty []Products
	getError := errors.New("Fallo en el get")
	db := StubDB{prodEmpty, getError}

	repo := NewRepository(db)

	// act
	prods, err := repo.GetAll()

	//assert
	assert.EqualError(t, getError, err.Error())
	assert.Equal(t, prodEmpty, prods)

}

func (m *MockDB) Read(data interface{}) error {
	if m.errRead != nil {
		return m.errRead
	}
	m.ReadWasCalled = true
	arrProds, ok := data.(*[]Products)
	if !ok {
		return errors.New("Hubo un fallo en read")
	}

	*arrProds = append(*arrProds, m.sliceProds...)

	return nil

}

func (m *MockDB) Write(data interface{}) error {
	if m.errWrite != nil {
		return m.errWrite
	}
	arrProds, ok := data.([]Products)
	if !ok {
		return errors.New("Hubo un fallo en read")
	}

	m.sliceProds = arrProds

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

	//act
	prod, err := repo.Update(updatedFirstProd.Id, updatedFirstProd.Name, updatedFirstProd.Type, updatedFirstProd.Price, updatedFirstProd.Amount)

	//assertion
	//PROBAR CON ID Q NO EXISTE , FALLO EN READ Y WRITE
	assert.Nil(t, err)
	assert.Equal(t, updatedFirstProd, prod)
	assert.True(t, mockDb.ReadWasCalled)
}

func TestUpdateFailWithId(t *testing.T) {

	// arrange
	emptyProds := Products{}
	mockDB := MockDB{}
	repo := NewRepository(&mockDB)

	prod, err := repo.Update(3, "oreo", "galletita", 300, 1)

	assert.NotNil(t, err)
	assert.Equal(t, emptyProds, prod)

}

func TestUpdateFailRead(t *testing.T) {

	errRead := errors.New("Ocurrio un error en read")
	mockDB := MockDB{errRead: errRead}
	repo := NewRepository(&mockDB)

	prods, err := repo.Update(5, "cocaco", "gaseosa", 300, 1)

	assert.EqualError(t, errRead, err.Error())
	assert.Equal(t, Products{}, prods)

}

func TestUpdateFailWrite(t *testing.T) {

	errWrite := errors.New("Ocurrio un error en write")
	mockDB := MockDB{errWrite: errWrite}
	repo := NewRepository(&mockDB)

	data, err := repo.Update(1, "sw", "es", 1, 2)

	assert.EqualError(t, errWrite, err.Error())
	assert.Equal(t, Products{}, data)

}

func TestDeleteFailWrite(t *testing.T) {

	errW := errors.New("Error en write papu")

	prods := []Products{
		{Id: 1,
			Name:   "jorgito",
			Type:   "alfajor",
			Amount: 1,
			Price:  1,
		},
	}

	mockDB := MockDB{errWrite: errW, sliceProds: prods}

	repo := NewRepository(&mockDB)

	err := repo.Delete(1)

	assert.NotNil(t, err)
	assert.Equal(t, errW, err)

}

func TestDeleteFailRead(t *testing.T) {

	errR := errors.New("Error en read papu")

	prods := []Products{
		{Id: 1,
			Name:   "jorgito",
			Type:   "alfajor",
			Amount: 1,
			Price:  1,
		},
	}

	mockDB := MockDB{errRead: errR, sliceProds: prods}

	repo := NewRepository(&mockDB)

	err := repo.Delete(1)

	assert.NotNil(t, err)
	assert.Equal(t, errR, err)

}
