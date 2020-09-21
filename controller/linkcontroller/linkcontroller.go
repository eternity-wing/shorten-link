package linkcontroller

import (
	"github.com/eternity-wing/short_link/lib/baseconversion/base62"
	"github.com/eternity-wing/short_link/repository/linkrepository"
	"github.com/eternity-wing/short_link/service/linkservice"
	"github.com/eternity-wing/short_link/service/shortlinkconversion"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func Get(c *fiber.Ctx) error {

	shorten := c.Params("shorten")

	convertor := shortlinkconversion.InitConvertor(base62.Convertor{})

	filter := bson.M{"id": convertor.Decode(shorten)}
	lnkRepo := linkrepository.NewRepository()
	link := lnkRepo.Find(filter)
	if link == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "The requested resource does not exist"})
	}
	return c.Redirect(link.URL, fiber.StatusMovedPermanently)
}

func New(c *fiber.Ctx) error {
	return linkservice.ProcessNewLinkRequest(c)
}
