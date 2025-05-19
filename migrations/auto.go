package migrations

import (
	"gorm.io/gorm"
	"judo/internal/link"
	typesimpo "judo/internal/types"
	"log"
)

func RunMigrations(db *gorm.DB) {
	if err := db.AutoMigrate(&link.Link{}, &typesimpo.User{}, &typesimpo.Stat{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
