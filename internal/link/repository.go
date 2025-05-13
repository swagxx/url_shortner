package link

import (
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

func (repo *LinkRepository) GetAll() ([]*Link, error) {
	var links []*Link
	result := repo.DataBase.DB.Find(&links)
	if result.Error != nil {
		return nil, result.Error
	}
	return links, nil
}
