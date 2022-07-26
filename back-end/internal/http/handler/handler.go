package handler

import (
	"github.com/amirhnajafiz/personal-website/back-end/internal/database/store"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	ProjectsCollection store.ProjectsCollection
}

func (h *Handler) RegisterClient(api fiber.Router) {
	api.Get("/projects", h.GetVisibleProjects)
	api.Get("/project/:id", h.GetProjectById)
}

func (h *Handler) RegisterAdmin(api fiber.Router) {
	api.Get("/projects", h.GetAllProjects)
	api.Put("/project/:id", h.UpsertProject)
	api.Delete("/project/:id", h.RemoveProject)
}
