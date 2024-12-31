package repository

import (
	"github.com/Mohamed-Kalandar-Sulaiman/Identity_Service/src/database"
	"github.com/Mohamed-Kalandar-Sulaiman/Identity_Service/src/models"

	"gorm.io/gorm"
)

func CreateUser(email, password, context string) error {
    user := models.User{
        Email:    email,
        Password: password,
        Context:  context, // Store user-specific context
    }

    result := database.DB.Create(&user)
    return result.Error
}

func GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    result := database.DB.Where("email = ?", email).First(&user)
    if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
        return nil, result.Error
    }
    return &user, result.Error
}

func GetUserContext(userID uint) (*models.User, error) {
    var user models.User
    result := database.DB.First(&user, userID)
    if result.Error != nil {
        return nil, result.Error
    }
    return &user, nil
}
