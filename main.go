package main

import (
	"context"
	"log"

	"portal/frontend"
	"portal/utils"

	"github.com/a-h/templ"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

var ctx = context.Background()

func main() {
	cfg := utils.LoadConfig()
  executor, c, d := utils.ConnectToTurso(cfg)
  defer c.Close()
  defer d.Close()

  _, err := executor.GetApp(ctx, 1)
  if err != nil {
    log.Println(err.Error())
  }

	app := fiber.New(fiber.Config{Prefork: cfg.Prefork})

  app.Static("/assets", "./assets")

  app.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
  // Websocket for live reload
  app.Get("/ws", websocket.New(func (c *websocket.Conn) {
    var (
			mt  int
			msg []byte
			err error
		)
    err = c.WriteMessage(websocket.TextMessage, []byte("2000"))
    if err != nil {
      log.Println(err)
    }
    for {
      if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("write:", err)
				break
			}
    }
  }))
  
	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, frontend.Sign())
	})
  app.Post("/sign/in", func (c *fiber.Ctx) error {
    // email := c.FormValue("email")
    // password := c.FormValue("password")
    return c.SendString("That user doesn't exist")
  })
	panic(app.Listen(cfg.Address))
}

func Render(c *fiber.Ctx, component templ.Component) error {
	c.Set("Content-Type", "text/html")
	return component.Render(c.Context(), c.Response().BodyWriter())
}
