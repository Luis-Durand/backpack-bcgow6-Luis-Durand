package products

import "fmt"

type Service interface {
	GetAll() ([]Products, error)
	Store(name, productType string, price int, amount int) (Products, error)
	Update(id int, name, productType string, price int, amount int) (Products, error)
	Delete(id int) error
	UpdateNameAndPrice(id int, name string, price int) (Products, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]Products, error) {
	prods, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return prods, nil

}

func (s *service) Store(name, productType string, price int, amount int) (Products, error) {
	newId, err := s.repository.LastId()
	if err != nil {
		return Products{}, nil
	}
	newId++
	prod, err := s.repository.Store(newId, name, productType, price, amount)
	if err != nil {
		return Products{}, err
	}

	return prod, nil
}
func (s *service) Update(id int, name, productType string, price int, amount int) (Products, error) {
	prod, err := s.repository.Update(id, name, productType, price, amount)
	if err != nil {
		return Products{}, err
	}

	return prod, nil
}

func (s *service) Delete(id int) error {

	err := s.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("Error al eliminar")
	}

	return nil
}

func (s *service) UpdateNameAndPrice(id int, name string, price int) (Products, error) {

	prods, err := s.repository.UpdateNameAndPrice(id, name, price)
	if err != nil {
		return Products{}, fmt.Errorf("No se encontro el producto %d", id)
	}

	return prods, nil

}
