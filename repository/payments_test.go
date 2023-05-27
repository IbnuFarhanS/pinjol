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

func setupTestDB_Payments(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjol port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Perform any necessary database setup

	return db
}

// ================== SAVE =========================
func TestSaveAccPayments(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewPaymentsRepositoryImpl(gormDB)

	// Menyiapkan data payments baru
	newPayment := model.Payments{
		TransactionsID:  1,
		PaymentMethodID: 1,
		Payment_Amount:  1000000,
	}
	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "payments" (.+) VALUES (.+)`).
		WithArgs(
			newPayment.TransactionsID,
			newPayment.PaymentMethodID,
			newPayment.Payment_Amount,
			newPayment.Payment_Date,
			newPayment.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Save
	_, err := repo.Save(newPayment)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)
}

// ================== UPDATE =========================
func TestUpdatePayments(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewPaymentsRepositoryImpl(gormDB)

	// Menyiapkan data payments yang akan diupdate
	updatePayment := model.Payments{
		ID:              1,
		TransactionsID:  1,
		PaymentMethodID: 1,
		Payment_Amount:  1000000,
		Payment_Date:    time.Now(), // Atur waktu yang sesuai
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "payments" SET (.+) WHERE "payments"."id" = (.+)`).
		WithArgs(
			updatePayment.TransactionsID,
			updatePayment.PaymentMethodID,
			updatePayment.Payment_Amount,
			updatePayment.Payment_Date.UTC(),
			updatePayment.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Update
	result, err := repo.Update(updatePayment)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan payments yang diupdate sesuai dengan yang diharapkan
	assert.Equal(t, updatePayment, result)
}

// ================== DELETE =========================
func TestDeletePayments(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewPaymentsRepositoryImpl(gormDB)

	// ID payments yang akan dihapus
	paymentID := int64(1)

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "payments" WHERE "payments"."id" = (.+)`).
		WithArgs(paymentID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Delete
	result, err := repo.Delete(paymentID)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan payments yang dihapus sesuai dengan yang diharapkan
	expectedPayment := model.Payments{} // Atur sesuai dengan nilai yang diharapkan
	assert.Equal(t, expectedPayment, result)
}

// ================== FIND BY ID =========================
func TestFindByIdPayments(t *testing.T) {
	db := setupTestDB_Payments(t)
	repo := NewPaymentsRepositoryImpl(db)

	// Test FindById for ID 1
	foundPayments, err := repo.FindById(1)
	require.Nil(t, err)

	// Expected Payments with ID 1
	expectedPayments := model.Payments{
		ID:              1,
		TransactionsID:  1,
		PaymentMethodID: 1,
		Payment_Amount:  1000000,
		Payment_Date:    time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedPayments, foundPayments)
}

// ================== FIND ALL =========================
func TestFindAllPayments(t *testing.T) {
	db := setupTestDB_Payments(t)
	repo := NewPaymentsRepositoryImpl(db)

	// Create multiple Payments in the database
	dummyPayment := []model.Payments{
		{
			ID:              1,
			TransactionsID:  1,
			PaymentMethodID: 1,
			Payment_Amount:  1000000,
			Payment_Date:    time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
		},
		{
			ID:              2,
			TransactionsID:  1,
			PaymentMethodID: 1,
			Payment_Amount:  1000000,
			Payment_Date:    time.Date(2023, 5, 27, 0, 0, 0, 0, time.Local),
		},
		// Add more Payments if needed
	}

	// Test FindAll
	foundPayments, err := repo.FindAll()
	require.NoError(t, err)
	require.Equal(t, len(dummyPayment), len(foundPayments))
}