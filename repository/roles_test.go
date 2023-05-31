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

func setupTestDB_Roles(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjol port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Perform any necessary database setup

	return db
}

// ================== SAVE =========================
func TestSaveRoles(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewRoleRepositoryImpl(gormDB)

	// Menyiapkan data roles baru
	newRoles := model.Role{
		Name: "Roles 1",
	}
	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "roles" (.+) VALUES (.+)`).
		WithArgs(
			newRoles.Name,
			newRoles.CreatedAt,
			newRoles.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Save
	_, err := repo.Save(newRoles)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)
}

// ================== UPDATE =========================
func TestUpdateRoles(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewRoleRepositoryImpl(gormDB)

	// Menyiapkan data roles yang akan diupdate
	updateRoles := model.Role{
		ID:        1,
		Name:      "Ibnu",
		CreatedAt: time.Now(), // Atur waktu yang sesuai
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "roles" SET (.+) WHERE "roles"."id" = (.+)`).
		WithArgs(
			updateRoles.Name,
			updateRoles.CreatedAt.UTC(),
			updateRoles.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Update
	result, err := repo.Update(updateRoles)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan roles yang diupdate sesuai dengan yang diharapkan
	assert.Equal(t, updateRoles, result)
}

// ================== DELETE =========================
func TestDeleteRoles(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewRoleRepositoryImpl(gormDB)

	// ID roles yang akan dihapus
	roleID := int64(1)

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "roles" WHERE "roles"."id" = (.+)`).
		WithArgs(roleID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Delete
	result, err := repo.Delete(uint(roleID))

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan roles yang dihapus sesuai dengan yang diharapkan
	expectedUser := model.Role{} // Atur sesuai dengan nilai yang diharapkan
	assert.Equal(t, expectedUser, result)
}

// ================== FIND BY ID =========================
func TestFindByIdRoles(t *testing.T) {
	db := setupTestDB_Roles(t)
	repo := NewRoleRepositoryImpl(db)

	// Test FindById for ID 1
	foundRoles, err := repo.FindById(1)
	require.NotNil(t, err)

	// Expected Roles with ID 1
	expectedRoles := model.Role{
		ID:        1,
		Name:      "users",
		CreatedAt: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedRoles, foundRoles)
}

// ================== FIND BY NAME =========================
func TestFindByNameRoles(t *testing.T) {
	db := setupTestDB_Roles(t)
	repo := NewRoleRepositoryImpl(db)

	// Test FindByName for Name
	foundRoles, err := repo.FindByName("users")
	require.NoError(t, err)

	// Expected Roles with Name
	expectedRoles := model.Role{
		ID:        1,
		Name:      "users",
		CreatedAt: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedRoles, foundRoles)

	// Test FindByName for non-existing name "Nonexistent"
	_, err = repo.FindByName("Nonexistent")
	require.Error(t, err)
}

// ================== FIND ALL =========================
func TestFindAllRoles(t *testing.T) {
	db := setupTestDB_Roles(t)
	repo := NewRoleRepositoryImpl(db)

	// Create multiple roles in the database
	dummyRoles := []model.Role{
		{
			ID:        1,
			Name:      "ibnu",
			CreatedAt: time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
		},
		{
			ID:        2,
			Name:      "admin",
			CreatedAt: time.Date(2023, 5, 27, 17, 15, 20, 540820000, time.FixedZone("WIB", 7*60*60)),
		},
		// Add more roles if needed
	}

	// Test FindAll
	foundRoles, err := repo.FindAll()
	require.NoError(t, err)
	require.Equal(t, len(dummyRoles), len(foundRoles))
}
