package main

import (
	"fmt"
	"log"
	"os"

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

	port := fmt.Sprintf(":%v", os.Getenv("PORT"))
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CORS_ORIGINS"),
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		redirectURL := os.Getenv("LOGIN_SUCCESS_REDIRECT_URL")
		if redirectURL == "" {
			redirectURL = DefaultRedirectURL
		}

		return c.Redirect(redirectURL)
	})
	app.Get("/login", authentication.LoginHandler(&authentication.LoginParams{
		Auth:  auth,
		Store: store,
	}))
	app.Get("/callback", authentication.CallbackHandler(&authentication.CallbackParams{
		Auth:  auth,
		Store: store,
	}))
	app.Get("/info", func(c *fiber.Ctx) error {
		sess, _ := sessStore.Get(c)
		sub := sess.Get(authentication.SESSION_HEADER_SUBJECT)
		return c.JSON(fiber.Map{
			SubjectKey: sub,
		})
	})

	app.Listen(port)
}
