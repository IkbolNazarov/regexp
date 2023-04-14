package repository

import (
	"admin/internal/db"
	"admin/internal/models"
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	Connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{Connection: conn}
}

func (r *Repository) AddUser(card *models.Agents) error {
	tx := db.DataB.Table(models.GetTableName()).Create(&card)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *Repository) AddAgent(card *models.Agents) error {
	tx := db.DataB.Table(models.GetTableName()).Create(&card)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *Repository) AddService(service *models.CardRule) error {
	tx := db.DataB.Table(models.GetRuleTable()).Create(&service)
	if tx.Error != nil {
		return tx.Error	
	}
	return nil
}

func (r *Repository) GetCardRule() ([]*models.CardRule, int64, error) {
	var cardRule []*models.CardRule
	tx := db.DataB.Table(models.GetRuleTable()).Find(&cardRule)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	log.Println()
	var l int64
	tx = db.DataB.Table(models.GetRuleTable()).Count(&l)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	return cardRule, l, nil
}	