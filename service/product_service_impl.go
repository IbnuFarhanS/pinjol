package service

import (
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/repository"
)

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
}

// Delete implements BorrowerService
func (s *ProductServiceImpl) Delete(id uint) (model.Product, error) {
	return s.ProductRepository.Delete(id)
}

// FindAll implements BorrowerService
func (s *ProductServiceImpl) FindAll() ([]model.Product, error) {
	return s.ProductRepository.FindAll()
}

// FindById implements BorrowerService
func (s *ProductServiceImpl) FindById(id uint) (model.Product, error) {
	return s.ProductRepository.FindById(id)
}

// FindByUsername implements BorrowerService
func (s *ProductServiceImpl) FindByName(name string) (model.Product, error) {
	return s.ProductRepository.FindByName(name)
}

// Save implements BorrowerService
func (s *ProductServiceImpl) Save(newProduct model.Product) (model.Product, error) {

	newPro := model.Product{
		Name:        newProduct.Name,
		Installment: newProduct.Installment,
		Interest:    newProduct.Interest,
		CreatedAt:   newProduct.CreatedAt,
	}
	return s.ProductRepository.Save(newPro)

}

// Update implements BorrowerService
func (s *ProductServiceImpl) Update(updateProduct model.Product) (model.Product, error) {

	var pro model.Product
	create_at := pro.CreatedAt

	newPro := model.Product{
		ID:          updateProduct.ID,
		Name:        updateProduct.Name,
		Installment: updateProduct.Installment,
		Interest:       updateProduct.Interest,
		CreatedAt:  create_at,
	}

	return s.ProductRepository.Update(newPro)
}

func NewProductServiceImpl(ProductRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		ProductRepository: ProductRepository,
	}
}
