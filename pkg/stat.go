package pkg

import (
	"errors"
	"judo/configs"
	"judo/internal/stat"
	"judo/pkg/constants"
	"judo/pkg/di"
	"judo/pkg/handlerset"
	"log"
	"net/http"
	"time"
)

type StatHandler struct {
	Config   *configs.Config
	StatRepo di.IStatRepository
}

func NewStatHandler(conf *configs.Config, stat *stat.StatRepository) *StatHandler {
	return &StatHandler{
		Config:   conf,
		StatRepo: stat,
	}
}

func (s *StatHandler) StatByDate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from, to, err := parseTime(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		by := r.URL.Query().Get("by")

		if by != constants.CheckDay && by != constants.CheckMonth {
			http.Error(w, "by must be either 'day' or 'month'", http.StatusBadRequest)
			return
		}
		stats, err := s.StatRepo.GetStats(by, from, to)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		log.Println(stats)
		handlerset.HandlerSet(w, stats, http.StatusOK)

	}
}

func parseTime(r *http.Request) (time.Time, time.Time, error) {
	from, err := time.Parse("2006-01-02", r.URL.Query().Get("from"))
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("data (from) format error")
	}
	to, err := time.Parse("2006-01-02", r.URL.Query().Get("to"))
	if err != nil {
		return time.Time{}, time.Time{}, errors.New("data (to) format error")
	}
	return from, to, nil
}
