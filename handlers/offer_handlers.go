package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/helpers"
	"github.com/zanamira43/appointment-api/models"
	"github.com/zanamira43/appointment-api/repository"
)

type OfferHandler struct {
	OfferRepository *repository.GormOfferRepository
}

func NewOfferHandler(offerRepo *repository.GormOfferRepository) *OfferHandler {
	return &OfferHandler{OfferRepository: offerRepo}
}

// Create New Offer
func (h *OfferHandler) CreateOffers(c echo.Context) error {
	offer := new(dto.Offer)

	if err := c.Bind(&offer); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate the request body
	if err := helpers.ValidateOffer(offer); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err := h.OfferRepository.CreateOffers(offer)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to create offer")
	}
	return c.JSON(http.StatusOK, offer)
}

// Get all offers
func (h *OfferHandler) GetAllOffers(c echo.Context) error {
	offers, err := h.OfferRepository.GetAllOffers()
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to get offers")
	}
	return c.JSON(http.StatusOK, offers)
}

// Get single offer
func (h *OfferHandler) GetOffer(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Offer Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Offer Id")
	}

	offer, err := h.OfferRepository.GetOfferByID(id)
	if err != nil {
		log.Error("Offer Not Found", err.Error())
		return c.JSON(http.StatusNotFound, "Offer Not Found")
	}
	return c.JSON(http.StatusOK, offer)
}

// update single offer by id
func (h *OfferHandler) UpdateOffer(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Offer Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Offer Id")
	}

	var dtoOffer dto.Offer
	if err := c.Bind(&dtoOffer); err != nil {
		log.Error("Invalid Request data", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Request data")
	}

	// Validate the request body
	if err := helpers.ValidateOffer(&dtoOffer); err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	offer := new(models.Offer)
	offer, err = h.OfferRepository.UpdateOfferByID(id, &dtoOffer)
	if err != nil {
		log.Error("Failed to update offer", err.Error())
		return c.JSON(http.StatusInternalServerError, "Failed to update offer")
	}
	return c.JSON(http.StatusOK, offer)
}

// delete offer by id
func (h *OfferHandler) DeleteOffer(c echo.Context) error {
	id, err := helpers.GetParam(c)
	if err != nil {
		log.Error("Invalid Offer Id", err.Error())
		return c.JSON(http.StatusBadRequest, "Invalid Offer Id")
	}
	err = h.OfferRepository.DeleteOfferByID(id)
	if err != nil {
		log.Error("Offer not found", err.Error())
		return c.JSON(http.StatusInternalServerError, "Offer not found")
	}
	return c.NoContent(http.StatusNoContent)
}
