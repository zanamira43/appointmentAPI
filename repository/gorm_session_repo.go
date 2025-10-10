package repository

import (
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
func (r *GormSessionRepository) CreateSession(session *dto.Session) error {
	return r.DB.Create(&session).Error
}

// get all sessions data from sql database
func (r *GormSessionRepository) GetAllSessions(page, limit int, search string) ([]models.Session, int64, error) {
	var sessions []models.Session
	var total int64

	// create blank query to build upon
	query := r.DB.Model(&models.Session{})

	// search by session subject
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("subtject LIKE ? ", searchPattern)
	}

	// get total number of sessions
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err = query.Order("id desc").Find(&sessions).Error
	if err != nil {
		return nil, 0, err
	}
	return sessions, total, nil
}

// get all sessions data from sql database by patient id
func (r *GormSessionRepository) GetSessionsByPatientID(page, limit int, search string, patientId uint) ([]models.Session, int64, error) {

	var sessions []models.Session
	var total int64

	// create blank query to build upon
	query := r.DB.Where("patient_id = ?", patientId).Model(&models.Session{})

	// search by session subject
	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("subject LIKE ? ", searchPattern)
	}

	// get total number of sessions
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err = query.Find(&sessions).Error
	if err != nil {
		return nil, 0, err
	}
	return sessions, total, nil

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
func (r *GormSessionRepository) UpdateSession(id uint, sessionDto *dto.Session) (*models.Session, error) {
	var session models.Session
	err := r.DB.Where("id = ?", id).First(&session).Error
	if err != nil {
		return nil, err
	}

	if sessionDto.PatientID != 0 {
		session.PatientID = sessionDto.PatientID
	}
	if sessionDto.Subject != "" {
		session.Subject = sessionDto.Subject
	}

	if sessionDto.CommunicationTypes != "" {
		session.CommunicationTypes = sessionDto.CommunicationTypes
	}

	if sessionDto.SessionDate != "" {
		session.SessionDate = sessionDto.SessionDate
	}

	if session.Detail != "" {
		session.Detail = sessionDto.Detail
	}
	if sessionDto.Status != "" {
		session.Status = sessionDto.Status
	}

	err = r.DB.Save(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// delete offer data by id from sql database
func (r *GormSessionRepository) DeleteSession(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&dto.Session{}).Error
}
