package client

import "github.com/gofiber/fiber/v2"

func GetRequest(uri string) *fiber.Agent {
	a := fiber.AcquireAgent()
	req := a.Request()
	req.Header.SetMethod(fiber.MethodGet)
	req.SetRequestURI(uri)

	if err := a.Parse(); err != nil {
		panic(err)
	}

	return a
}
