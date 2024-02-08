package facades

import (
	"log"

	"github.com/go-unity/framework/contracts/route"

	"github.com/go-unity/fiber"
)

func Route(driver string) route.Route {
	instance, err := fiber.App.MakeWith(fiber.RouteBinding, map[string]any{
		"driver": driver,
	})
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return instance.(*fiber.Route)
}
