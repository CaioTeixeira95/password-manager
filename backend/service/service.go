package service

import (
	"fmt"

	"github.com/CaioTeixeira95/password-manager/backend/model"
	"github.com/CaioTeixeira95/password-manager/backend/repository"
)

type PasswordCardService struct {
	passwordCardRepository *repository.PasswordCardRepository
}

func NewPasswordCardService(passwordCardRepository *repository.PasswordCardRepository) *PasswordCardService {
	return &PasswordCardService{passwordCardRepository: passwordCardRepository}
}

func (s *PasswordCardService) CreatePasswordCard(newPasswordCard model.PasswordCard) (*model.PasswordCard, error) {
	if err := s.passwordCardRepository.Insert(newPasswordCard); err != nil {
		return nil, fmt.Errorf("error creating a new password card: %w", err)
	}

	return &newPasswordCard, nil
}

func (s *PasswordCardService) ListPasswordCards() []model.PasswordCard {
	return s.passwordCardRepository.GetAll()
}

func (s *PasswordCardService) UpdatePasswordCard(newPasswordCard model.PasswordCard) (*model.PasswordCard, error) {
	if err := s.passwordCardRepository.Update(newPasswordCard); err != nil {
		return nil, fmt.Errorf("error updating password card: %w", err)
	}

	return &newPasswordCard, nil
}

func (s *PasswordCardService) DeletePasswordCard(passwordCardID string) error {
	if err := s.passwordCardRepository.Delete(passwordCardID); err != nil {
		return fmt.Errorf("error deleting password card: %w", err)
	}

	return nil
}
