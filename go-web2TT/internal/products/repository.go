package products

import "fmt"

type Products struct {
	Id     int    `json:"id"`
	Name   string `json:"name" `
	Type   string `json:"type" `
	Price  int    `json:"price" `
	Amount int    `json:"amoun" `
}

var newArr []Products
var lastId int

type Repository interface {
	GetAll() ([]Products, error)
	LastId() (int, error)
	Store(id int, name, productType string, price int, amount int) (Products, error)
	Update(id int, name, productType string, price int, amount int) (Products, error)
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Products, error) {

	return newArr, nil

}

func (r *repository) LastId() (int, error) {

	return lastId, nil
}

func (r *repository) Store(id int, name, productType string, price int, amount int) (Products, error) {

	newProd := Products{id, name, productType, price, amount}
	newArr = append(newArr, newProd)
	lastId = newProd.Id
	return newProd, nil
}

func (r *repository) Update(id int, name, productType string, price int, amount int) (Products, error) {

	newProd := Products{Name: name, Type: productType, Price: price, Amount: amount}
	updated := false

	for i := range newArr {
		if newArr[i].Id == id {
			newProd.Id = id
			newArr[i] = newProd
			updated = true
		}
	}

	if !updated {
		return Products{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return newProd, nil
}
