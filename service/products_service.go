package service

import "github.com/IbnuFarhanS/pinjol/model"

type ProductsService interface {
	Save(newProducts model.Products) (model.Products, error)
	Update(updateProducts model.Products) (model.Products, error)
	Delete(id int64) (model.Products, error)
	FindById(id int64) (model.Products, error)
	FindAll() ([]model.Products, error)
	FindByName(name string) (model.Products, error)
}
