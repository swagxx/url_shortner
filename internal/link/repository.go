package link

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"judo/pkg/db"
)

type LinkRepository struct {
	DataBase *db.DB
}

func NewLinkRepository(dataBase *db.DB) *LinkRepository {
	return &LinkRepository{
		DataBase: dataBase,
	}
}

func (repo *LinkRepository) Create(link *Link) (*Link, error) {
	result := repo.DataBase.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	result := repo.DataBase.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) UpdateLink(link *Link) (*Link, error) {
	result := repo.DataBase.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) DeleteLink(id uint) error {
	if err := repo.FindById(id); err != nil {
		return err
	}
	result := repo.DataBase.DB.Delete(&Link{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *LinkRepository) FindById(id uint) error {
	result := repo.DataBase.DB.First(&Link{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *LinkRepository) CountLinks() (int64, error) {
	var count int64
	err := repo.DataBase.DB.Model(&Link{}).
		Where("deleted_at is null").
		Count(&count).Error

	if err != nil {
		return 0, err
	}
	return count, nil
}

func (repo *LinkRepository) GetLinks(limit, offset uint) ([]*Link, int64, error) {
	var links []*Link
	query := repo.DataBase.DB.Model(&Link{}).Where("deleted_at is null").Session(&gorm.Session{})

	err := query.
		Order("links.id desc").
		Limit(int(limit)).
		Offset(int(offset)).
		Find(&links).Error

	if err != nil {
		return nil, 0, err
	}

	count, err := repo.CountLinks()
	if err != nil {
		return nil, 0, err
	}

	return links, count, nil

}
