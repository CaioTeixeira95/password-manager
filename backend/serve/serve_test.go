package serve

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/CaioTeixeira95/password-manager/backend/model"
	"github.com/CaioTeixeira95/password-manager/backend/repository"
	"github.com/CaioTeixeira95/password-manager/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostPasswordCards(t *testing.T) {
	app := fiber.New()
	service := service.NewPasswordCardService(
		repository.CustomPasswordCardRepository([]model.PasswordCard{
			{
				ID:       "card-id-1",
				Name:     "AWS",
				Username: "username",
				Password: "supersecret",
				URL:      "https://aws.com/login",
			},
		}),
	)

	s := NewServe(app, service)
	s.initHandlers()

	url := "/password-cards"

	t.Run("return BadRequest for invalid body request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(`invalid`))
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.JSONEq(t, `{"error":"invalid character 'i' looking for beginning of value", "message":"The request is invalid in some way.", "status":400}`, string(respBody))

		// Validation error
		req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(`{"id": ""}`))
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp, err = app.Test(req)
		require.NoError(t, err)

		respBody, err = io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.JSONEq(t, `{"error":"invalid id", "message":"Validation error.", "status":400}`, string(respBody))
	})

	t.Run("return Conflict for duplicated entries", func(t *testing.T) {
		// duplicated ID
		reqBody := `
			{
				"id": "card-id-1",
				"name": "AWS",
				"username": "username",
				"password": "supersecret",
				"url": "https://aws.com/login"
			}
		`
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(reqBody))
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusConflict, resp.StatusCode)
		assert.JSONEq(t, `{"error":"password with ID \"card-id-1\" already exists", "message":"Conflict.", "status":409}`, string(respBody))

		// duplicated URL
		reqBody = `
			{
				"id": "card-id-2",
				"name": "AWS",
				"username": "username",
				"password": "supersecret",
				"url": "https://aws.com/login"
			}
		`
		req, err = http.NewRequest(http.MethodPost, url, strings.NewReader(reqBody))
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp, err = app.Test(req)
		require.NoError(t, err)

		respBody, err = io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusConflict, resp.StatusCode)
		assert.JSONEq(t, `{"error":"password with URL \"https://aws.com/login\" already exists", "message":"Conflict.", "status":409}`, string(respBody))
	})

	t.Run("creates a new password card successfully", func(t *testing.T) {
		reqBody := `
		{
			"id": "card-id-2",
			"name": "Google Cloud Platform",
			"username": "username",
			"password": "supersecret",
			"url": "https://cloud.google.com/login"
		}
	`
		req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(reqBody))
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		assert.JSONEq(t, reqBody, string(respBody))
	})
}

