package link

import (
	"github.com/asaskevich/govalidator"
	"github.com/eternity-wing/short_link/handler"
	"github.com/eternity-wing/short_link/lib/baseconversion/base62"
	"github.com/eternity-wing/short_link/lib/shortlinkconversion"
	"github.com/eternity-wing/short_link/model"
	"github.com/eternity-wing/short_link/repository/counterrepository"
	"github.com/eternity-wing/short_link/repository/linkrepository"
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

	filter := bson.M{"id": convertor.Decode(shorten)}
	lnkRepo := linkrepository.NewRepository()
	link := lnkRepo.Find(filter)
	if link == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "The requested resource does not exist"})
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

	ctrRepo := counterrepository.NewRepository()
	lnkRepo := linkrepository.NewRepository()
	link := model.Link{
		UserID: req.UserID,
		URL:    req.URL,
		ID:     ctrRepo.GetNextSequenceValue("link"),
	}

	lnkRepo.Create(&link)
	convertor := shortlinkconversion.InitConvertor(base62.Convertor{})

	return c.Status(fiber.StatusOK).JSON(response{
		URL:      link.URL,
		ShortUrl: convertor.GetShorten(link.ID),
	})
}
