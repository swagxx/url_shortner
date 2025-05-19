package stat

import (
	"gorm.io/datatypes"
	typesimpo "judo/internal/types"
	"judo/pkg/constants"
	"judo/pkg/db"
	"judo/pkg/response"
	"time"
)

type StatRepository struct {
	*db.DB
}

func NewStatRepository(db *db.DB) *StatRepository {
	return &StatRepository{db}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat typesimpo.Stat
	currentDate := datatypes.Date(time.Now())
	repo.DB.Find(&stat, "link_id = ? and data = ?", linkId, currentDate)
	if stat.ID == 0 {
		repo.DB.Create(&typesimpo.Stat{
			LinkID: linkId,
			Clicks: 1,
			Data:   currentDate,
		})
		return

	}
	stat.Clicks++
	repo.DB.Save(&stat)

}

func (repo *StatRepository) GetStats(by string, from, to time.Time) ([]response.StatResponse, error) {
	var stats []response.StatResponse
	var selectQuery string
	
	switch by {
	case constants.CheckDay:
		selectQuery = "to_char(data, 'YYYY-MM-DD') as period, sum(clicks)"
	case constants.CheckMonth:
		selectQuery = "to_char(data, 'YYYY-MM') as period, sum(clicks)"
	}
	err := repo.DB.Model(typesimpo.Stat{}).Select(selectQuery).
		Where("data BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats).
		Error

	if err != nil {
		return nil, err
	}

	return stats, nil
}
