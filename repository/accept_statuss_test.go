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

func setupTestDB_AccStats(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjol port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Perform any necessary database setup

	return db
}

// ================== SAVE =========================
func TestSaveAccStats(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewAcceptStatusRepositoryImpl(gormDB)

	// Menyiapkan data accept status baru
	newAccStat := model.AcceptStatus{
		TransactionsID: 1,
		Status:         true,
	}
	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "accept_statuses" (.+) VALUES (.+)`).
		WithArgs(
			newAccStat.TransactionsID,
			newAccStat.Status,
			newAccStat.Created_At,
			newAccStat.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Save
	_, err := repo.Save(newAccStat)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)
}

// ================== UPDATE =========================
func TestUpdateAccStats(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewAcceptStatusRepositoryImpl(gormDB)

	// Menyiapkan data accept status yang akan diupdate
	updateAccStat := model.AcceptStatus{
		ID:             1,
		TransactionsID: 1,
		Status:         true,
		Created_At:     time.Now(), // Atur waktu yang sesuai
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "accept_statuses" SET (.+) WHERE "accept_statuses"."id" = (.+)`).
		WithArgs(
			updateAccStat.TransactionsID,
			updateAccStat.Status,
			updateAccStat.Created_At.UTC(),
			updateAccStat.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Update
	result, err := repo.Update(updateAccStat)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan accept status yang diupdate sesuai dengan yang diharapkan
	assert.Equal(t, updateAccStat, result)
}

// ================== DELETE =========================
func TestDeleteAccStats(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewAcceptStatusRepositoryImpl(gormDB)

	// ID accept status yang akan dihapus
	accstatID := int64(1)

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "accept_statuses" WHERE "accept_statuses"."id" = (.+)`).
		WithArgs(accstatID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Delete
	result, err := repo.Delete(accstatID)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan accept status yang dihapus sesuai dengan yang diharapkan
	expectedAccStat := model.AcceptStatus{} // Atur sesuai dengan nilai yang diharapkan
	assert.Equal(t, expectedAccStat, result)
}

// ================== FIND BY ID =========================
func TestFindByIdAccStats(t *testing.T) {
	db := setupTestDB_AccStats(t)
	repo := NewAcceptStatusRepositoryImpl(db)

	// Test FindById for ID 1
	foundAccStats, err := repo.FindById(1)
	require.Nil(t, err)

	// Expected AccStats with ID 1
	expectedAccStats := model.AcceptStatus{
		ID:             1,
		TransactionsID: 1,
		Status:         true,
		Created_At:     time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedAccStats, foundAccStats)
}

// ================== FIND ALL =========================
func TestFindAllAccStats(t *testing.T) {
	db := setupTestDB_AccStats(t)
	repo := NewAcceptStatusRepositoryImpl(db)

	// Create multiple AccStats in the database
	dummyAccStats := []model.AcceptStatus{
		{
			ID:             1,
			TransactionsID: 1,
			Status:         true,
			Created_At:     time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
		},
		{
			ID:             2,
			TransactionsID: 1,
			Status:         false,
			Created_At:     time.Date(2023, 5, 27, 0, 0, 0, 0, time.Local),
		},
		// Add more AccStats if needed
	}

	// Test FindAll
	foundAccStats, err := repo.FindAll()
	require.NoError(t, err)
	require.Equal(t, len(dummyAccStats), len(foundAccStats))
}
