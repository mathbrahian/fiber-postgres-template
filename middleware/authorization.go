package middleware

import (
	"github.com/gofiber/fiber/v2"
)

import (
	"fmt"
	"errors"
	"strings"
	"net/http"
	"io/ioutil"
)

/* 
	REQUIRED(Any middleware must have this)

	For every middleware we need a config.
	In config we also need to define a function which allows us to skip the middleware if return true.
	By convention it should be named as "Filter" but any other name will work too.
*/
type Config struct {
	Filter         func(c *fiber.Ctx) bool // Required
	RequestAUTH		func(c *fiber.Ctx) (*string, error)
	// Unauthorized fiber.Handler  // middleware specfic
}

/*
	Middleware specific

	Our middleware's config default values if not passed
*/
var ConfigDefault = Config{
	Filter:       nil,
	RequestAUTH: nil,
	// Unauthorized: nil,
}

/*
	Middleware specific

	Function for generating default config
*/
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values if not passed
	if cfg.Filter == nil {
		cfg.Filter = ConfigDefault.Filter
	}

	if cfg.RequestAUTH == nil {
		cfg.RequestAUTH = func(c *fiber.Ctx) (*string, error){

			authHeader := c.Get("Authorization")

			if authHeader == "" {
				return nil, errors.New("Authorization header is required")
			}

			bearerToken := strings.SplitN(authHeader, " ", -1)


			if len(bearerToken) == 2 && bearerToken[0]=="Bearer"{
				token := bearerToken[1]
				req, err := http.NewRequest(http.MethodGet, "http://host.docker.internal:6050/user/api/auth", nil)
				if err != nil {
					fmt.Printf("client: error making http request: %s\n", err)
				}
				req.Header.Set("Content-Type", "application/json")
				bearerToken2 := fmt.Sprintf("Bearer %s", token)
				fmt.Printf(bearerToken2)
				req.Header.Set("Authorization", bearerToken2)			
				res, err := http.DefaultClient.Do(req)
				if err != nil {
					fmt.Printf("client: error making http request: %s\n", err)
				}
				body, err := ioutil.ReadAll(res.Body)
				bodyString := string(body)
				return &bodyString, nil
			}

			return nil, errors.New("Authorization header is required")

		}
	}

	return cfg
}

/*
    REQUIRED(Any middleware must have this)

	Our main middleware function used to initialize our middleware.
	By convention we name it "New" but any other name will work too.
*/
func New(config Config) fiber.Handler {

	// For setting default config
	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {
		// Don't execute middleware if Filter returns true
		if cfg.Filter != nil && cfg.Filter(c) {
			fmt.Println("Midddle was skipped")
			return c.Next()
		}
		fmt.Println("Midddle was run")

		requestAuth, err := cfg.RequestAUTH(c)

		if err == nil {
			c.Locals("requestAuth", *requestAuth)
			return c.Next()
		}
		fmt.Println("Midddle error")

		return c.Next()
		// return cfg.Unauthorized(c)
	}
}
