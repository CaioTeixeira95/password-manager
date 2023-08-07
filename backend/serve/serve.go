package serve

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/CaioTeixeira95/password-manager/backend/model"
	"github.com/CaioTeixeira95/password-manager/backend/repository"
	"github.com/CaioTeixeira95/password-manager/backend/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

type Serve struct {
	app                 *fiber.App
	passwordCardService *service.PasswordCardService
}

func NewServe(app *fiber.App, passwordCardService *service.PasswordCardService) *Serve {
	return &Serve{
		app:                 app,
		passwordCardService: passwordCardService,
	}
}

func (s *Serve) Run(port int) error {
	s.initHandlers()
	if err := s.app.Listen(fmt.Sprintf(":%d", port)); err != nil {
		return fmt.Errorf("error starting serve: %w", err)
	}

	return nil
}

func (s *Serve) initHandlers() {
	s.app.Use(recover.New())
	s.app.Use(logger.New())
	s.app.Use(cors.New())

	s.app.Route("/password-cards", func(router fiber.Router) {
		router.Get("/", handleGetPasswordCards(s.passwordCardService))
		router.Post("/", handlePostPasswordCards(s.passwordCardService))

		router.Route("/:id", func(router fiber.Router) {
			router.Put("/", handlePutPasswordCards(s.passwordCardService))
			router.Delete("/", handleDeletePasswordCards(s.passwordCardService))
		})
	})
}

func handleGetPasswordCards(s *service.PasswordCardService) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.JSON(s.ListPasswordCards())
	}
}

func handlePostPasswordCards(s *service.PasswordCardService) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var passwordCardRequest model.PasswordCard
		if err := c.BodyParser(&passwordCardRequest); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "The request is invalid in some way.",
				Error:   err.Error(),
			})
		}

		if err := passwordCardRequest.Validate(); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Validation error.",
				Error:   err.Error(),
			})
		}

		_, err := s.CreatePasswordCard(passwordCardRequest)
		if err != nil {
			log.Printf("error creating password card: %s", err.Error())

			var errExists repository.ErrPasswordCardAlreadyExists
			if errors.As(err, &errExists) {
				return c.Status(http.StatusConflict).JSON(ErrorResponse{
					Status:  http.StatusConflict,
					Message: "Conflict.",
					Error:   errExists.Error(),
				})
			}

			return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal Server Error.",
			})
		}

		return c.Status(http.StatusCreated).JSON(passwordCardRequest)
	}
}

func handlePutPasswordCards(s *service.PasswordCardService) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		passwordCardID := c.Params("id")

		var passwordCardRequest model.PasswordCard
		if err := c.BodyParser(&passwordCardRequest); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "The request is invalid in some way.",
				Error:   err.Error(),
			})
		}

		passwordCardRequest.ID = passwordCardID
		if err := passwordCardRequest.Validate(); err != nil {
			return c.Status(http.StatusBadRequest).JSON(ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "Validation error.",
				Error:   err.Error(),
			})
		}

		_, err := s.UpdatePasswordCard(passwordCardRequest)
		if err != nil {
			log.Printf("error creating password card: %s", err.Error())

			var errExists repository.ErrPasswordCardAlreadyExists
			if errors.As(err, &errExists) {
				return c.Status(http.StatusConflict).JSON(ErrorResponse{
					Status:  http.StatusConflict,
					Message: "Conflict.",
					Error:   errExists.Error(),
				})
			}

			var errNotFound repository.ErrPasswordCardNotFound
			if errors.As(err, &errNotFound) {
				return c.Status(http.StatusNotFound).JSON(ErrorResponse{
					Status:  http.StatusNotFound,
					Message: "Password Card not found.",
					Error:   errNotFound.Error(),
				})
			}

			return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal Server Error.",
			})
		}

		return c.JSON(passwordCardRequest)
	}
}

func handleDeletePasswordCards(s *service.PasswordCardService) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		passwordCardID := c.Params("id")

		err := s.DeletePasswordCard(passwordCardID)
		if err != nil {
			log.Printf("error deleting password card: %s", err.Error())

			var errNotFound repository.ErrPasswordCardNotFound
			if errors.As(err, &errNotFound) {
				return c.Status(http.StatusNotFound).JSON(ErrorResponse{
					Status:  http.StatusNotFound,
					Message: "Password Card not found.",
					Error:   errNotFound.Error(),
				})
			}

			return c.Status(http.StatusInternalServerError).JSON(ErrorResponse{
				Status:  http.StatusInternalServerError,
				Message: "Internal Server Error.",
			})
		}

		return c.SendStatus(http.StatusNoContent)
	}
}
