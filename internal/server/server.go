package server

import (
	"github.com/gofiber/fiber/v2"

	"JusticeNet/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "JusticeNet",
			AppName:      "JusticeNet",
		}),

		db: database.New(),
	}

	return server
}
