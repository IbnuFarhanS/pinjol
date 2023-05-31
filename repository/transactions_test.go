package repository

// import (
// 	"testing"
// 	"time"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/IbnuFarhanS/pinjol/model"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"gorm.io/gorm/logger"
// )

// func setupTestDB_Transactions(t *testing.T) *gorm.DB {
// 	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjol_test port=5432 sslmode=disable"
// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	require.NoError(t, err)

// 	// Perform any necessary database setup

// 	return db
// }

// // ================== SAVE =========================
// func TestSaveTransactions(t *testing.T) {
// 	// Inisialisasi mock DB
// 	db, mock, _ := sqlmock.New()
// 	defer db.Close()

// 	// Inisialisasi GORM DB menggunakan mock DB
// 	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})

// 	// Inisialisasi repository dengan GORM DB
// 	repo := NewTransactionRepositoryImpl(gormDB)

// 	// Menyiapkan transactions status baru
// 	newTra := model.Transaction{
// 		ProductID: 1,
// 		UserID:    1,
// 		Status:    false,
// 	}
// 	// Menyiapkan query dan hasil yang diharapkan
// 	mock.ExpectBegin()
// 	mock.ExpectExec(`INSERT INTO "transactions" (.+) VALUES (.+)`).
// 		WithArgs(
// 			newTra.ProductID,
// 			newTra.UserID,
// 			newTra.Status,
// 			newTra.CreatedAt,
// 			newTra.ID).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	// Memanggil fungsi Save
// 	_, err := repo.Save(newTra)

// 	// Memastikan tidak ada error yang terjadi
// 	assert.NoError(t, err)
// }

// // ================== UPDATE =========================
// func TestUpdateTransactions(t *testing.T) {
// 	// Inisialisasi mock DB
// 	db, mock, _ := sqlmock.New()
// 	defer db.Close()

// 	// Inisialisasi GORM DB menggunakan mock DB
// 	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})

// 	// Inisialisasi repository dengan GORM DB
// 	repo := NewTransactionRepositoryImpl(gormDB)

// 	// Menyiapkan data transactions yang akan diupdate
// 	updateTra := model.Transaction{
// 		ID:        1,
// 		ProductID: 1,
// 		UserID:    1,
// 		Status:    false,
// 		CreatedAt: time.Now(), // Atur waktu yang sesuai
// 	}

// 	// Menyiapkan query dan hasil yang diharapkan
// 	mock.ExpectBegin()
// 	mock.ExpectExec(`UPDATE "transactions" SET (.+) WHERE "transactions"."id" = (.+)`).
// 		WithArgs(
// 			updateTra.ProductID,
// 			updateTra.UserID,
// 			updateTra.Status,
// 			updateTra.CreatedAt.UTC(),
// 			updateTra.ID).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	// Memanggil fungsi Update
// 	result, err := repo.Update(updateTra)

// 	// Memastikan tidak ada error yang terjadi
// 	assert.NoError(t, err)

// 	// Memastikan transactions yang diupdate sesuai dengan yang diharapkan
// 	assert.Equal(t, updateTra, result)
// }

// // ================== DELETE =========================
// func TestDeleteTransactions(t *testing.T) {
// 	// Inisialisasi mock DB
// 	db, mock, _ := sqlmock.New()
// 	defer db.Close()

// 	// Inisialisasi GORM DB menggunakan mock DB
// 	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
// 		Logger: logger.Default.LogMode(logger.Silent),
// 	})

// 	// Inisialisasi repository dengan GORM DB
// 	repo := NewTransactionRepositoryImpl(gormDB)

// 	// ID transactions yang akan dihapus
// 	traID := int64(1)

// 	// Menyiapkan query dan hasil yang diharapkan
// 	mock.ExpectBegin()
// 	mock.ExpectExec(`DELETE FROM "transactions" WHERE "transactions"."id" = (.+)`).
// 		WithArgs(traID).
// 		WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	// Memanggil fungsi Delete
// 	result, err := repo.Delete(uint(traID))

// 	// Memastikan tidak ada error yang terjadi
// 	assert.NoError(t, err)

// 	// Memastikan transactions yang dihapus sesuai dengan yang diharapkan
// 	expectedTra := model.Transaction{} // Atur sesuai dengan nilai yang diharapkan
// 	assert.Equal(t, expectedTra, result)
// }

// // ================== FIND BY ID =========================
// func TestFindByIdTransactions(t *testing.T) {
// 	db := setupTestDB_Transactions(t)
// 	repo := NewTransactionRepositoryImpl(db)

// 	// Test FindById for ID 1
// 	foundTransactions, err := repo.FindById(1)
// 	require.Nil(t, err)

// 	// Expected Transactions with ID 1
// 	expectedTransactions := model.Transaction{
// 		ID:        1,
// 		UserID:    1,
// 		ProductID: 1,
// 		Status:    false,
// 		CreatedAt: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
// 		DueDate:   time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
// 	}

// 	require.Equal(t, expectedTransactions, foundTransactions)
// }

// // ================== FIND ALL =========================
// func TestFindAllTransactions(t *testing.T) {
// 	db := setupTestDB_Transactions(t)
// 	repo := NewTransactionRepositoryImpl(db)

// 	// Create multiple Transactions in the database
// 	dummyTransactions := []model.Transaction{
// 		{
// 			ID:        1,
// 			UserID:    1,
// 			ProductID: 1,
// 			Status:    false,
// 			CreatedAt: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
// 			DueDate:   time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
// 		},
// 		{
// 			ID:        2,
// 			UserID:    1,
// 			ProductID: 1,
// 			Status:    false,
// 			CreatedAt: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
// 			DueDate:   time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
// 		},
// 		// Add more Transactions if needed
// 	}

// 	// Test FindAll
// 	foundTransactions, err := repo.FindAll()
// 	require.NoError(t, err)
// 	require.Equal(t, len(dummyTransactions), len(foundTransactions))
// }
