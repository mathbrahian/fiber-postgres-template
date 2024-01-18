package middleware

import (
	"github.com/gofiber/fiber/v2"
)

import (
	"fmt"
	"errors"
	"strings"
	"net/http"
	// "io/ioutil"
	"encoding/json"
)

import uuid "github.com/google/uuid"

type DeserializeUser struct {
	ID uuid.UUID  `json:"id"`
}

/* 
	REQUIRED(Any middleware must have this)

	For every middleware we need a config.
	In config we also need to define a function which allows us to skip the middleware if return true.
	By convention it should be named as "Filter" but any other name will work too.
*/
type Config struct {
	Filter         func(c *fiber.Ctx) bool // Required
	RequestAUTH		func(c *fiber.Ctx) (*DeserializeUser, error)
	Unauthorized fiber.Handler  // middleware specfic
}

/*
	Middleware specific

	Our middleware's config default values if not passed
*/
var ConfigDefault = Config{
	Filter:       nil,
	RequestAUTH: nil,
	Unauthorized: nil,
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

		cfg.RequestAUTH = func(c *fiber.Ctx) (*DeserializeUser, error){

			authHeader := c.Get("Authorization")

			if authHeader == "" {
				return nil, errors.New("Authorization header is required")
			}

			bearerTokenList := strings.SplitN(authHeader, " ", -1)

			if len(bearerTokenList) != 2 && bearerTokenList[0] !="Bearer"{
				return nil, errors.New("Error parsing token")
			}

			req, err := http.NewRequest(http.MethodGet, "http://host.docker.internal:6050/user/api/auth", nil)
			if err != nil {
				return nil, errors.New("Could not create request")
			}			
			
		
			bearerToken := fmt.Sprintf("Bearer %s", bearerTokenList[1])
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", bearerToken)

			res, err := http.DefaultClient.Do(req)
			
			if err != nil {
				return nil, errors.New("Error making http request")
			}

			defer res.Body.Close()

			var user DeserializeUser

			err = json.NewDecoder(res.Body).Decode(&user)
			

			if err != nil {
				return nil, errors.New("Could not deserialezer response body")
			}

		
			return &user, nil
		}
	}

	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
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

		requestAuth, err := cfg.RequestAUTH(c)


		if err == nil {
			c.Locals("requestAuth", *requestAuth)
			return c.Next()
		}
	
		return cfg.Unauthorized(c)
	}
}
