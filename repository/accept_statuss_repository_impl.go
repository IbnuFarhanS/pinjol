package repository

import (
	"errors"
	"time"

	"github.com/IbnuFarhanS/pinjol/helper"
	"github.com/IbnuFarhanS/pinjol/model"
	"gorm.io/gorm"
)

type AcceptStatusRepositoryImpl struct {
	Db *gorm.DB
}

func NewAcceptStatusRepositoryImpl(Db *gorm.DB) AcceptStatusRepository {
	return &AcceptStatusRepositoryImpl{Db: Db}
}

// Delete implements AcceptStatusRepository
func (r *AcceptStatusRepositoryImpl) Delete(id int64) (model.AcceptStatus, error) {
	var bor model.AcceptStatus
	result := r.Db.Where("id = ?", id).Delete(&bor)
	helper.ErrorPanic(result.Error)
	return bor, nil
}

// FindAll implements AcceptStatusRepository
func (r *AcceptStatusRepositoryImpl) FindAll() ([]model.AcceptStatus, error) {
	var bor []model.AcceptStatus
	results := r.Db.Find(&bor)
	helper.ErrorPanic(results.Error)
	return bor, nil
}

// FindById implements AcceptStatusRepository
func (r *AcceptStatusRepositoryImpl) FindById(id int64) (model.AcceptStatus, error) {
	var bor model.AcceptStatus
	result := r.Db.Find(&bor, "id = ?", id)
	if result.Error != nil {
		return bor, errors.New("AcceptStatus is not found")
	}
	return bor, nil
}

// Save implements AcceptStatusRepository
func (r *AcceptStatusRepositoryImpl) Save(newAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	currentTime := time.Now()
	newAcceptStatus.Created_At = currentTime
	result := r.Db.Create(&newAcceptStatus)
	helper.ErrorPanic(result.Error)
	return newAcceptStatus, nil
}

// Update implements AcceptStatusRepository
func (r *AcceptStatusRepositoryImpl) Update(updatedAcceptStatus model.AcceptStatus) (model.AcceptStatus, error) {
	result := r.Db.Model(&model.AcceptStatus{}).Where("id = ?", updatedAcceptStatus.ID).Updates(updatedAcceptStatus)
	helper.ErrorPanic(result.Error)
	return updatedAcceptStatus, nil
}
