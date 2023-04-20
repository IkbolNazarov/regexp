package repository

import (
	"cards/internal/db"
	"cards/internal/models"


	"gorm.io/gorm"
)

type Repository struct {
	connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{connection: conn}
}

func (r *Repository) AddService(service *models.Regexp) error {
	tx := db.DataB.Table(models.GetRuleTable()).Create(&service)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *Repository) GetCardRule() ([]*models.Regexp, int, error) {
	var cardRule []*models.Regexp
	tx := db.DataB.Table("regexp").Preload("Agent").Joins("join agents on regexp.agent_id = agents.id").Order("agents.sort DESC").Find(&cardRule)
	//tx := db.DataB.Raw("Select * from regexp join agent on cards.agent=agent.id order by agent.sort")
	//tx := db.DataB.Preload("Agent").Order("agents.sort").Find(&cardRule)
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
