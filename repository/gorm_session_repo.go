package repository

import (
	"time"

	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormSessionRepository struct {
	DB *gorm.DB
}

func NewGormSessionRepository(db *gorm.DB) *GormSessionRepository {
	return &GormSessionRepository{DB: db}
}

// insert new offer data into sql database
func (r *GormSessionRepository) CreateSession(session *dto.SessionDto) error {
	return r.DB.Create(&session).Error
}

// get all offers data from sql database
func (r *GormSessionRepository) GetAllSessions() ([]models.Session, error) {
	var sessions []models.Session
	err := r.DB.Find(&sessions).Error
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

// get offer data by id from sql database
func (r *GormSessionRepository) GetSessionByID(id uint) (*models.Session, error) {
	var session models.Session
	err := r.DB.Where("id = ?", id).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// update offer data by id from sql database
func (r *GormSessionRepository) UpdateSession(id uint, sessionDto *dto.SessionDto) (*models.Session, error) {
	var session models.Session
	err := r.DB.Where("id = ?", id).First(&session).Error
	if err != nil {
		return nil, err
	}

	if sessionDto.PatientID != 0 {
		session.PatientID = sessionDto.PatientID
	}
	if session.Duration != 0 {
		session.Duration = sessionDto.Duration
	}
	if sessionDto.Status != "" {
		session.Status = sessionDto.Status
	}
	if sessionDto.Notes != "" {
		session.Notes = sessionDto.Notes
	}
	if sessionDto.SessionDate != (time.Time{}) {
		session.SessionDate = sessionDto.SessionDate
	}

	err = r.DB.Save(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// delete offer data by id from sql database
func (r *GormSessionRepository) DeleteSession(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&dto.SessionDto{}).Error
}
