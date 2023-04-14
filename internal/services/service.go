package services

import (
	"admin/internal/models"
	"admin/internal/repository"
	"errors"
	"log"
	"regexp"
)

type Services struct {
	Repository *repository.Repository
}

func NewServices(rep *repository.Repository) *Services {
	return &Services{Repository: rep}
}

// func (s *Services) AddUser(card *models.Card) (err error) {

// 	CardRule, l := s.Repository.GetCardRule()
// 	var i int64
// 	for i=0; i<l; i++{
// 		log.Println(CardRule[i].Regexp)
// 		match, _ := regexp.MatchString(CardRule[i].Regexp,card.Card1)
// 		if match {
// 			return errors.New("wrong card number!")
// 		}
// 	}
// 	err = s.Repository.AddUser(card)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (s *Services) AddService(service *models.CardRule) error {
	err := s.Repository.AddService(service)
	if err != nil {
		return err
	}
	return nil
}

func (s *Services) AddAgent(agent *models.Agents) error {
	err := s.Repository.AddAgent(agent)
	if err != nil {
		return err
	}
	return nil
}

func (s *Services) GetService(card string) ([]*models.CardRule, error) {
	counter := 0
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

		log.Println(l)
		log.Println(match)
		log.Println(CardRule[i].Regexp)
	}
	if counter == 0 {
		return nil, errors.New("Wrong Card!")
	}
	log.Println(counter)
	return response, nil
}
