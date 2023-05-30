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

func (m *mockProductsRepository) Delete(id int64) (model.Products, error) {
	// Simulate deleting a product by ID
	if id == 1 {
		product := model.Products{
			ID:          1,
			Name:        "Product 1",
			Installment: 12,
			Bunga:       5,
			Created_At:  time.Now(),
		}
		return product, nil
	}
	return model.Products{}, nil
}

func (m *mockProductsRepository) FindAll() ([]model.Products, error) {
	// Simulate finding all products
	products := []model.Products{
		{
			ID:          1,
			Name:        "Product 1",
			Installment: 12,
			Bunga:       5,
			Created_At:  time.Now(),
		},
		{
			ID:          2,
			Name:        "Product 2",
			Installment: 24,
			Bunga:       8,
			Created_At:  time.Now(),
		},
	}
	return products, nil
}

func (m *mockProductsRepository) FindById(id int64) (model.Products, error) {
	// Simulate finding a product by ID
	if id == 1 {
		product := model.Products{
			ID:          1,
			Name:        "Product 1",
			Installment: 12,
			Bunga:       5,
			Created_At:  time.Now(),
		}
		return product, nil
	}
	return model.Products{}, nil
}

func (m *mockProductsRepository) FindByName(name string) (model.Products, error) {
	// Simulate finding a product by name
	if name == "Product 1" {
		product := model.Products{
			ID:          1,
			Name:        "Product 1",
			Installment: 12,
			Bunga:       5,
			Created_At:  time.Now(),
		}
		return product, nil
	}
	return model.Products{}, nil
}

func (m *mockProductsRepository) Save(newProducts model.Products) (model.Products, error) {
	// Simulate validation errors
	if newProducts.Name == "" {
		return model.Products{}, errors.New("name tidak boleh kosong")
	}
	if newProducts.Installment == 0 {
		return model.Products{}, errors.New("installment tidak boleh kosong")
	}
	if newProducts.Bunga == 0 {
		return model.Products{}, errors.New("bunga tidak boleh kosong")
	}

	// Simulate successful save
	return newProducts, nil
}

func (m *mockProductsRepository) Update(updateProducts model.Products) (model.Products, error) {
	// Simulate validation errors
	if updateProducts.Name == "" {
		return model.Products{}, errors.New("name tidak boleh kosong")
	}
	if updateProducts.Installment == 0 {
		return model.Products{}, errors.New("installment tidak boleh kosong")
	}
	if updateProducts.Bunga == 0 {
		return model.Products{}, errors.New("bunga tidak boleh kosong")
	}

	// Simulate successful update
	return updateProducts, nil
}

func TestProductsService(t *testing.T) {
	repo := &mockProductsRepository{}
	productService := service.NewProductsServiceImpl(repo)

	t.Run("Delete", func(t *testing.T) {
		productID := int64(1)
		product, err := productService.Delete(productID)
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
		product, err := productService.FindById(productID)
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
		newProduct := model.Products{
			Name:        "New Product",
			Installment: 12,
			Bunga:       5,
			Created_At:  time.Now(),
		}

		savedProduct, err := productService.Save(newProduct)
		assert.NoError(t, err)
		assert.Equal(t, newProduct.Name, savedProduct.Name)
	})

	t.Run("Save_ValidationError", func(t *testing.T) {
		invalidProduct := model.Products{
			Name:        "", // Invalid name
			Installment: 12,
			Bunga:       5,
			Created_At:  time.Now(),
		}

		_, err := productService.Save(invalidProduct)
		assert.Error(t, err)
		assert.EqualError(t, err, "name tidak boleh kosong")
	})

	t.Run("Update", func(t *testing.T) {
		updatedProduct := model.Products{
			ID:          1,
			Name:        "Updated Product",
			Installment: 24,
			Bunga:       8,
			Created_At:  time.Now(),
		}

		updated, err := productService.Update(updatedProduct)
		assert.NoError(t, err)
		assert.Equal(t, updatedProduct.Name, updated.Name)
	})

	t.Run("Update_ValidationError", func(t *testing.T) {
		invalidProduct := model.Products{
			ID:          1,
			Name:        "", // Invalid name
			Installment: 24,
			Bunga:       8,
			Created_At:  time.Now(),
		}

		_, err := productService.Update(invalidProduct)
		assert.Error(t, err)
		assert.EqualError(t, err, "name tidak boleh kosong")
	})
}
