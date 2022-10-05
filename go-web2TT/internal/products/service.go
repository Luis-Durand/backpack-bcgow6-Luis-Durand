package products

type Service interface {
	GetAll() ([]Products, error)
	Store(name, productType string, price int, amount int) (Products, error)
	Update(id int, name, productType string, price int, amount int) (Products, error)
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
