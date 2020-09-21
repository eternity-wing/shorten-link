package linkservice

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type validationHandler struct {
	next newLinkHandler
}

func (v *validationHandler) setNext(next newLinkHandler) {
	v.next = next
}

func (v *validationHandler) execute(c *fiber.Ctx, req *request) error {
	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":  "Validation error",
			"title": "There was a validation error",
			"error": err.Error(),
		})
	}
	return v.next.execute(c, req)
}
