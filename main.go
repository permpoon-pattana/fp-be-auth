package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/permpoon-pattana/authentication"
)

func main() {
	godotenv.Load()
	store := getSessionStore()
	auth, err := authentication.NewAuthenticator()
	if err != nil {
		log.Panic(err)
	}
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowCredentials: true,
	}))
	app.Get("/login", authentication.LoginHandler(&authentication.LoginParams{
		Auth:  auth,
		Store: store,
	}))
	app.Get("/callback", authentication.CallbackHandler(&authentication.CallbackParams{
		Auth:  auth,
		Store: store,
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("http://localhost:3000")
	})
	app.Get("/info", func(c *fiber.Ctx) error {
		sess, _ := sessStore.Get(c)
		sub := sess.Get(authentication.SESSION_HEADER_SUBJECT)
		return c.JSON(fiber.Map{
			"sub": sub,
		})
	})
	// app.Listen(os.Getenv("PORT"))
	app.Listen(":5000")
}
