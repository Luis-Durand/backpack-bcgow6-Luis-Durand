package products

import (
	"fmt"

	"github.com/Luis-Durand/backpack-bcgow6-Luis-Durand/go-web2TT/pkg/store"
)

type Products struct {
	Id     int    `json:"id"`
	Name   string `json:"name" `
	Type   string `json:"type" `
	Price  int    `json:"price" `
	Amount int    `json:"amount" `
}

type Repository interface {
	GetAll() ([]Products, error)
	LastId() (int, error)
	Store(id int, name, productType string, price int, amount int) (Products, error)
	Update(id int, name, productType string, price int, amount int) (Products, error)
	Delete(id int) error
	UpdateNameAndPrice(id int, name string, price int) (Products, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]Products, error) {
	var prods []Products
	if err := r.db.Read(&prods); err != nil {
		return nil, err
	}

	return prods, nil

}

func (r *repository) LastId() (int, error) {

	var arrProd []Products
	if err := r.db.Read(&arrProd); err != nil {
		return 0, nil
	}

	if len(arrProd) == 0 {
		return 0, nil
	}

	return arrProd[len(arrProd)-1].Id, nil
}

func (r *repository) Store(id int, name, productType string, price int, amount int) (Products, error) {

	var ps []Products
	err := r.db.Read(&ps)
	if err != nil {
		return Products{}, nil
	}

	newProd := Products{id, name, productType, price, amount}
	ps = append(ps, newProd)
	if err := r.db.Write(ps); err != nil {
		return Products{}, nil
	}
	return newProd, nil
}

func (r *repository) Update(id int, name, productType string, price int, amount int) (Products, error) {

	var arrProd []Products

	r.db.Read(&arrProd)

	newProd := Products{Name: name, Type: productType, Price: price, Amount: amount}
	updated := false

	for i := range arrProd {
		if arrProd[i].Id == id {
			newProd.Id = id
			arrProd[i] = newProd
			updated = true
		}
	}

	if err := r.db.Write(arrProd); err != nil {
		return Products{}, nil
	}

	if !updated {
		return Products{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return newProd, nil
}

func (r *repository) Delete(id int) error {

	var arrProds []Products

	r.db.Read(&arrProds)

	deleted := false
	var index int
	for i := range arrProds {
		if arrProds[i].Id == id {
			index = i
			deleted = true
			break
		}

	}

	if !deleted {

		return fmt.Errorf("Error %d no encontrado", id)

	}

	arrProds = append(arrProds[:index], arrProds[index+1:]...)

	if err := r.db.Write(arrProds); err != nil {
		return err
	}

	return nil

}

func (r *repository) UpdateNameAndPrice(id int, name string, price int) (Products, error) {

	var arrProd []Products

	r.db.Read(&arrProd)
	updated := false
	var updatedProd Products
	for i := range arrProd {
		if arrProd[i].Id == id {
			arrProd[i].Name = name
			arrProd[i].Price = price
			updatedProd = arrProd[i]
			updated = true
			break
		}

	}
	if !updated {
		return Products{}, fmt.Errorf("No se encontro el producto %d", id)
	}

	if err := r.db.Read(arrProd); err != nil {
		return Products{}, err
	}

	return updatedProd, nil

}
