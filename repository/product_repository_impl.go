package repository

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

type ProductsRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductsRepositoryImpl(Db *gorm.DB) ProductsRepository {
	return &ProductsRepositoryImpl{Db: Db}
}

// Delete implements ProductsRepository
func (r *ProductsRepositoryImpl) Delete(id int64) (model.Products, error) {
	var pro model.Products
	result := r.Db.Where("id = ?", id).Delete(&pro)
	helper.ErrorPanic(result.Error)
	return pro, nil
}

// FindAll implements ProductsRepository
func (r *ProductsRepositoryImpl) FindAll() ([]model.Products, error) {
	var pro []model.Products
	results := r.Db.Find(&pro)
	helper.ErrorPanic(results.Error)
	return pro, nil
}

// FindById implements ProductsRepository
func (r *ProductsRepositoryImpl) FindById(id int64) (model.Products, error) {
	var pro model.Products
	result := r.Db.Find(&pro, "id = ?", id)
	if result.Error != nil {
		return pro, errors.New("products is not found")
	}
	return pro, nil
}

// Save implements ProductsRepository
func (r *ProductsRepositoryImpl) Save(newProducts model.Products) (model.Products, error) {
	currentTime := time.Now()
	newProducts.Created_At = currentTime
	result := r.Db.Create(&newProducts)
	helper.ErrorPanic(result.Error)
	return newProducts, nil
}

// Update implements ProductsRepository
func (r *ProductsRepositoryImpl) Update(updatedProducts model.Products) (model.Products, error) {
	result := r.Db.Model(&model.Products{}).Where("id = ?", updatedProducts.ID).Updates(updatedProducts)
	helper.ErrorPanic(result.Error)
	return updatedProducts, nil
}

// FindByName implements ProductsRepository
func (r *ProductsRepositoryImpl) FindByName(name string) (model.Products, error) {
	var pro model.Products
	result := r.Db.First(&pro, "name = ?", name)

	if result.Error != nil {
		return pro, errors.New("invalid name")
	}
	return pro, nil
}
