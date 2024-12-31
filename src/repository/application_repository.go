package repository

import (
    "github.com/Mohamed-Kalandar-Sulaiman/Identity_Service/src/database"
    "github.com/Mohamed-Kalandar-Sulaiman/Identity_Service/src/models"
)

func GetApplicationContext(appName string) (*models.Application, error) {
    var app models.Application
    result := database.DB.Where("app_name = ?", appName).First(&app)
    return &app, result.Error
}

func GetPublicKey(appName string) (string, error) {
    var app models.Application
    result := database.DB.Where("app_name = ?", appName).First(&app)
    if result.Error != nil {
        return "", result.Error
    }
    return app.PublicKey, nil
}
