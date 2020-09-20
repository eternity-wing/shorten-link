package handler

import (
	"github.com/gofiber/fiber/v2"
)

type errorResponse struct {
	status int
	errType string
	title string
	err string
}

func SendInvalidJSONRequestResponse(c *fiber.Ctx) error {
	return sendErrorResponse(c, errorResponse{
		status: fiber.StatusBadRequest,
		errType:  "Invalid request",
		title: "Invalid request format",
		err: "cannot parse JSON",
	})
}


func SendValidationFailedResponse(c *fiber.Ctx, err error) error {

	return sendErrorResponse(c, errorResponse{
		status: fiber.StatusBadRequest,
		errType:  "Validation error",
		title: "There was a validation error",
		err: err.Error(),
	})

}


func SendInternalServerErrorResponse(c *fiber.Ctx, err error) error {

	return sendErrorResponse(c, errorResponse{
		status: fiber.StatusInternalServerError,
		errType:  "Server Error",
		title: "Internal server error occurred",
		err: err.Error(),
	})
}

func sendErrorResponse(c *fiber.Ctx, res errorResponse) error  {
	return c.Status(res.status).JSON(fiber.Map{
		"type":  res.errType,
		"title": res.title,
		"error": res.err,
	})
}