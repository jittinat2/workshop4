package http

import (
	"net/http"
	"workshop4/internal/entity"
	"workshop4/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	u usecase.UserUsecase
}

func New(u usecase.UserUsecase) *Handler { return &Handler{u: u} }

func (h *Handler) Register(r fiber.Router) {
	r.Get("/", h.list)
	r.Post("/", h.create)
	r.Get(":id", h.get)
	r.Put(":id", h.update)
	r.Delete(":id", h.delete)
}

func (h *Handler) list(c *fiber.Ctx) error {
	users, err := h.u.List()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *Handler) create(c *fiber.Ctx) error {
	var in entity.User
	if err := c.BodyParser(&in); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := h.u.Create(&in); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(in)
}

func (h *Handler) get(c *fiber.Ctx) error {
	id := c.Params("id")
	u, err := h.u.Get(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(u)
}

func (h *Handler) update(c *fiber.Ctx) error {
	id := c.Params("id")
	var in entity.User
	if err := c.BodyParser(&in); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	in.ID = id
	if err := h.u.Update(&in); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(in)
}

func (h *Handler) delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.u.Delete(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(http.StatusNoContent)
}
