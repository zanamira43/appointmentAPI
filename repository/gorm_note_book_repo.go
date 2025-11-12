package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormNoteBookRepostiry struct {
	DB *gorm.DB
}

func NewGormNoteBookRepostiry(db *gorm.DB) *GormNoteBookRepostiry {
	return &GormNoteBookRepostiry{DB: db}
}

// create
func (r *GormNoteBookRepostiry) CreateNoteBook(noteBook *dto.NoteBook) error {
	return r.DB.Create(&noteBook).Error
}

// get all
func (r *GormNoteBookRepostiry) GetAllNoteBooks(page, limit int, search string) ([]models.NoteBook, int64, error) {
	var noteBooks []models.NoteBook
	var total int64

	// create blank query to build upon
	query := r.DB.Model(&models.NoteBook{})

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("content LIKE ?", searchPattern)
	}

	// get total number of notebooks
	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	if page > 0 && limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err = query.Order("id desc").Find(&noteBooks).Error
	if err != nil {
		return nil, 0, err
	}
	return noteBooks, total, nil
}

// get single by id
func (r *GormNoteBookRepostiry) GetNoteBookByID(id uint) (*models.NoteBook, error) {
	var noteBook models.NoteBook
	err := r.DB.First(&noteBook, id).Error
	if err != nil {
		return nil, err
	}
	return &noteBook, nil
}

// update
func (r *GormNoteBookRepostiry) UpdateNoteBookByID(id uint, noteBook *dto.NoteBook) (*models.NoteBook, error) {
	var nb models.NoteBook
	err := r.DB.First(&nb, id).Error
	if err != nil {
		return nil, err
	}

	if noteBook.Content != "" {
		nb.Content = noteBook.Content
	}

	err = r.DB.Save(&nb).Error
	if err != nil {
		return nil, err
	}

	return &nb, nil
}

// delete
func (r *GormNoteBookRepostiry) DeleteNoteBookByID(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&models.NoteBook{}).Error
}
