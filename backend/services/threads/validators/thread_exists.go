package validators

import (
	"backend/helpers"
	"backend/models"
)

func ThreadExists(id int) bool {
	db, err := helpers.OpenDatabase()
	if err != nil {
		return false
	}

	var exists bool
	err = db.Model(&models.Thread{}).
		Select("count(*) > 0").
		Where("id = ?", id).
		Find(&exists).Error

	if err != nil {
		return false
	}

	return exists
}
