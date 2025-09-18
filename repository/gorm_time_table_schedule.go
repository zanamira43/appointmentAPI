package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormTimeTableRepository struct {
	DB *gorm.DB
}

func NewGormTimeTableRepository(db *gorm.DB) *GormTimeTableRepository {
	return &GormTimeTableRepository{DB: db}
}

// create appointment repository
func (r *GormTimeTableRepository) CreateTimeTables(aps *dto.TimeTable, UserID uint) error {
	// create a new AppointmentSchedule struct

	request := models.TimeTable{
		PatientName: aps.PatientName,
		WeekDay:     aps.WeekDay,
		StartTime:   aps.StartTime,
		EndTime:     aps.EndTime,
		UserId:      UserID,
	}

	if aps.PatientID != nil {
		request.PatientID = aps.PatientID
	}

	return r.DB.Create(&request).Error
}

// get all appointments repository
func (r *GormTimeTableRepository) GetAllTimeTables(page, limit int, search string) ([]models.TimeTable, int64, error) {
	var tts []models.TimeTable
	var total int64

	// create blank query to build upon
	query := r.DB.Model(&models.TimeTable{})

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("patient_name LIKE ? OR week_day LIKE ?", searchPattern, searchPattern)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err = query.Preload("User").Preload("Patient").Order("id desc").Find(&tts).Error
	if err != nil {
		return nil, 0, err
	}

	return tts, total, nil
}

// get single appointment repository
func (r *GormTimeTableRepository) GetTimeTableByID(id uint) (*models.TimeTable, error) {
	var tt models.TimeTable
	err := r.DB.Preload("User").First(&tt, id).Error
	if err != nil {
		return nil, err
	}
	return &tt, nil
}

// update appointment repository
func (r *GormTimeTableRepository) UpdateTimeTableByID(id uint, dtt *dto.TimeTable, UserID uint) (*models.TimeTable, error) {
	var tt models.TimeTable
	err := r.DB.First(&tt, id).Error
	if err != nil {
		return nil, err
	}

	if dtt.PatientID != nil {
		tt.PatientID = dtt.PatientID
	}

	if dtt.PatientName != "" {
		tt.PatientName = dtt.PatientName
	}

	if len(dtt.WeekDay) != 0 {
		tt.WeekDay = dtt.WeekDay
	}

	if dtt.StartTime != "" {
		tt.StartTime = dtt.StartTime
	}

	if dtt.EndTime != "" {
		tt.EndTime = dtt.EndTime
	}

	if UserID != 0 {
		tt.UserId = UserID
	}

	err = r.DB.Save(&tt).Error
	if err != nil {
		return nil, err
	}
	return &tt, nil
}

// delete appointment repository
func (r *GormTimeTableRepository) DeleteTimeTableByID(id uint) error {
	return r.DB.Delete(&models.TimeTable{}, id).Error
}
