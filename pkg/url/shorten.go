package url

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type Request struct {
	URL     string        `json:"url"`
	Short   string        `json:"short"`
	Expires time.Duration `json:"expires"`
}

type Response struct {
	URL            string        `json:"url"`
	Short          string        `json:"short"`
	Expires        time.Duration `json:"expires"`
	XRateLimit     int           `json:"rate_limit"`
	XRateLimitRest time.Duration `json:"rate_limit_reset"`
}

func Shorten(ctx *fiber.Ctx) error {
	body := &Request{}

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "JSON parse error"})
	}

	//TODO: implement shorten

	res := Response{
		URL:     body.URL,
		Short:   "",
		Expires: body.Expires,
	}

	return ctx.Status(fiber.StatusOK).JSON(res)
}
