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

func setupTestDB_PaymentMethods(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjol port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Perform any necessary database setup

	return db
}

// ================== SAVE =========================
func TestSavePaymentMethods(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewPaymentMethodRepositoryImpl(gormDB)

	// Menyiapkan data produk baru
	newPaymentMethod := model.PaymentMethod{
		Name: "Payment Method",
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "payment_methods" (.+) VALUES (.+)`).
		WithArgs(newPaymentMethod.Name, newPaymentMethod.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Save
	_, err := repo.Save(newPaymentMethod)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)
}

// ================== FIND BY ID =========================
func TestFindByIdPaymentMethods(t *testing.T) {
	db := setupTestDB_PaymentMethods(t)
	repo := NewPaymentMethodRepositoryImpl(db)

	// Test FindById for ID 1
	foundPaymentMethods, err := repo.FindById(1)
	require.NoError(t, err)

	// Expected PaymentMethods with ID 1
	expectedPaymentMethods := model.PaymentMethod{
		ID:         1,
		Name:       "Transfer Bank",
		Created_At: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedPaymentMethods, foundPaymentMethods)
}

// ================== FIND BY NAME =========================
func TestFindByNamePaymentMethods(t *testing.T) {
	db := setupTestDB_PaymentMethods(t)
	repo := NewPaymentMethodRepositoryImpl(db)

	// Test FindByUsername for Name
	foundPaymentMethods, err := repo.FindByName("Transfer Bank")
	require.NoError(t, err)

	// Expected PaymentMethods with Name
	expectedPaymentMethods := model.PaymentMethod{
		ID:         1,
		Name:       "Transfer Bank",
		Created_At: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedPaymentMethods, foundPaymentMethods)

	// Test FindByUsername for non-existing username "Nonexistent"
	_, err = repo.FindByName("Nonexistent")
	require.Error(t, err)
	require.EqualError(t, err, "invalid name")
}

// ================== FIND ALL =========================
func TestFindAllPaymentMethods(t *testing.T) {
	db := setupTestDB_PaymentMethods(t)
	repo := NewPaymentMethodRepositoryImpl(db)

	// Create multiple PaymentMethods in the database
	PaymentMethods := []model.PaymentMethod{
		{
			ID:         1,
			Name:       "Transfer Bank",
			Created_At: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
		},
		{
			ID:         2,
			Name:       "Virtual Account",
			Created_At: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
		},
		// Add more PaymentMethods if needed
	}

	// Test FindAll
	foundPaymentMethods, err := repo.FindAll()
	require.NoError(t, err)
	require.Equal(t, len(PaymentMethods), len(foundPaymentMethods))

	// Compare each PaymentMethods in the expected list with the found PaymentMethodss
	for _, expectedPaymentMethods := range PaymentMethods {
		found := false
		for _, actualPaymentMethods := range foundPaymentMethods {
			if expectedPaymentMethods.ID == actualPaymentMethods.ID {
				require.Equal(t, expectedPaymentMethods, actualPaymentMethods)
				found = true
				break
			}
		}
		require.True(t, found, "PaymentMethods not found: ID %d", expectedPaymentMethods.ID)
	}
}

// ================== UPDATE =========================
func TestUpdatePaymentMethods(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewPaymentMethodRepositoryImpl(gormDB)

	// Menyiapkan data payment methods yang akan diupdate
	updatePaymentMethods := model.PaymentMethod{
		ID:         1,
		Name:       "Updated PaymentMethods",
		Created_At: time.Now(), // Atur waktu yang sesuai
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "payment_methods" SET (.+) WHERE "payment_methods"."id" = (.+)`).
		WithArgs(updatePaymentMethods.Name, updatePaymentMethods.Created_At.UTC(), updatePaymentMethods.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Update
	result, err := repo.Update(updatePaymentMethods)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan produk yang diupdate sesuai dengan yang diharapkan
	assert.Equal(t, updatePaymentMethods, result)
}

// ================== DELETE =========================
func TestDeletePaymentMethods(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewPaymentMethodRepositoryImpl(gormDB)

	// ID payment methods yang akan dihapus
	payment_methodsID := int64(1)

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "payment_methods" WHERE "payment_methods"."id" = (.+)`).
		WithArgs(payment_methodsID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Delete
	result, err := repo.Delete(payment_methodsID)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan payment methods yang dihapus sesuai dengan yang diharapkan
	expectedPaymentMethods := model.PaymentMethod{} // Atur sesuai dengan nilai yang diharapkan
	assert.Equal(t, expectedPaymentMethods, result)
}
