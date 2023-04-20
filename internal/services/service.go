package services

import (
	"errors"
	"fmt"
	"log"
	"regexp"

	"cards/internal/models"
	"cards/internal/repository"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep}
}

func (s *Services) AddService(service *models.Regexp) error {
	err := s.Repository.AddService(service)
	if err != nil {
		return err
	}
	return nil
}

func (s *Services) GetService(card string) ([]*models.Regexp, error) {
	var counter int64 = 0
	var response []*models.Regexp
	CardRule, l, err := s.Repository.GetCardRule()
	if err != nil {
		return nil, fmt.Errorf("GetCardRule %v", err)
	}
	log.Println(l)
	for i := 0; i < l; i++ {
		log.Println("6")
		match, _ := regexp.MatchString("^4[0-9]{12}(?:[0-9]{3})?$", "4000001234567899")
		log.Println(match)
		log.Println("6")
		//match, _ := regexp.MatchString(CardRule[i].RegularExp, card)
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
