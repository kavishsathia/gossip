package helpers

import (
	"backend/models"

	"gorm.io/gorm"
)

func IncrementAura(userId int) {
	db, err := OpenDatabase()
	if err != nil {
		print("Error when incrementing aura")
		return
	}

	userCountUpdateResult := db.Model(&models.User{}).
		Where("id = ?", uint(userId)).
		Update("aura", gorm.Expr("aura + ?", 1))

	if userCountUpdateResult.Error != nil {
		print("Error when incrementing aura")
		return
	}
}
