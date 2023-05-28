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
	repo := NewUsersRepositoryImpl(gormDB)

	// Menyiapkan data user baru
	newUsers := model.Users{
		Username:     "User",
		Password:     "User123",
		Nik:          "1234",
		Name:         "Ibnu",
		Alamat:       "Bandung",
		Phone_Number: "084578458",
		Limit:        2000000,
		RolesID:      1,
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`INSERT INTO "users" (.+) VALUES (.+)`).
		WithArgs(
			newUsers.Username,
			newUsers.Password,
			newUsers.Name,
			newUsers.Alamat,
			newUsers.Phone_Number,
			newUsers.Limit,
			newUsers.RolesID,
			newUsers.Created_At,
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
	repo := NewUsersRepositoryImpl(gormDB)

	// Menyiapkan data user yang akan diupdate
	updateUser := model.Users{
		ID:           1,
		Username:     "User",
		Password:     "User123",
		Nik:          "1234",
		Name:         "Ibnu",
		Alamat:       "Bandung",
		Phone_Number: "084578458",
		Limit:        2000000,
		RolesID:      0,
		Created_At:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE "users" SET (.+) WHERE "users"."id" = (.+)`).
		WithArgs(
			updateUser.Username,
			updateUser.Password,
			updateUser.Name,
			updateUser.Alamat,
			updateUser.Phone_Number,
			updateUser.Limit,
			updateUser.RolesID,
			updateUser.Created_At,
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
	repo := NewUsersRepositoryImpl(gormDB)

	// ID user yang akan dihapus
	userID := int64(1)

	// Menyiapkan query dan hasil yang diharapkan
	mock.ExpectBegin()
	mock.ExpectExec(`DELETE FROM "users" WHERE "users"."id" = (.+)`).
		WithArgs(userID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Memanggil fungsi Delete
	result, err := repo.Delete(userID)

	// Memastikan tidak ada error yang terjadi
	assert.NoError(t, err)

	// Memastikan user yang dihapus sesuai dengan yang diharapkan
	expectedUser := model.Users{} // Atur sesuai dengan nilai yang diharapkan
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
	repo := NewUsersRepositoryImpl(gormDB)

	// Setup dummy data
	createDummyUser := []model.Users{
		{
			ID:           1,
			Username:     "User",
			Password:     "User123",
			Nik:          "1234",
			Name:         "Ibnu",
			Alamat:       "Bandung",
			Phone_Number: "084578458",
			Limit:        2000000,
			RolesID:      1,
			Created_At:   time.Now(),
		},
	}

	// Expect query
	rows := sqlmock.NewRows([]string{"id", "username", "password", "nik", "name", "alamat", "phone_number", "limit", "id_role", "created_at"})
	for _, d := range createDummyUser {
		rows.AddRow(d.ID, d.Username, d.Password, d.Nik, d.Name, d.Alamat, d.Phone_Number, d.Limit, d.RolesID, d.Created_At)
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
	repo := NewUsersRepositoryImpl(db)

	// Test FindByUsername for Name
	foundUsers, err := repo.FindByUsername("ibnu")
	require.NoError(t, err)

	// Expected Users with Name
	expectedUsers := model.Users{
		ID:           1,
		Username:     "ibnu",
		Password:     "ibnu",
		Nik:          "1234",
		Name:         "ibnu",
		Alamat:       "bandung",
		Phone_Number: "084579856598",
		Limit:        2000000,
		RolesID:      1,
		Created_At:   time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
	}

	require.Equal(t, expectedUsers, foundUsers)

	// Test FindByUsername for non-existing username "Nonexistent"
	_, err = repo.FindByUsername("Nonexistent")
	require.Error(t, err)
}

// ================== FIND ALL =========================
func TestFindAllUsers(t *testing.T) {
	db := setupTestDB_Users(t)
	repo := NewUsersRepositoryImpl(db)

	// Create multiple Users in the database
	dummyUser := []model.Users{
		{
			ID:           1,
			Username:     "ibnu",
			Password:     "ibnu",
			Nik:          "1234",
			Name:         "ibnu",
			Alamat:       "bandung",
			Phone_Number: "084579856598",
			Limit:        2000000,
			RolesID:      1,
			Created_At:   time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
		},
		{
			ID:           2,
			Username:     "ab",
			Password:     "$2a$10$/GrmpqvDB/exiMDZQ1KtTuOAFeWbK/vZXyCCKcHqZE5o2Ild8fPUm",
			Nik:          "ab",
			Name:         "ab",
			Alamat:       "ab",
			Phone_Number: "ab",
			Limit:        2000000,
			RolesID:      1,
			Created_At:   time.Date(2023, 5, 27, 17, 15, 20, 540820000, time.FixedZone("WIB", 7*60*60)),
		},
		// Add more Users if needed
	}

	// Test FindAll
	foundUsers, err := repo.FindAll()
	require.NoError(t, err)
	require.Equal(t, len(dummyUser), len(foundUsers))
}
