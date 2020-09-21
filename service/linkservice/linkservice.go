package linkservice

import (
	"github.com/eternity-wing/short_link/lib/baseconversion/base62"
	"github.com/eternity-wing/short_link/model"
	"github.com/eternity-wing/short_link/repository/counterrepository"
	"github.com/eternity-wing/short_link/repository/linkrepository"
	"github.com/eternity-wing/short_link/service/shortlinkconversion"
	"github.com/gofiber/fiber/v2"
)

type request struct {
	URL    string `valid:"url"`
	UserID int    `valid:"type(int),optional"`
}

type response struct {
	URL      string `json:"url"`
	ShortUrl string `json:"short_url"`
}

type reqHandler interface {
	execute(c *fiber.Ctx, req *request) error
	setNext(reqHandler)
}

type newLinkHandler struct {
	ctrRepo counterrepository.Repository
	lnkRepo linkrepository.Repository
}

func (v *newLinkHandler) execute(c *fiber.Ctx, req *request) error {
	link := model.Link{
		UserID: req.UserID,
		URL:    req.URL,
		ID:     v.ctrRepo.GetNextSequenceValue("link"),
	}

	v.lnkRepo.Create(&link)
	convertor := shortlinkconversion.InitConvertor(base62.Convertor{})

	return c.Status(fiber.StatusOK).JSON(response{
		URL:      link.URL,
		ShortUrl: convertor.GetShorten(link.ID),
	})
}

func ProcessNewLinkRequest(c *fiber.Ctx) error {
	req := new(request)

	lh := newLinkHandler{
		ctrRepo: *counterrepository.NewRepository(),
		lnkRepo: *linkrepository.NewRepository(),
	}
	vh := validationHandler{next: lh}
	dto := dataTransferObjectHandler{next: vh}
	return dto.execute(c, req)
}
