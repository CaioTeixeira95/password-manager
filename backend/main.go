package main

import (
	"flag"
	"log"

	"github.com/CaioTeixeira95/password-manager/backend/repository"
	"github.com/CaioTeixeira95/password-manager/backend/serve"
	"github.com/CaioTeixeira95/password-manager/backend/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := flag.Int("port", 8000, "Web server port")

	flag.Parse()

	s := serve.NewServe(
		fiber.New(),
		service.NewPasswordCardService(
			repository.NewPasswordCardRepository(),
		),
	)

	if err := s.Run(*port); err != nil {
		log.Fatal(err)
	}
}
