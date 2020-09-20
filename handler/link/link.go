package link

import (
	"github.com/asaskevich/govalidator"
	"github.com/eternity-wing/short_link/handler"
	"github.com/eternity-wing/short_link/lib/baseconversion/base62"
	"github.com/eternity-wing/short_link/lib/shortlinkconversion"
	"github.com/eternity-wing/short_link/model"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type request struct {
	URL    string `valid:"url"`
	UserID int    `valid:"type(int),optional"`
}

type response struct {
	URL      string `json:"url"`
	ShortUrl string `json:"short_url"`
}

func GetLink(c *fiber.Ctx) error {

	shorten := c.Params("shorten")

	convertor := shortlinkconversion.InitConvertor(base62.Convertor{})
	ID, _ := convertor.Decode(shorten)

	filter := bson.M{"id": ID}
	link, err := model.GetLink(filter)
	if err != nil || link == nil {

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}
	return c.Redirect(link.URL, fiber.StatusMovedPermanently)
}

func NewLink(c *fiber.Ctx) error {

	req := new(request)

	if err := c.BodyParser(&req); err != nil {
		return handler.SendInvalidJSONRequestResponse(c)
	}

	_, err := govalidator.ValidateStruct(req)
	if err != nil {
		return handler.SendValidationFailedResponse(c, err)
	}

	link := model.Link{
		UserID: req.UserID,
		URL:    req.URL,
	}

	_, err = model.NewLink(&link)

	if err != nil {
		return handler.SendInternalServerErrorResponse(c, err)
	}
	convertor := shortlinkconversion.InitConvertor(base62.Convertor{})
	shorten, _ := convertor.GetShorten(link.ID)


	return c.Status(fiber.StatusOK).JSON(response{
		URL:      link.URL,
		ShortUrl: shorten,
	})
}
