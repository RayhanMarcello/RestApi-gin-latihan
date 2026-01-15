package service

import (
	"errors"
	"latihan-api-gorm/entity"
	"latihan-api-gorm/repository"
)

type ProductService interface {
	GetAll() ([]entity.Product, error)
	GetAllById(id int) (entity.Product,error)
	Create(name string, price int) (entity.Product,error)
}

type productService struct{
	repo repository.Repository
}

func NewProductService(repo repository.Repository) *productService{
	return &productService{repo}
}

func (s *productService) GetAll() ([]entity.Product,error){
	return s.repo.FindAll()
} 

func (s *productService) GetAllById(id int) (entity.Product,error){
	if id <= 0 {
		return entity.Product{}, errors.New("id must grater than 0") 
	}
	p, err := s.repo.FindByID(id)
	if err != nil {
		return entity.Product{}, err
	}
	return p,nil
}

func (s *productService) Create(name string, price int)(entity.Product,error){
	if price < 0 {
		return entity.Product{}, errors.New("price cant be less than 0")
	}	
	if name == ""{
		return entity.Product{}, errors.New("product name required!")
	}
	product := entity.Product{
		Name : name,
		Price: price,
	}
	return s.repo.Create(product)
}
