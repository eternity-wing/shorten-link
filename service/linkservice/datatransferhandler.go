package linkservice

import (
	"github.com/gofiber/fiber/v2"
)

type dataTransferObjectHandler struct {
	next validationHandler
}

func (d *dataTransferObjectHandler) setNext(next validationHandler) {
	d.next = next
}

func (d *dataTransferObjectHandler) execute(c *fiber.Ctx, req *request) error {

	if err := c.BodyParser(&req); err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":  "Invalid request",
			"title": "Invalid request format",
			"error": "cannot parse JSON",
		})
	}
	return d.next.execute(c, req)
}
