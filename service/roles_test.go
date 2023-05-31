package service_test

import (
	"errors"
	"testing"
	"time"

	"github.com/IbnuFarhanS/pinjol/model"
	"github.com/IbnuFarhanS/pinjol/service"
	"github.com/stretchr/testify/assert"
)

type mockRoleRepository struct{}

func (m *mockRoleRepository) Delete(id uint) (model.Role, error) {
	// Simulate deleting a role
	if id == 1 {
		role := model.Role{
			ID:        1,
			Name:      "Admin",
			CreatedAt: time.Now(),
		}
		return role, nil
	}
	return model.Role{}, errors.New("role not found")
}

func (m *mockRoleRepository) FindAll() ([]model.Role, error) {
	// Simulate finding all roles
	roles := []model.Role{
		{
			ID:        1,
			Name:      "Admin",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "User",
			CreatedAt: time.Now(),
		},
	}
	return roles, nil
}

func (m *mockRoleRepository) FindById(id uint) (model.Role, error) {
	// Simulate finding a role by ID
	if id == 1 {
		role := model.Role{
			ID:        1,
			Name:      "Admin",
			CreatedAt: time.Now(),
		}
		return role, nil
	}
	return model.Role{}, errors.New("role not found")
}

func (m *mockRoleRepository) FindByName(name string) (model.Role, error) {
	// Simulate finding a role by name
	if name == "Admin" {
		role := model.Role{
			ID:        1,
			Name:      "Admin",
			CreatedAt: time.Now(),
		}
		return role, nil
	}
	return model.Role{}, errors.New("role not found")
}

func (m *mockRoleRepository) Save(newRole model.Role) (model.Role, error) {
	// Simulate saving a new role
	role := model.Role{
		Name:      newRole.Name,
		CreatedAt: newRole.CreatedAt,
	}
	return role, nil
}

func (m *mockRoleRepository) Update(updatedRole model.Role) (model.Role, error) {
	// Simulate updating a role
	role := model.Role{
		ID:        updatedRole.ID,
		Name:      updatedRole.Name,
		CreatedAt: updatedRole.CreatedAt,
	}
	return role, nil
}

func TestRoleService(t *testing.T) {
	repo := &mockRoleRepository{}
	roleService := service.NewRoleServiceImpl(repo)

	t.Run("Delete_ValidRole", func(t *testing.T) {
		id := uint(1)

		role, err := roleService.Delete(id)
		assert.NoError(t, err)
		assert.Equal(t, id, role.ID)
	})

	t.Run("Delete_InvalidRole", func(t *testing.T) {
		id := uint(2)

		_, err := roleService.Delete(id)
		assert.Error(t, err)
		assert.EqualError(t, err, "role not found")
	})

	t.Run("FindAll", func(t *testing.T) {
		roles, err := roleService.FindAll()
		assert.NoError(t, err)
		assert.Len(t, roles, 2)
	})

	t.Run("FindById_ValidRole", func(t *testing.T) {
		id := uint(1)

		role, err := roleService.FindById(id)
		assert.NoError(t, err)
		assert.Equal(t, id, role.ID)
	})

	t.Run("FindById_InvalidRole", func(t *testing.T) {
		id := uint(2)

		_, err := roleService.FindById(id)
		assert.Error(t, err)
		assert.EqualError(t, err, "role not found")
	})

	t.Run("FindByName_ValidRole", func(t *testing.T) {
		name := "Admin"

		role, err := roleService.FindByName(name)
		assert.NoError(t, err)
		assert.Equal(t, name, role.Name)
	})

	t.Run("FindByName_InvalidRole", func(t *testing.T) {
		name := "Superadmin"

		_, err := roleService.FindByName(name)
		assert.Error(t, err)
		assert.EqualError(t, err, "role not found")
	})

	t.Run("Save", func(t *testing.T) {
		newRole := model.Role{
			Name:      "New Role",
			CreatedAt: time.Now(),
		}

		role, err := roleService.Save(newRole)
		assert.NoError(t, err)
		assert.Equal(t, newRole.Name, role.Name)
	})

	t.Run("Update", func(t *testing.T) {
		updatedRole := model.Role{
			ID:        1,
			Name:      "Updated Role",
			CreatedAt: time.Now(),
		}

		role, err := roleService.Update(updatedRole)
		assert.NoError(t, err)
		assert.Equal(t, updatedRole.Name, role.Name)
	})
}
