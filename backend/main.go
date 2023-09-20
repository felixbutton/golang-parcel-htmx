package main

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
}

func main() {
	dirname := getPathToRoot()
	engine := html.New("./public/views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	engine.Reload(true)

	fmt.Println(dirname + "/public")

	app.Use(logger.New())
	app.Use(recover.New())
	app.Static("/assets", dirname+"/public", fiber.Static{
		CacheDuration: 30 * 24 * time.Hour,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/index", fiber.Map{
			"name": "unkwown",
		})
	})

	app.Post("hello", func(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		return c.Render("actions/hello", fiber.Map{
			"name": p.Name,
		})
	})

	app.Post("clicked", func(c *fiber.Ctx) error {
		return c.Render("actions/clicked", fiber.Map{})
	})

	log.Fatal(app.Listen(":8080"))
}

func getPathToRoot() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("get directory failed")
	}

	return filepath.Dir(filename)
}
