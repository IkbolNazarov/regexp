package repository

import (
	"regexp/internal/db"
	"regexp/internal/models"

	"gorm.io/gorm"
)

type Repository struct {
	connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{connection: conn}
}

func (r *Repository) AddService(service *models.CardRule) error {
	tx := db.DataB.Table(models.GetRuleTable()).Create(&service)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *Repository) GetCardRule() ([]*models.CardRule, int, error) {
	var cardRule []*models.CardRule
	tx := db.DataB.Preload("Agent").Find(&cardRule)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	var l int64
	tx = db.DataB.Table(models.GetRuleTable()).Count(&l)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	return cardRule, len(cardRule), nil
}

func (r *Repository) AddAgent(card *models.Agents) error {
	tx := db.DataB.Create(&card)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
