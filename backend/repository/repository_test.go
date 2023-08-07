package repository

import (
	"sync"
	"testing"

	"github.com/CaioTeixeira95/password-manager/backend/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPasswordCardRepositoryInsert(t *testing.T) {
	r := CustomPasswordCardRepository([]model.PasswordCard{
		{
			ID:       "card-id-1",
			Name:     "AWS",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		},
	})

	t.Run("returns error when try to insert a new password card with an existant ID", func(t *testing.T) {
		err := r.Insert(model.PasswordCard{
			ID:       "card-id-1",
			Name:     "AWS",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrPasswordCardAlreadyExists{id: "card-id-1"})
	})

	t.Run("returns error when tries to insert a new password card with an existant URL", func(t *testing.T) {
		err := r.Insert(model.PasswordCard{
			ID:       "card-id-2",
			Name:     "AWS",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrPasswordCardAlreadyExists{url: "https://aws.com/login"})
	})

	t.Run("🎉 inserts a new password card successfully", func(t *testing.T) {
		err := r.Insert(model.PasswordCard{
			ID:       "card-id-2",
			Name:     "Google Cloud Platform",
			Username: "username",
			Password: "supersecret",
			URL:      "https://cloud.google.com/",
		})
		require.NoError(t, err)

		assert.Len(t, r.passwordCards, 2)
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
				Name:     "Google Cloud Platform",
				Username: "username",
				Password: "supersecret",
				URL:      "https://cloud.google.com/",
			},
		}, r.passwordCards)
	})

	t.Run("ensures no race condition", func(t *testing.T) {
		pr := NewPasswordCardRepository()

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			err := pr.Insert(model.PasswordCard{
				ID:       "card-id-1",
				Name:     "AWS",
				Username: "username",
				Password: "supersecret",
				URL:      "https://aws.com/login",
			})
			require.NoError(t, err)
		}()

		go func() {
			defer wg.Done()
			err := pr.Insert(model.PasswordCard{
				ID:       "card-id-2",
				Name:     "Google Cloud Platform",
				Username: "username",
				Password: "supersecret",
				URL:      "https://cloud.google.com/",
			})
			require.NoError(t, err)
		}()

		wg.Wait()

		assert.Len(t, pr.passwordCards, 2)
		assert.Contains(t, pr.passwordCards, model.PasswordCard{
			ID:       "card-id-1",
			Name:     "AWS",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		})
		assert.Contains(t, pr.passwordCards, model.PasswordCard{
			ID:       "card-id-2",
			Name:     "Google Cloud Platform",
			Username: "username",
			Password: "supersecret",
			URL:      "https://cloud.google.com/",
		})
	})
}

func TestPasswordCardRepositoryUpdate(t *testing.T) {
	r := CustomPasswordCardRepository([]model.PasswordCard{
		{
			ID:       "card-id-1",
			Name:     "AWS",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		},
		{
			ID:       "card-id-2",
			Name:     "Google Cloud Platform",
			Username: "username",
			Password: "supersecret",
			URL:      "https://cloud.google.com/",
		},
	})

	t.Run("returns error when tries to update a password card with an existant URL", func(t *testing.T) {
		err := r.Update(model.PasswordCard{
			ID:       "card-id-2",
			Name:     "Google Cloud Platform",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrPasswordCardAlreadyExists{url: "https://aws.com/login"})
	})

	t.Run("returns error when password card is not found", func(t *testing.T) {
		err := r.Update(model.PasswordCard{
			ID:       "card-id-3",
			Name:     "Google Cloud Platform",
			Username: "username",
			Password: "supersecret",
			URL:      "https://another.google.com/login",
		})

		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrPasswordCardNotFound{id: "card-id-3"})
	})

	t.Run("🎉 updates a password card successfully", func(t *testing.T) {
		err := r.Update(model.PasswordCard{
			ID:       "card-id-2",
			Name:     "Google Cloud Platform - GCP",
			Username: "username",
			Password: "supersecret",
			URL:      "https://another.google.com/login",
		})
		require.NoError(t, err)

		assert.Len(t, r.passwordCards, 2)
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
				Name:     "Google Cloud Platform - GCP",
				Username: "username",
				Password: "supersecret",
				URL:      "https://another.google.com/login",
			},
		}, r.passwordCards)
	})
}

func TestPasswordCardRepositoryDelete(t *testing.T) {
	r := CustomPasswordCardRepository([]model.PasswordCard{
		{
			ID:       "card-id-1",
			Name:     "AWS",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		},
		{
			ID:       "card-id-2",
			Name:     "Google Cloud Platform",
			Username: "username",
			Password: "supersecret",
			URL:      "https://cloud.google.com/",
		},
	})

	t.Run("returns error when password card is not found", func(t *testing.T) {
		err := r.Delete("card-id-3")
		assert.Error(t, err)
		assert.ErrorIs(t, err, ErrPasswordCardNotFound{id: "card-id-3"})
	})

	t.Run("🎉 deletes a password card successfully", func(t *testing.T) {
		err := r.Delete("card-id-1")
		require.NoError(t, err)

		assert.Len(t, r.passwordCards, 1)
		assert.Equal(t, []model.PasswordCard{
			{
				ID:       "card-id-2",
				Name:     "Google Cloud Platform",
				Username: "username",
				Password: "supersecret",
				URL:      "https://cloud.google.com/",
			},
		}, r.passwordCards)
	})

	t.Run("ensures no race condition", func(t *testing.T) {
		pr := CustomPasswordCardRepository([]model.PasswordCard{
			{
				ID:       "card-id-1",
				Name:     "AWS",
				Username: "username",
				Password: "supersecret",
				URL:      "https://aws.com/login",
			},
			{
				ID:       "card-id-2",
				Name:     "Google Cloud Platform",
				Username: "username",
				Password: "supersecret",
				URL:      "https://cloud.google.com/",
			},
		})

		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			err := pr.Delete("card-id-2")
			require.NoError(t, err)
		}()

		go func() {
			defer wg.Done()
			err := pr.Delete("card-id-1")
			require.NoError(t, err)
		}()

		wg.Wait()

		assert.Empty(t, pr.passwordCards)
	})
}

func TestGetAll(t *testing.T) {
	r := NewPasswordCardRepository()

	assert.Empty(t, r.GetAll())

	err := r.Insert(model.PasswordCard{
		ID:       "card-id-1",
		Name:     "AWS",
		Username: "username",
		Password: "supersecret",
		URL:      "https://aws.com/login",
	})
	require.NoError(t, err)

	err = r.Insert(model.PasswordCard{
		ID:       "card-id-2",
		Name:     "Google Cloud Platform",
		Username: "username",
		Password: "supersecret",
		URL:      "https://cloud.google.com/",
	})
	require.NoError(t, err)

	assert.Len(t, r.passwordCards, 2)
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
			Name:     "Google Cloud Platform",
			Username: "username",
			Password: "supersecret",
			URL:      "https://cloud.google.com/",
		},
	}, r.passwordCards)
}
