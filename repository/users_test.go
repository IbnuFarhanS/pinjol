package repository

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupTestDB_Users(t *testing.T) *gorm.DB {
	dsn := "host=localhost user=postgres password=sql1234 dbname=pinjol port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err)

	// Perform any necessary database setup

	return db
}

// ================== SAVE =========================
func TestSaveUsers(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewUserRepositoryImpl(gormDB)

	// Menyiapkan data user baru
	newUsers := model.User{
		Username:    "User",
		Password:    "User123",
		NIK:         "1234",
		Name:        "Ibnu",
		Address:     "Bandung",
		PhoneNumber: "084578458",
		Limit:       2000000,
		RoleID:      1,
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectQuery(`SELECT \* FROM "users" WHERE "username" = \? OR "nik" = \? OR "phone_number" = \?`).
		WithArgs(newUsers.Username, newUsers.NIK, newUsers.PhoneNumber).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users" (.+) VALUES (.+)`).
		WithArgs(
			newUsers.Username,
			newUsers.Password,
			newUsers.NIK,
			newUsers.Name,
			newUsers.Address,
			newUsers.PhoneNumber,
			newUsers.Limit,
			newUsers.RoleID,
			newUsers.CreatedAt,
			newUsers.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Save
	_, err := repo.Save(newUsers)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)
}

// ================== UPDATE =========================
func TestUpdateUsers(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewUserRepositoryImpl(gormDB)

	// Menyiapkan data user yang akan diupdate
	updateUser := model.User{
		ID:          1,
		Username:    "User",
		Password:    "User123",
		NIK:         "1234",
		Name:        "Ibnu",
		Address:     "Bandung",
		PhoneNumber: "084578458",
		Limit:       2000000,
		RoleID:      0,
		CreatedAt:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET (.+) WHERE "users"."id" = (.+)`).
		WithArgs(
			updateUser.Username,
			updateUser.Password,
			updateUser.NIK,
			updateUser.Name,
			updateUser.Address,
			updateUser.PhoneNumber,
			updateUser.Limit,
			updateUser.RoleID,
			updateUser.CreatedAt,
			updateUser.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Update
	result, err := repo.Update(updateUser)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan user yang diupdate sesuai dengan yang diharapkan
	assert.Equal(t, updateUser, result)
}

// ================== DELETE =========================
func TestDeleteUsers(t *testing.T) {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewUserRepositoryImpl(gormDB)

	// ID user yang akan dihapus
	userID := int64(1)

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "users" WHERE "users"."id" = (.+)`).
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Delete
	result, err := repo.Delete(uint(userID))

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan user yang dihapus sesuai dengan yang diharapkan
	expectedUser := model.User{} // Atur sesuai dengan nilai yang diharapkan
	assert.Equal(t, expectedUser, result)
}

// ================== FIND BY ID =========================
type TestUsers struct {
	suite.Suite
}

func (suite *TestUsers) TestFindByIdUsers() {
	// Inisialisasi mock DB
	db, mock, _ := sqlmock.New()
	defer db.Close()

	// Inisialisasi GORM DB menggunakan mock DB
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	// Inisialisasi repository dengan GORM DB
	repo := NewUserRepositoryImpl(gormDB)

	// Setup dummy data
	createDummyUser := []model.User{
		{
			ID:          1,
			Username:    "User",
			Password:    "User123",
			NIK:         "1234",
			Name:        "Ibnu",
			Address:     "Bandung",
			PhoneNumber: "084578458",
			Limit:       2000000,
			RoleID:      1,
			CreatedAt:   time.Now(),
		},
	}

	// Expect query
	rows := sqlmock.NewRows([]string{"id", "username", "password", "nik", "name", "alamat", "phone_number", "limit", "id_role", "created_at"})
	for _, d := range createDummyUser {
		rows.AddRow(d.ID, d.Username, d.Password, d.NIK, d.Name, d.Address, d.PhoneNumber, d.Limit, d.RoleID, d.CreatedAt)
	}
	mock.ExpectQuery("SELECT id, username, password, nik, name, alamat, phone_number, limit, id_role, created_at FROM users").WillReturnRows(rows)

	// Panggil metode yang diuji
	users, err := repo.FindById(1)
	if err != nil {
		return
	}

	suite.Equal(createDummyUser[0].ID, users)

}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(TestUsers))
}

// ================== FIND BY USERNAME =========================
func TestFindByUsernameUsers(t *testing.T) {
	db := setupTestDB_Users(t)
	repo := NewUserRepositoryImpl(db)

	// Test FindByUsername for Name
	foundUsers, err := repo.FindByUsername("ibnu")
	require.NoError(t, err)

	// Expected Users with Name
	expectedUsers := model.User{
		ID:          1,
		Username:    "ibnu",
		Password:    "ibnu",
		NIK:         "1234",
		Name:        "ibnu",
		Address:     "",
		PhoneNumber: "084579856598",
		Limit:       2000000,
		RoleID:      1,
		CreatedAt:   time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedUsers, foundUsers)

	// Test FindByUsername for non-existing username "Nonexistent"
	_, err = repo.FindByUsername("Nonexistent")
	require.Error(t, err)
}

// ================== FIND ALL =========================
func TestFindAllUsers(t *testing.T) {
	db := setupTestDB_Users(t)
	repo := NewUserRepositoryImpl(db)

	// Create multiple Users in the database
	dummyUser := []model.User{
		{
			ID:          1,
			Username:    "ibnu",
			Password:    "ibnu",
			NIK:         "1234",
			Name:        "ibnu",
			Address:     "bandung",
			PhoneNumber: "084579856598",
			Limit:       2000000,
			RoleID:      1,
			CreatedAt:   time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
		},
		{
			ID:          2,
			Username:    "ab",
			Password:    "$2a$10$/GrmpqvDB/exiMDZQ1KtTuOAFeWbK/vZXyCCKcHqZE5o2Ild8fPUm",
			NIK:         "ab",
			Name:        "ab",
			Address:     "ab",
			PhoneNumber: "ab",
			Limit:       2000000,
			RoleID:      1,
			CreatedAt:   time.Date(2023, 5, 27, 17, 15, 20, 540820000, time.FixedZone("WIB", 7*60*60)),
		},
		// Add more Users if needed
	}

	// Test FindAll
	foundUsers, err := repo.FindAll()
	require.NoError(t, err)
	require.Equal(t, len(dummyUser), len(foundUsers))
}
