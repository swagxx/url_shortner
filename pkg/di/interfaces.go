package di

import (
	"judo/internal/link"
	typesimpo "judo/internal/types"
	"judo/pkg/response"
	"time"
)

type IStatRepository interface {
	AddClick(linkId uint)
	GetStats(by string, from, to time.Time) ([]response.StatResponse, error)
}

type ILinkRepository interface {
	Create(link *link.Link) (*link.Link, error)
	GetByHash(hash string) (*link.Link, error)
	UpdateLink(link *link.Link) (*link.Link, error)
	DeleteLink(id uint) error
	FindById(id uint) error
	CountLinks() (int64, error)
	GetLinks(limit, offset uint) ([]*link.Link, int64, error)
}

type IUserRepository interface {
	Create(user *typesimpo.User) (*typesimpo.User, error)
	Find(email string) (*typesimpo.User, error)
}
