package linkcontroller

import (
	"github.com/asaskevich/govalidator"
	"github.com/eternity-wing/short_link/lib/baseconversion/base62"
	"github.com/eternity-wing/short_link/repository/linkrepository"
	"github.com/eternity-wing/short_link/service/linkservice"
	"github.com/eternity-wing/short_link/service/shortlinkconversion"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func New(c *fiber.Ctx) error {
	return linkservice.ProcessNewLinkRequest(c)
}

func Get(c *fiber.Ctx) error {

	shorten := c.Params("shorten")
	if govalidator.IsAlphanumeric(shorten) != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"type":  "Validation error",
			"title": "There was a validation error",
			"error": "Url contains unallowable character(accepted characters are alphanumerical)",
		})
	}

	convertor := shortlinkconversion.InitConvertor(base62.Convertor{})

	filter := bson.M{"id": convertor.Decode(shorten)}
	lnkRepo := linkrepository.NewRepository()
	link := lnkRepo.Find(filter)
	if link == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "The requested resource does not exist"})
	}
	return c.Redirect(link.URL, fiber.StatusMovedPermanently)
}
