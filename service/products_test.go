package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/stretchr/testify/assert"
)

type mockProductsRepository struct{}

func (m *mockProductsRepository) Delete(id uint) (model.Product, error) {
	// Simulate deleting a product by ID
	if id == 1 {
		product := model.Product{
			ID:          1,
			Name:        "Product 1",
			Installment: 12,
			Interest:    5,
			CreatedAt:   time.Now(),
		}
		return product, nil
	}
	return model.Product{}, nil
}

func (m *mockProductsRepository) FindAll() ([]model.Product, error) {
	// Simulate finding all products
	products := []model.Product{
		{
			ID:          1,
			Name:        "Product 1",
			Installment: 12,
			Interest:    5,
			CreatedAt:   time.Now(),
		},
		{
			ID:          2,
			Name:        "Product 2",
			Installment: 24,
			Interest:    8,
			CreatedAt:   time.Now(),
		},
	}
	return products, nil
}

func (m *mockProductsRepository) FindById(id uint) (model.Product, error) {
	// Simulate finding a product by ID
	if id == 1 {
		product := model.Product{
			ID:          1,
			Name:        "Product 1",
			Installment: 12,
			Interest:    5,
			CreatedAt:   time.Now(),
		}
		return product, nil
	}
	return model.Product{}, nil
}

func (m *mockProductsRepository) FindByName(name string) (model.Product, error) {
	// Simulate finding a product by name
	if name == "Product 1" {
		product := model.Product{
			ID:          1,
			Name:        "Product 1",
			Installment: 12,
			Interest:    5,
			CreatedAt:   time.Now(),
		}
		return product, nil
	}
	return model.Product{}, nil
}

func (m *mockProductsRepository) Save(newProducts model.Product) (model.Product, error) {
	// Simulate validation errors
	if newProducts.Name == "" {
		return model.Product{}, errors.New("name tidak boleh kosong")
	}
	if newProducts.Installment == 0 {
		return model.Product{}, errors.New("installment tidak boleh kosong")
	}
	if newProducts.Interest == 0 {
		return model.Product{}, errors.New("Interest tidak boleh kosong")
	}

	// Simulate successful save
	return newProducts, nil
}

func (m *mockProductsRepository) Update(updateProducts model.Product) (model.Product, error) {
	// Simulate validation errors
	if updateProducts.Name == "" {
		return model.Product{}, errors.New("name tidak boleh kosong")
	}
	if updateProducts.Installment == 0 {
		return model.Product{}, errors.New("installment tidak boleh kosong")
	}
	if updateProducts.Interest == 0 {
		return model.Product{}, errors.New("Interest tidak boleh kosong")
	}

	// Simulate successful update
	return updateProducts, nil
}

func TestProductsService(t *testing.T) {
	repo := &mockProductsRepository{}
	productService := service.NewProductServiceImpl(repo)

	t.Run("Delete", func(t *testing.T) {
		productID := int64(1)
		product, err := productService.Delete(uint(productID))
		assert.NoError(t, err)
		assert.Equal(t, productID, product.ID)
	})

	t.Run("FindAll", func(t *testing.T) {
		products, err := productService.FindAll()
		assert.NoError(t, err)
		assert.Len(t, products, 2)
	})

	t.Run("FindById", func(t *testing.T) {
		productID := int64(1)
		product, err := productService.FindById(uint(productID))
		assert.NoError(t, err)
		assert.Equal(t, productID, product.ID)
	})

	t.Run("FindByName", func(t *testing.T) {
		productName := "Product 1"
		product, err := productService.FindByName(productName)
		assert.NoError(t, err)
		assert.Equal(t, productName, product.Name)
	})

	t.Run("Save", func(t *testing.T) {
		newProduct := model.Product{
			Name:        "New Product",
			Installment: 12,
			Interest:    5,
			CreatedAt:   time.Now(),
		}

		savedProduct, err := productService.Save(newProduct)
		assert.NoError(t, err)
		assert.Equal(t, newProduct.Name, savedProduct.Name)
	})

	t.Run("Save_ValidationError", func(t *testing.T) {
		invalidProduct := model.Product{
			Name:        "", // Invalid name
			Installment: 12,
			Interest:    5,
			CreatedAt:   time.Now(),
		}

		_, err := productService.Save(invalidProduct)
		assert.Error(t, err)
		assert.EqualError(t, err, "name tidak boleh kosong")
	})

	t.Run("Update", func(t *testing.T) {
		updatedProduct := model.Product{
			ID:          1,
			Name:        "Updated Product",
			Installment: 24,
			Interest:    8,
			CreatedAt:   time.Now(),
		}

		updated, err := productService.Update(updatedProduct)
		assert.NoError(t, err)
		assert.Equal(t, updatedProduct.Name, updated.Name)
	})

	t.Run("Update_ValidationError", func(t *testing.T) {
		invalidProduct := model.Product{
			ID:          1,
			Name:        "", // Invalid name
			Installment: 24,
			Interest:    8,
			CreatedAt:   time.Now(),
		}

		_, err := productService.Update(invalidProduct)
		assert.Error(t, err)
		assert.EqualError(t, err, "name tidak boleh kosong")
	})
}
