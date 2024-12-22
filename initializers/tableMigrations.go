package initializers

import (
	"github.com/Desk888/go-jwt/models"
)

func MigrateTables() {
	DB.AutoMigrate(&models.User{})
}