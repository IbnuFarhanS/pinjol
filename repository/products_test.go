package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB_Products(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjol port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Perform any necessary database setup

	return db
}

// ================== SAVE =========================
func TestSaveProducts(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewProductsRepositoryImpl(gormDB)

	// Menyiapkan data produk baru
	newProduct := model.Products{
		Name:        "Product",
		Installment: 6,
		Bunga:       0.2,
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "products" (.+) VALUES (.+)`).
		WithArgs(newProduct.Name, newProduct.Installment, newProduct.Bunga, newProduct.Created_At, newProduct.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Save
	_, err := repo.Save(newProduct)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)
}

// ================== FIND BY ID =========================
func TestFindByIdProducts(t *testing.T) {
	db := setupTestDB_Products(t)
	repo := NewProductsRepositoryImpl(db)

	// Test FindById for ID 1
	foundProducts, err := repo.FindById(1)
	require.NoError(t, err)

	// Expected Products with ID 1
	expectedProducts := model.Products{
		ID:          1,
		Name:        "cicilan 6 bulan",
		Installment: 6,
		Bunga:       0.2,
		Created_At:  time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedProducts, foundProducts)
}

// ================== FIND BY NAME =========================
func TestFindByNameProducts(t *testing.T) {
	db := setupTestDB_Products(t)
	repo := NewProductsRepositoryImpl(db)

	// Test FindByName for Name
	foundProducts, err := repo.FindByName("cicilan 6 bulan")
	require.NoError(t, err)

	// Expected Products with Name
	expectedProducts := model.Products{
		ID:          1,
		Name:        "cicilan 6 bulan",
		Installment: 6,
		Bunga:       0.2,
		Created_At:  time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedProducts, foundProducts)

	// Test FindByUsername for non-existing username "Nonexistent"
	_, err = repo.FindByName("Nonexistent")
	require.Error(t, err)
	require.EqualError(t, err, "invalid name")
}

// ================== FIND ALL =========================
func TestFindAll(t *testing.T) {
	db := setupTestDB_Products(t)
	repo := NewProductsRepositoryImpl(db)

	// Create multiple Productss in the database
	Productss := []model.Products{
		{
			ID:          1,
			Name:        "cicilan 6 bulan",
			Installment: 6,
			Bunga:       0.2,
			Created_At:  time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
		},
		{
			ID:          2,
			Name:        "cicilan 12 bulan",
			Installment: 12,
			Bunga:       0.4,
			Created_At:  time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
		},
		// Add more Productss if needed
	}

	// Test FindAll
	foundProductss, err := repo.FindAll()
	require.NoError(t, err)
	require.Equal(t, len(Productss), len(foundProductss))

	// Compare each Products in the expected list with the found Productss
	for _, expectedProducts := range Productss {
		found := false
		for _, actualProducts := range foundProductss {
			if expectedProducts.ID == actualProducts.ID {
				require.Equal(t, expectedProducts, actualProducts)
				found = true
				break
			}
		}
		require.True(t, found, "Products not found: ID %d", expectedProducts.ID)
	}
}

// ================== UPDATE =========================
func TestUpdateProducts(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewProductsRepositoryImpl(gormDB)

	// Menyiapkan data produk yang akan diupdate
	updatedProduct := model.Products{
		ID:          1,
		Name:        "Updated Product",
		Installment: 12,
		Bunga:       0.3,
		Created_At:  time.Now(), // Atur waktu yang sesuai
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "products" SET (.+) WHERE "products"."id" = (.+)`).
		WithArgs(updatedProduct.Name, updatedProduct.Installment, updatedProduct.Bunga, updatedProduct.Created_At.UTC(), updatedProduct.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Update
	result, err := repo.Update(updatedProduct)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan produk yang diupdate sesuai dengan yang diharapkan
	assert.Equal(t, updatedProduct, result)
}

// ================== DELETE =========================
func TestDeleteProducts(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewProductsRepositoryImpl(gormDB)

	// ID produk yang akan dihapus
	productID := int64(1)

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "products" WHERE "products"."id" = (.+)`).
		WithArgs(productID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Delete
	result, err := repo.Delete(productID)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan produk yang dihapus sesuai dengan yang diharapkan
	expectedProduct := model.Products{} // Atur sesuai dengan nilai yang diharapkan
	assert.Equal(t, expectedProduct, result)
}
