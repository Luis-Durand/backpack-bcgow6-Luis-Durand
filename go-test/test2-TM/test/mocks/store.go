package mocks

import (
	"errors"

	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-test/test2-TM/internal/products"
)

type MockStorage struct {
	AllProds []products.Products
	ErrWrite error
	ErrRead  error
}

func (m *MockStorage) Read(data interface{}) error {

	if m.ErrRead != nil {
		return m.ErrRead
	}

	prods, ok := data.(*[]products.Products)
	if !ok {
		return errors.New("error en read")
	}

	*prods = append(*prods, m.AllProds...)

	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.ErrWrite != nil {
		return m.ErrWrite
	}

	prods, ok := data.([]products.Products)
	if !ok {
		return errors.New("error en write")
	}

	m.AllProds = prods

	return nil

}
