package repository

import (
	"fmt"
	"sync"

	"github.com/CaioTeixeira95/password-manager/backend/model"
)

type PasswordCardRepository struct {
	passwordCards []model.PasswordCard
	mu            sync.Mutex
}

func NewPasswordCardRepository() *PasswordCardRepository {
	return CustomPasswordCardRepository(make([]model.PasswordCard, 0))
}

func CustomPasswordCardRepository(passwordCards []model.PasswordCard) *PasswordCardRepository {
	return &PasswordCardRepository{passwordCards: passwordCards}
}

type ErrPasswordCardAlreadyExists struct {
	id, url string
}

// Error implements error type interface.
func (e ErrPasswordCardAlreadyExists) Error() string {
	if e.id != "" {
		return fmt.Sprintf("password with ID %q already exists", e.id)
	}
	return fmt.Sprintf("password with URL %q already exists", e.url)
}

type ErrPasswordCardNotFound struct {
	id string
}

// Error implements error type interface.
func (e ErrPasswordCardNotFound) Error() string {
	return fmt.Sprintf("password with ID %q not found", e.id)
}

func (pr *PasswordCardRepository) Insert(newPasswordCard model.PasswordCard) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	for _, passwordCard := range pr.passwordCards {
		if passwordCard.ID == newPasswordCard.ID {
			return ErrPasswordCardAlreadyExists{id: newPasswordCard.ID}
		}
		if passwordCard.URL == newPasswordCard.URL {
			return ErrPasswordCardAlreadyExists{url: newPasswordCard.URL}
		}
	}

	pr.passwordCards = append(pr.passwordCards, newPasswordCard)

	return nil
}

func (pr *PasswordCardRepository) Update(updatedPasswordCard model.PasswordCard) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	for i, passwordCard := range pr.passwordCards {
		// verify if the updated URL already exists for other cards
		if passwordCard.ID != updatedPasswordCard.ID && passwordCard.URL == updatedPasswordCard.URL {
			return ErrPasswordCardAlreadyExists{url: updatedPasswordCard.URL}
		}

		if passwordCard.ID == updatedPasswordCard.ID {
			pr.passwordCards[i] = updatedPasswordCard
			return nil
		}
	}

	return ErrPasswordCardNotFound{id: updatedPasswordCard.ID}
}

func (pr *PasswordCardRepository) Delete(passwordCardID string) error {
	pr.mu.Lock()
	defer pr.mu.Unlock()

	for i, passwordCard := range pr.passwordCards {
		if passwordCard.ID == passwordCardID {
			pr.passwordCards = append(pr.passwordCards[:i], pr.passwordCards[i+1:]...)
			return nil
		}
	}

	return ErrPasswordCardNotFound{id: passwordCardID}
}

func (pr *PasswordCardRepository) GetAll() []model.PasswordCard {
	return pr.passwordCards
}
