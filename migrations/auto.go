package migrations

import (
	"gorm.io/gorm"
	"judo/internal/link"
	"judo/internal/user"
	"log"
)

func RunMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&link.Link{}, &user.User{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
