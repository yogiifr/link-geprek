// Package services for GenerateShortCode
package services

import (
	"link-geprek/backend/internal/models"
	"math/rand"
	"time"

	"github.com/speps/go-hashids"
	"gorm.io/gorm"
)

type Shortener struct {
	db *gorm.DB
}

func NewShortener(db *gorm.DB) *Shortener {
	return &Shortener{db: db}
}

func (s *Shortener) GenerateShortCode() (string, error) {
	hd := hashids.NewData()
	hd.Salt = "geprek-salt"
	hd.MinLength = 6
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}

	// Ran Num buat encode
	randInts := []int{int(time.Now().Unix()), rand.Intn(10000)}
	code, err := h.Encode(randInts)
	if err != nil {
		return "", err
	}

	// Collisions check
	var count int64
	s.db.Model(&models.URL{}).Where("short_code = ?", code).Count(&count)
	if count > 0 {
		return s.GenerateShortCode()
	}

	return code, nil
}
