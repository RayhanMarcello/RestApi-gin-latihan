package repository

import (
	"latihan-api-gorm/entity"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entity.Product, error)
	Create(product entity.Product) (entity.Product, error)
	FindByID(id int) (entity.Product, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) FindAll() ([]entity.Product, error){
	var product []entity.Product
	err := r.db.Find(&product).Error
	return product,err
}

func (r *repository) Create(product entity.Product) (entity.Product,error){
	products := entity.Product{
		Name : product.Name,
		Price : product.Price,
	}
	err := r.db.Create(&products).Error
	return products, err
}

func (r *repository) FindByID(id int) (entity.Product,error){
	var product entity.Product
	err := r.db.First(&product, id).Error
	return product, err 
} 

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}