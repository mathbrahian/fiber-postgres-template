package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"

	uuid "github.com/google/uuid"
)

type DeserializeUser struct {
	ID uuid.UUID `json:"id"`
}

type Config struct {
	Filter       func(c *fiber.Ctx) bool // Required
	RequestAUTH  func(c *fiber.Ctx) (*DeserializeUser, error)
	Unauthorized fiber.Handler // middleware specfic
}

var ConfigDefault = Config{
	Filter:       nil,
	RequestAUTH:  nil,
	Unauthorized: nil,
}

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	cfg := config[0]

	if cfg.Filter == nil {
		cfg.Filter = ConfigDefault.Filter
	}

	if cfg.RequestAUTH == nil {

		cfg.RequestAUTH = func(c *fiber.Ctx) (*DeserializeUser, error) {

			authHeader := c.Get("Authorization")

			if authHeader == "" {
				return nil, errors.New("authorization header is required")
			}

			bearerTokenList := strings.SplitN(authHeader, " ", -1)

			if len(bearerTokenList) != 2 && bearerTokenList[0] != "Bearer" {
				return nil, errors.New("error parsing token")
			}

			viper.SetConfigFile(".env")

			viper.ReadInConfig()

			host_auth, _ := viper.Get("HOST_AUTH").(string)

			req, err := http.NewRequest(http.MethodGet, host_auth, nil)
			if err != nil {
				return nil, errors.New("could not create request")
			}

			bearerToken := fmt.Sprintf("Bearer %s", bearerTokenList[1])
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", bearerToken)

			res, err := http.DefaultClient.Do(req)

			if err != nil {
				return nil, errors.New("error making http request")
			}

			if res.StatusCode != 202 {
				return nil, errors.New("invalid token")
			}

			defer res.Body.Close()

			var user DeserializeUser

			err = json.NewDecoder(res.Body).Decode(&user)

			if err != nil {
				return nil, errors.New("could not deserialezer response body")
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

func New(config Config) fiber.Handler {
	cfg := configDefault(config)

	return func(c *fiber.Ctx) error {
		if cfg.Filter != nil && cfg.Filter(c) {
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