func TestPutPasswordCards(t *testing.T) {
	app := fiber.New()
	service := service.NewPasswordCardService(
		repository.CustomPasswordCardRepository([]model.PasswordCard{
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
				URL:      "https://cloud.google.com/login",
			},
		}),
	)

	s := NewServe(app, service)
	s.initHandlers()

	url := "/password-cards/%s"

	t.Run("return BadRequest for invalid body request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(url, "card-id-1"), strings.NewReader(`invalid`))
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.JSONEq(t, `{"error":"invalid character 'i' looking for beginning of value", "message":"The request is invalid in some way.", "status":400}`, string(respBody))

		// Validation error
		req, err = http.NewRequest(http.MethodPut, fmt.Sprintf(url, "card-id-1"), strings.NewReader(`{"name": ""}`))
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp, err = app.Test(req)
		require.NoError(t, err)

		respBody, err = io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.JSONEq(t, `{"error":"invalid name", "message":"Validation error.", "status":400}`, string(respBody))
	})

	t.Run("return Conflict for duplicated entries", func(t *testing.T) {
		reqBody := `
			{
				"name": "Google Cloud Platform",
				"username": "username",
				"password": "supersecret",
				"url": "https://aws.com/login"
			}
		`
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(url, "card-id-2"), strings.NewReader(reqBody))
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusConflict, resp.StatusCode)
		assert.JSONEq(t, `{"error":"password with URL \"https://aws.com/login\" already exists", "message":"Conflict.", "status":409}`, string(respBody))
	})

	t.Run("return NotFound when a non-existent is used", func(t *testing.T) {
		reqBody := `
			{
				"name": "Heroku",
				"username": "username",
				"password": "supersecret",
				"url": "https://heroku.com/login"
			}
		`
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(url, "card-id-3"), strings.NewReader(reqBody))
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.JSONEq(t, `{"error":"password with ID \"card-id-3\" not found", "message":"Password Card not found.", "status":404}`, string(respBody))
	})

	t.Run("updates a password card successfully", func(t *testing.T) {
		reqBody := `
			{
				"name": "Google Cloud Platform - GCP",
				"username": "username",
				"password": "mynewsupersecret",
				"url": "https://another.google.com/login"
			}
		`
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(url, "card-id-2"), strings.NewReader(reqBody))
		require.NoError(t, err)

		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)
		require.NoError(t, err)

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, `
			{
				"id": "card-id-2",
				"name": "Google Cloud Platform - GCP",
				"username": "username",
				"password": "mynewsupersecret",
				"url": "https://another.google.com/login"
			}
		`, string(respBody))
	})
}

func TestDeletePasswordCards(t *testing.T) {
	app := fiber.New()
	service := service.NewPasswordCardService(
		repository.CustomPasswordCardRepository([]model.PasswordCard{
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
				URL:      "https://cloud.google.com/login",
			},
		}),
	)

	s := NewServe(app, service)
	s.initHandlers()

	url := "/password-cards/%s"

	t.Run("return NotFound when a non-existent is used", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(url, "card-id-3"), nil)
		require.NoError(t, err)

		resp, err := app.Test(req)
		require.NoError(t, err)

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusNotFound, resp.StatusCode)
		assert.JSONEq(t, `{"error":"password with ID \"card-id-3\" not found", "message":"Password Card not found.", "status":404}`, string(respBody))
	})

	t.Run("deletes a password card successfully", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf(url, "card-id-2"), nil)
		require.NoError(t, err)

		resp, err := app.Test(req)
		require.NoError(t, err)

		assert.Equal(t, http.StatusNoContent, resp.StatusCode)
	})
}

func TestGetPasswordCards(t *testing.T) {
	app := fiber.New()
	r := repository.NewPasswordCardRepository()
	service := service.NewPasswordCardService(r)

	s := NewServe(app, service)
	s.initHandlers()

	url := "/password-cards"

	t.Run("gets all the password cards successfully", func(t *testing.T) {
		// with no password cards
		req, err := http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		resp, err := app.Test(req)
		require.NoError(t, err)

		respBody, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, `[]`, string(respBody))

		// with password cards
		r.Insert(model.PasswordCard{
			ID:       "card-id-1",
			Name:     "AWS",
			Username: "username",
			Password: "supersecret",
			URL:      "https://aws.com/login",
		})

		r.Insert(model.PasswordCard{
			ID:       "card-id-2",
			Name:     "GCP",
			Username: "username",
			Password: "supersecret",
			URL:      "https://cloud.google.com/login",
		})

		req, err = http.NewRequest(http.MethodGet, url, nil)
		require.NoError(t, err)

		resp, err = app.Test(req)
		require.NoError(t, err)

		respBody, err = io.ReadAll(resp.Body)
		resp.Body.Close()

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.JSONEq(t, `
			[
				{
					"id": "card-id-1",
					"name": "AWS",
					"username": "username",
					"password": "supersecret",
					"url": "https://aws.com/login"
				},
				{
					"id": "card-id-2",
					"name": "GCP",
					"username": "username",
					"password": "supersecret",
					"url": "https://cloud.google.com/login"
				}
			]
		`, string(respBody))
	})
}
