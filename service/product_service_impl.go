package service

import (
	"fmt"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
	"github.com/go-playground/validator/v10"
)

type ProductsServiceImpl struct {
	ProductsRepository repository.ProductsRepository
}

// Delete implements BorrowerService
func (s *ProductsServiceImpl) Delete(id int64) (model.Products, error) {
	return s.ProductsRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *ProductsServiceImpl) FindAll() ([]model.Products, error) {
	return s.ProductsRepository.FindAll()
}

// FindById implements BorrowerService
func (s *ProductsServiceImpl) FindById(id int64) (model.Products, error) {
	fmt.Println("INI ADALAH PRODUCT SERVICE")
	return s.ProductsRepository.FindById(id)
}

// FindByUsername implements BorrowerService
func (s *ProductsServiceImpl) FindByName(name string) (model.Products, error) {
	return s.ProductsRepository.FindByName(name)
}

// Save implements BorrowerService
func (s *ProductsServiceImpl) Save(newProducts model.Products) (model.Products, error) {

	newPro := model.Products{
		Name:        newProducts.Name,
		Installment: newProducts.Installment,
		Bunga:       newProducts.Bunga,
		Created_At:  newProducts.Created_At,
	}
	return s.ProductsRepository.Save(newPro)

}

// Update implements BorrowerService
func (s *ProductsServiceImpl) Update(updateProducts model.Products) (model.Products, error) {

	var pro model.Products
	create_at := pro.Created_At

	newPro := model.Products{
		ID:          updateProducts.ID,
		Name:        updateProducts.Name,
		Installment: updateProducts.Installment,
		Bunga:       updateProducts.Bunga,
		Created_At:  create_at,
	}

	return s.ProductsRepository.Update(newPro)
}

func NewProductsServiceImpl(productsRepository repository.ProductsRepository, validate *validator.Validate) ProductsService {
	return &ProductsServiceImpl{
		ProductsRepository: productsRepository,
	}
}
