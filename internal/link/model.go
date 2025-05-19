package link

import (
	"crypto/rand"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"math/big"
)

const (
	lenHash = 7
)

type Link struct {
	gorm.Model
	URL   string `json:"url"`
	Hash  string `json:"hash" gorm:"uniqueIndex"`
	Stats []Stat `gorm:"constraint:OnUpdate:CASCADE, OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		URL: url,
	}
	link.GenerateHash()
	return link
}
func (l *Link) GenerateHash() {
	l.Hash = generateHash(lenHash)
}

func generateHash(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[idx.Int64()]
	}
	return string(b)
}

type Stat struct {
	gorm.Model
	LinkID uint           `json:"link_id"`
	Clicks uint           `json:"click"`
	Data   datatypes.Date `json:"data"`
}
