package service

import (
	"testing"

	"github.com/CaioTeixeira95/password-manager/backend/model"
	"github.com/CaioTeixeira95/password-manager/backend/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreatePasswordCard(t *testing.T) {
	r := repository.NewPasswordCardRepository()
	s := NewPasswordCardService(r)

	pc, err := s.CreatePasswordCard(model.PasswordCard{
		ID:       "card-id-1",
		Name:     "AWS",
		Username: "username",
		Password: "supersecret",
		URL:      "https://aws.com/login",
	})
	require.NoError(t, err)
	assert.NotNil(t, pc)

	pc, err = s.CreatePasswordCard(model.PasswordCard{
		ID:       "card-id-1",
		Name:     "AWS",
		Username: "username",
		Password: "supersecret",
		URL:      "https://aws.com/login",
	})

	assert.EqualError(t, err, `error creating a new password card: password with ID "card-id-1" already exists`)
	assert.Nil(t, pc)
}

func TestListPasswordCards(t *testing.T) {
	r := repository.NewPasswordCardRepository()
	s := NewPasswordCardService(r)

	assert.Empty(t, s.ListPasswordCards())

	r = repository.CustomPasswordCardRepository([]model.PasswordCard{
		{
			ID:       "card-id-1",
			Name:     "AWS",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		},
		{
			ID:       "card-id-2",
			Name:     "GCP",
			Username: "username",
			Password: "supersecret",
			URL:      "https://cloud.google.com/login",
		},
	})
	s = NewPasswordCardService(r)

	assert.Equal(t, []model.PasswordCard{
		{
			ID:       "card-id-1",
			Name:     "AWS",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		},
		{
			ID:       "card-id-2",
			Name:     "GCP",
			Username: "username",
			Password: "supersecret",
			URL:      "https://cloud.google.com/login",
		},
	}, s.ListPasswordCards())
}

func TestUpdatePasswordCardService(t *testing.T) {
	r := repository.CustomPasswordCardRepository([]model.PasswordCard{
		{
			ID:       "card-id-1",
			Name:     "AWS",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		},
	})
	s := NewPasswordCardService(r)

	pc, err := s.UpdatePasswordCard(model.PasswordCard{
		ID:       "card-id-1",
		Name:     "Amazon Web Services",
		Username: "username",
		Password: "newsupersecret",
		URL:      "https://aws.com/login",
	})
	require.NoError(t, err)
	assert.NotNil(t, pc)

	pc, err = s.UpdatePasswordCard(model.PasswordCard{
		ID:       "card-id-2",
		Name:     "AWS",
		Username: "username",
		Password: "supersecret",
		URL:      "https://aws.com/login",
	})

	assert.EqualError(t, err, `error updating password card: password with URL "https://aws.com/login" already exists`)
	assert.Nil(t, pc)

	pc, err = s.UpdatePasswordCard(model.PasswordCard{
		ID:       "card-id-2",
		Name:     "AWS",
		Username: "username",
		Password: "supersecret",
		URL:      "https://another.aws.com/login",
	})

	assert.EqualError(t, err, `error updating password card: password with ID "card-id-2" not found`)
	assert.Nil(t, pc)
}

func TestDeletePasswordCard(t *testing.T) {
	r := repository.CustomPasswordCardRepository([]model.PasswordCard{
		{
			ID:       "card-id-1",
			Name:     "AWS",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		},
		{
			ID:       "card-id-2",
			Name:     "GCP",
			Username: "username",
			Password: "supersecret",
			URL:      "https://cloud.google.com/login",
		},
	})
	s := NewPasswordCardService(r)

	err := s.DeletePasswordCard("card-id-1")
	require.NoError(t, err)

	err = s.DeletePasswordCard("card-id-1")
	assert.EqualError(t, err, `error deleting password card: password with ID "card-id-1" not found`)
}
