package controller

import (
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2"
	cek "github.com/nekowawolf/create-db-go"
)

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetTugasakhirData(c *fiber.Ctx) error {
	ps := cek.GetAllTugas()
	return c.JSON(ps)
}
