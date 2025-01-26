package validators

import (
	"backend/helpers"
	"backend/models"
)

func UserOwnsComment(id int, userId int) bool {
	db, err := helpers.OpenDatabase()
	if err != nil {
		return false
	}

	var exists bool
	err = db.Model(&models.ThreadComment{}).
		Select("count(*) > 0").
		Where("id = ? AND user_id = ?", id, userId).
		Find(&exists).Error

	if err != nil {
		return false
	}

	return exists
}
