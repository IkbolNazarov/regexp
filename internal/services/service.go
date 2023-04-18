package services

import (
	"regexp/internal/models"
	"regexp/internal/repository"
	"errors"
	"regexp"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep}
}

func (s *Services) AddService(service *models.CardRule) error {
	err := s.Repository.AddService(service)
	if err != nil {
		return err
	}
	return nil
}


func (s *Services) GetService(card string) ([]*models.CardRule, error) {
	var counter int64 = 0
	var response []*models.CardRule

	CardRule, l, err := s.Repository.GetCardRule()
	if err != nil {
		return nil, err
	}
	var i int64
	for i = 0; i < l; i++ {
		match, _ := regexp.MatchString(CardRule[i].Regexp, card)
		if match {
			response = append(response, CardRule[i])
			counter++
		}
	}
	if counter == 0 {
		return nil, errors.New("Wrong Card!")
	}

	return response, nil
}



func (s *Services) AddAgent(agent *models.Agents) error {
	err := s.Repository.AddAgent(agent)
	if err != nil {
		return err
	}
	return nil
}