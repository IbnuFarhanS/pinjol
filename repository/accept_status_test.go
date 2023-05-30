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

func setupTestDB_AcceptStatus(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjol port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Perform any necessary database setup

	return db
}

func TestFindAllAccStat(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewAcceptStatusRepositoryImpl(gormDB)

	// Test case: Berhasil menemukan semua metode pembayaran
	mock.ExpectQuery(`SELECT \* FROM "accept_statuses"`).
		WillReturnRows(sqlmock.NewRows([]string{"id", "id_transaction", "status"}).
			AddRow(1, 1, false).
			AddRow(2, 1, true))

	accStats, err := repo.FindAll()
	assert.NoError(t, err)
	assert.Len(t, accStats, 2)
}

func TestFindByIdAsccStat(t *testing.T) {
	db := setupTestDB_AcceptStatus(t)
	repo := NewAcceptStatusRepositoryImpl(db)

	// Test FindById for ID 1
	foundPayMeth, err := repo.FindById(1)
	require.NoError(t, err)

	// Expected Paymeth with ID 1
	expectedPayMeth := model.AcceptStatus{
		ID:            1,
		TransactionID: 1,
		Status:        true,
		CreatedAt:     time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedPayMeth, foundPayMeth)
}

func TestSaveAcceptStatus(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewAcceptStatusRepositoryImpl(gormDB)

	// Menyiapkan data status penerimaan baru
	newAcceptStatus := model.AcceptStatus{
		TransactionID: 1,
		Status:        false,
		CreatedAt:     time.Now(),
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "accept_statuses" (.+) VALUES (.+) RETURNING "id","created_at"`).
		WithArgs(
			newAcceptStatus.TransactionID,
			newAcceptStatus.Status,
			newAcceptStatus.CreatedAt).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).
			AddRow(1, newAcceptStatus.CreatedAt))
	mock.ExpectCommit()

	// Memanggil fungsi Save
	savedAcceptStatus, err := repo.Save(newAcceptStatus)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan data status penerimaan yang disimpan sesuai dengan yang diharapkan
	assert.Equal(t, newAcceptStatus.TransactionID, savedAcceptStatus.TransactionID)
	assert.Equal(t, newAcceptStatus.Status, savedAcceptStatus.Status)
}

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

	// Test case: Berhasil mengupdate metode pembayaran
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "accept_statuses" SET (.+) WHERE "accept_statuses"."id" = \?`).
		WithArgs(false, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	updatedAcceptStatus := model.AcceptStatus{
		ID:            1,
		TransactionID: 1,
		Status:        false,
	}

	updatedStatus, err := repo.Update(updatedAcceptStatus)

	assert.NoError(t, err)
	assert.Equal(t, updatedAcceptStatus, updatedStatus)
}

func TestDeleteAccStat(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewAcceptStatusRepositoryImpl(gormDB)

	// ID AccStat yang akan dihapus
	accstatID := uint(1)

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "accept_statuses" WHERE "accept_statuses"."id" = (.+)`).
		WithArgs(accstatID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Delete
	result, err := repo.Delete(uint(accstatID))

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan AccStat yang dihapus sesuai dengan yang diharapkan
	expectedAccStat := model.AcceptStatus{} // Atur sesuai dengan nilai yang diharapkan
	assert.Equal(t, expectedAccStat, result)
}
