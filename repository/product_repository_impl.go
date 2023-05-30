package repository

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

// Delete implements ProductRepository
func (r *ProductRepositoryImpl) Delete(id uint) (model.Product, error) {
	var pro model.Product
	result := r.Db.Where("id = ?", id).Delete(&pro)
	helper.ErrorPanic(result.Error)
	return pro, nil
}

// FindAll implements ProductRepository
func (r *ProductRepositoryImpl) FindAll() ([]model.Product, error) {
	var pro []model.Product
	results := r.Db.Find(&pro)
	helper.ErrorPanic(results.Error)
	return pro, nil
}

// FindById implements ProductRepository
func (r *ProductRepositoryImpl) FindById(id uint) (model.Product, error) {
	var pro model.Product
	result := r.Db.First(&pro, "id = ?", id)
	if result.Error != nil {
		return pro, errors.New("Product is not found")
	}
	return pro, nil
}

// Save implements ProductRepository
func (r *ProductRepositoryImpl) Save(newProduct model.Product) (model.Product, error) {
	currentTime := time.Now()
	newProduct.CreatedAt = currentTime
	result := r.Db.Create(&newProduct)
	helper.ErrorPanic(result.Error)
	return newProduct, nil
}

// Update implements ProductRepository
func (r *ProductRepositoryImpl) Update(updatedProduct model.Product) (model.Product, error) {
	result := r.Db.Model(&model.Product{}).Where("id = ?", updatedProduct.ID).Updates(updatedProduct)
	helper.ErrorPanic(result.Error)
	return updatedProduct, nil
}

// FindByName implements ProductRepository
func (r *ProductRepositoryImpl) FindByName(name string) (model.Product, error) {
	var pro model.Product
	result := r.Db.First(&pro, "name = ?", name)

	if result.Error != nil {
		return pro, errors.New("invalid name")
	}
	return pro, nil
}
