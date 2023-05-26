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
		RolesID:      1,
		Created_At:   time.Now(), // Atur waktu yang sesuai
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
			updateUser.Created_At.UTC(),
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
func TestFindByIdUsers(t *testing.T) {
	db := setupTestDB_Users(t)
	repo := NewUsersRepositoryImpl(db)

	// Test FindById for ID 1
	foundUsers, err := repo.FindById(1)
	require.NoError(t, err)

	// Expected Users with ID 1
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
}

// // ================== FIND BY NAME =========================
// func TestFindByNameProducts(t *testing.T) {
// 	db := setupTestDB_Products(t)
// 	repo := NewProductsRepositoryImpl(db)

// 	// Test FindByName for Name
// 	foundUsers, err := repo.FindByName("cicilan 6 bulan")
// 	require.NoError(t, err)

// 	// Expected Products with Name
// 	expectedUsers := model.Products{
// 		ID:          1,
// 		Name:        "cicilan 6 bulan",
// 		Amount:      1000000,
// 		Installment: 6,
// 		Bunga:       0.2,
// 		Created_At:  time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
// 	}

// 	require.Equal(t, expectedUsers, foundUsers)

// 	// Test FindByUsername for non-existing username "Nonexistent"
// 	_, err = repo.FindByName("Nonexistent")
// 	require.Error(t, err)
// 	require.EqualError(t, err, "invalid name")
// }

// // ================== FIND ALL =========================
// func TestFindAll(t *testing.T) {
// 	db := setupTestDB_Products(t)
// 	repo := NewProductsRepositoryImpl(db)

// 	// Create multiple Productss in the database
// 	Productss := []model.Products{
// 		{
// 			ID:          1,
// 			Name:        "cicilan 6 bulan",
// 			Amount:      1000000,
// 			Installment: 6,
// 			Bunga:       0.2,
// 			Created_At:  time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
// 		},
// 		{
// 			ID:          2,
// 			Name:        "cicilan 12 bulan",
// 			Amount:      2000000,
// 			Installment: 12,
// 			Bunga:       0.4,
// 			Created_At:  time.Date(2023, 5, 26, 0, 0, 0, 0, time.Local),
// 		},
// 		// Add more Productss if needed
// 	}

// 	// Test FindAll
// 	foundUserss, err := repo.FindAll()
// 	require.NoError(t, err)
// 	require.Equal(t, len(Productss), len(foundUserss))

// 	// Compare each Products in the expected list with the found Productss
// 	for _, expectedUsers := range Productss {
// 		found := false
// 		for _, actualProducts := range foundUserss {
// 			if expectedUsers.ID == actualProducts.ID {
// 				require.Equal(t, expectedUsers, actualProducts)
// 				found = true
// 				break
// 			}
// 		}
// 		require.True(t, found, "Products not found: ID %d", expectedUsers.ID)
// 	}
// }
