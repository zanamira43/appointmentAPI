package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormProblemRepository struct {
	DB *gorm.DB
}

func NewGormProblemRepository(db *gorm.DB) *GormProblemRepository {
	return &GormProblemRepository{DB: db}
}

func (r *GormProblemRepository) CreateProblem(problem *dto.Problem) error {
	var p models.Problem

	p.PatientID = problem.PatientID
	p.MianpProblems = problem.MianpProblems
	p.SecondaryProblems = problem.SecondaryProblems
	p.NeedSessionsCount = problem.NeedSessionsCount
	p.SessionPrice = problem.SessionPrice
	p.PatientImage = problem.PatientImage
	p.Details = problem.Details
	return r.DB.Create(&p).Error
}

func (r *GormProblemRepository) GetAllProblems() ([]models.Problem, error) {
	var problems []models.Problem
	err := r.DB.Find(&problems).Error
	if err != nil {
		return nil, err
	}
	return problems, nil
}

func (r *GormProblemRepository) GetProblemByID(id uint) (*models.Problem, error) {
	var problem models.Problem
	err := r.DB.First(&problem, id).Error
	if err != nil {
		return nil, err
	}
	return &problem, nil
}

func (r *GormProblemRepository) GetProblemByPatientId(patientId uint) (*models.Problem, error) {
	var problem models.Problem
	err := r.DB.Where("patient_id = ?", patientId).First(&problem).Error
	if err != nil {
		return nil, err
	}
	return &problem, nil
}

func (r *GormProblemRepository) UpdateProblemByID(id uint, problem *dto.Problem) (*models.Problem, error) {
	var p models.Problem
	err := r.DB.First(&p, id).Error
	if err != nil {
		return nil, err
	}

	if problem.PatientID != 0 {
		p.PatientID = problem.PatientID
	}

	if len(problem.MianpProblems) != 0 {
		p.MianpProblems = problem.MianpProblems
	}

	if len(problem.SecondaryProblems) != 0 {
		p.SecondaryProblems = problem.SecondaryProblems
	}

	if problem.IsDollarPayment != p.IsDollarPayment {
		p.IsDollarPayment = problem.IsDollarPayment
	}

	if problem.NeedSessionsCount != 0 {
		p.NeedSessionsCount = problem.NeedSessionsCount
	}

	if problem.SessionPrice != 0 {
		p.SessionPrice = problem.SessionPrice
	}

	p.PatientImage = problem.PatientImage

	if problem.Details != "" {
		p.Details = problem.Details
	}

	err = r.DB.Save(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *GormProblemRepository) DeleteProblemByID(id uint) error {
	return r.DB.Delete(&models.Problem{}, id).Error
}
