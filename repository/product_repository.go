package repository

import "github.com/IbnuFarhanS/pinjol/model"

type ProductRepository interface {
	Save(newProduct model.Product) (model.Product, error)
	Update(updateProduct model.Product) (model.Product, error)
	Delete(id uint) (model.Product, error)
	FindById(id uint) (model.Product, error)
	FindAll() ([]model.Product, error)
	FindByName(name string) (model.Product, error)
}
