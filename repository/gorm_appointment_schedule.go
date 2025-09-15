package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormAppointmentRepository struct {
	DB *gorm.DB
}

func NewGormAppointmentScheduleRepository(db *gorm.DB) *GormAppointmentRepository {
	return &GormAppointmentRepository{DB: db}
}

// create appointment repository
func (r *GormAppointmentRepository) CreateAppointments(aps *dto.AppointmentSchedule, UserID uint) error {
	// create a new AppointmentSchedule struct
	request := models.AppointmentSchedule{
		PatientID:   &aps.PatientID,
		PatientName: aps.PatientName,
		WeekDay:     aps.WeekDay,
		Date:        aps.Date,
		StartTime:   aps.StartTime,
		EndTime:     aps.EndTime,
		UserId:      UserID,
	}

	return r.DB.Create(&request).Error
}

// get all appointments repository
func (r *GormAppointmentRepository) GetAllAppointments(page, limit int, search string) ([]models.AppointmentSchedule, int64, error) {
	var appointments []models.AppointmentSchedule
	var total int64

	// create blank query to build upon
	query := r.DB.Model(&models.AppointmentSchedule{})

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("patient_name LIKE ? OR date LIKE ?", searchPattern, searchPattern)
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err = query.Preload("User").Preload("Patient").Order("date desc").Find(&appointments).Error
	if err != nil {
		return nil, 0, err
	}

	return appointments, total, nil
}

// get single appointment repository
func (r *GormAppointmentRepository) GetAppointmentByID(id uint) (*models.AppointmentSchedule, error) {
	var appointment models.AppointmentSchedule
	err := r.DB.Preload("User").First(&appointment, id).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

// update appointment repository
func (r *GormAppointmentRepository) UpdateAppointmentByID(id uint, ap *dto.AppointmentSchedule, UserID uint) (*models.AppointmentSchedule, error) {
	var appointment models.AppointmentSchedule
	err := r.DB.First(&appointment, id).Error
	if err != nil {
		return nil, err
	}

	if ap.PatientID != 0 {
		appointment.PatientID = &ap.PatientID
	}

	if ap.PatientName != "" {
		appointment.PatientName = ap.PatientName
	}

	if ap.WeekDay != "" {
		appointment.WeekDay = ap.WeekDay
	}

	if ap.Date != "" {
		appointment.Date = ap.Date
	}

	if ap.StartTime != "" {
		appointment.StartTime = ap.StartTime
	}

	if ap.EndTime != "" {
		appointment.EndTime = ap.EndTime
	}

	if ap.Status != "" {
		appointment.Status = ap.Status
	}

	if UserID != 0 {
		appointment.UserId = UserID
	}

	err = r.DB.Save(&appointment).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

// delete appointment repository
func (r *GormAppointmentRepository) DeleteAppointmentByID(id uint) error {
	return r.DB.Delete(&models.AppointmentSchedule{}, id).Error
}
