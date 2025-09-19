package repository

import (
	"github.com/zanamira43/appointment-api/dto"
	"github.com/zanamira43/appointment-api/models"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	DB *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{DB: db}
}

// create new user
func (r *GormUserRepository) CreateUser(user *dto.User) error {
	return r.DB.Create(&user).Error
}

// get list of users
func (r *GormUserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// get single user
func (r *GormUserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// get single user
func (r *GormUserRepository) GetUserByPhone(phone string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// update single user
func (r *GormUserRepository) UpdateUser(id uint, updateuser *dto.User) (*models.User, error) {
	user, err := r.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if updateuser.FirstName != "" {
		user.FirstName = updateuser.FirstName
	}
	if updateuser.LastName != "" {
		user.LastName = updateuser.LastName
	}

	if updateuser.Email != "" {
		user.Email = updateuser.Email
	}

	if updateuser.Phone != "" {
		user.Phone = updateuser.Phone
	}

	if updateuser.Role != "" {
		user.Role = updateuser.Role
	}

	if updateuser.Active != user.Active {
		user.Active = updateuser.Active
	}

	err = r.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// update single user
func (r *GormUserRepository) UpdateUserPassword(id uint, updateuser *dto.User) (*models.User, error) {
	user, err := r.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if updateuser.Password != nil {
		user.Password = updateuser.Password
	}

	err = r.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// delete single user
func (r *GormUserRepository) DeleteUser(id uint) error {
	return r.DB.Where("id = ?", id).Delete(&models.User{}).Error
}
