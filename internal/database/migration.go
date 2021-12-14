package database

import (
	"github.com/cmcki85/RESTurant-Full-Stack/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - auto migrates the database and creates the comment table.
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
