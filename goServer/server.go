package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func getPort() string {
	const defaultPort string = "9000"
	var port string = defaultPort

	args := os.Args[1:]
	if len(args) > 0 {
		portArg := args[0]

		if portArg != "" {
			port = portArg
		}
	} else {
		log.Println("Using default port", defaultPort)
	}

	return port
}

func main() {
	port := ":" + getPort()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		time := time.Now().Unix()
		str := strconv.FormatInt(time, 10)
		return c.SendString(str)
	})

	log.Fatal(app.Listen(port))
}
