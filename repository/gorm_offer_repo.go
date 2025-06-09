package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormOfferRepository struct {
	DB *gorm.DB
}

func NewGormOfferRepository(db *gorm.DB) *GormOfferRepository {
	return &GormOfferRepository{DB: db}
}

// insert new offer data into sql database
func (r *GormOfferRepository) CreateOffers(offer *dto.Offer) error {
	return r.DB.Create(&offer).Error
}

// get all offers data from sql database
func (r *GormOfferRepository) GetAllOffers() ([]models.Offer, error) {
	var offers []models.Offer
	err := r.DB.Find(&offers).Error
	if err != nil {
		return nil, err
	}
	return offers, nil
}

// get offer data by id from sql database
func (r *GormOfferRepository) GetOfferByID(id uint) (*models.Offer, error) {
	var offer models.Offer
	err := r.DB.Where("id = ?", id).First(&offer).Error
	if err != nil {
		return nil, err
	}
	return &offer, nil
}

// update offer data by id from sql database
func (r *GormOfferRepository) UpdateOfferByID(id uint, dtoOffer *dto.Offer) (*models.Offer, error) {
	var offer models.Offer
	err := r.DB.Where("id = ?", id).First(&offer).Error
	if err != nil {
		return nil, err
	}
	if dtoOffer.Title != "" {
		offer.Title = dtoOffer.Title
	}
	if dtoOffer.ServiceTypeID != 0 {
		offer.ServiceTypeID = dtoOffer.ServiceTypeID
	}
	if dtoOffer.Price != 0 {
		offer.Price = dtoOffer.Price
	}

	err = r.DB.Save(&offer).Error
	if err != nil {
		return nil, err
	}
	return &offer, nil
}

// delete offer data by id from sql database
func (r *GormOfferRepository) DeleteOfferByID(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&dto.Offer{}).Error
}
