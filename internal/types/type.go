package typesimpo

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"judo/internal/link"
)

type Stat struct {
	gorm.Model
	LinkID uint           `json:"link_id"`
	Clicks uint           `json:"click"`
	Data   datatypes.Date `json:"data"`
}

type User struct {
	gorm.Model
	Email    string `gorm:"index"`
	Password string
	Name     string
}

type LinkCreateRequest struct {
	URL string `json:"url" validate:"required,url"`
}

type LinkUpdateRequest struct {
	URL  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

type LinksResponse struct {
	ID   int
	URL  string
	Hash string
}

type AllLinksResponse struct {
	Links []*link.Link `json:"links"`
	Count int64        `json:"count"`
}
