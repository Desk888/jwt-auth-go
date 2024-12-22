package initializers

import (
	"github.com/Desk888/go-jwt/internal/models"
)

func MigrateTables() {
	DB.AutoMigrate(&models.User{})
}