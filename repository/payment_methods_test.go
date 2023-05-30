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

func TestFindAllPaymentMethods(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewPaymentMethodRepositoryImpl(gormDB)

	// Test case: Berhasil menemukan semua metode pembayaran
	mock.ExpectQuery(`SELECT \* FROM "payment_methods"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "Method 1").
			AddRow(2, "Method 2"))

	paymentMethods, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, paymentMethods, 2)
}

func TestFindByIdPaymentMethods(t *testing.T) {
	db := setupTestDB_PaymentMethods(t)
	repo := NewPaymentMethodRepositoryImpl(db)

	// Test FindById for ID 1
	foundPayMeth, err := repo.FindById(1)
	require.NoError(t, err)

	// Expected Paymeth with ID 1
	expectedPayMeth := model.PaymentMethod{
		ID:        1,
		Name:      "Transfer Bank",
		CreatedAt: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedPayMeth, foundPayMeth)
}

func TestFindByNamePaymentMethods(t *testing.T) {
	db := setupTestDB_PaymentMethods(t)
	repo := NewPaymentMethodRepositoryImpl(db)

	// Test FindByName for Name
	foundPayMeth, err := repo.FindByName("Transfer Bank")
	require.NoError(t, err)

	// Expected paymeth with Name
	expectedPayMeth := model.PaymentMethod{
		ID:        1,
		Name:      "Transfer Bank",
		CreatedAt: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedPayMeth, foundPayMeth)

	// Test FindByName for non-existing name "Nonexistent"
	_, err = repo.FindByName("Nonexistent")
	require.Error(t, err)
	require.EqualError(t, err, "invalid name")
}

func TestSavePaymentMethod(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewPaymentMethodRepositoryImpl(gormDB)

	// Test case: Berhasil menyimpan metode pembayaran baru
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "payment_methods" (.+) VALUES (.+)`).
		WithArgs("Method 1").
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	newPaymentMethod := model.PaymentMethod{
		Name: "Method 1",
	}

	paymentMethod, err := repo.Save(newPaymentMethod)
	assert.NoError(t, err)
	assert.Equal(t, "Method 1", paymentMethod.Name)
}

func TestUpdatePaymentMethod(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewPaymentMethodRepositoryImpl(gormDB)

	// Test case: Berhasil mengupdate metode pembayaran
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "payment_methods" SET (.+) WHERE "payment_methods"."id" = \?`).
		WithArgs("Updated Method 1", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	updatedPaymentMethod := model.PaymentMethod{
		ID:   1,
		Name: "Updated Method 1",
	}

	paymentMethod, err := repo.Update(updatedPaymentMethod)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Method 1", paymentMethod.Name)
}

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

	// ID paymeth yang akan dihapus
	paymethID := uint(1)

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "payment_methods" WHERE "payment_methods"."id" = (.+)`).
		WithArgs(paymethID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Delete
	result, err := repo.Delete(uint(paymethID))

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan paymeth yang dihapus sesuai dengan yang diharapkan
	expectedPaymeth := model.PaymentMethod{} // Atur sesuai dengan nilai yang diharapkan
	assert.Equal(t, expectedPaymeth, result)
}
