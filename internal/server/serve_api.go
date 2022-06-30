package server

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/taalhach/aroundhome-challennge/docs"
	"github.com/taalhach/aroundhome-challennge/internal/server/apihandlers"
)

const (
	port = 3000
)

// @title aroundhome-challennge API docs
// @version 0.1
// @description aroundhome's code aroundhome API specs.

// @contact.name Muhammad Talha
// @contact.email talhach891@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000

var serverApi = &cobra.Command{
	Use:                   "serve_api",
	Short:                 fmt.Sprintf("servers api on %v port", port),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := loadConfigs()
		if err != nil {
			fmt.Printf("failed to load configs: %v", err)
			os.Exit(1)
		}

		// init customised validator
		customValidator := &CustomValidator{validator: validator.New()}
		if err := customValidator.Init(); err != nil {
			fmt.Printf("Failed to init validator: %v\n", err)
			os.Exit(1)
		}

		e := echo.New()
		//attach custom error handler
		e.HTTPErrorHandler = customHTTPErrorHandler

		//attach validator
		e.Validator = customValidator

		// request logger middleware which logs each request in specified log form
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: fmt.Sprintf("method=${method}, uri=${uri}, status=${status}, latency=${latency_human}, in=${bytes_in}, out=${bytes_out}\n"),
		}))

		// middleware to recover from panics calls HttpErrorHandler
		e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
			StackSize:         1 << 10, // 1 KB
			DisableStackAll:   true,
			DisablePrintStack: true,
		}))

		// compress middleware
		e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}))

		// CORS middleware to CORS errors
		e.Use(middleware.CORS())

		// middleware to attach database connection with each request
		// context is extended and database connection pool is attached
		e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				if strings.Contains(c.Request().Header.Get("Content-Encoding"), "gzip") {
					var err error
					// Decompress the stream
					c.Request().Body, err = gzip.NewReader(c.Request().Body)
					if err != nil {
						return err
					}

					defer c.Request().Body.Close()
				}
				return h(c)
			}
		})

		// now add all endpoints here

		// redirect to swagger api docs
		e.GET("/", func(c echo.Context) error {
			return c.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
		})

		// ping/pong request
		e.GET("ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})

		e.GET("partners", apihandlers.PartnersList)
		e.GET("partners/:id", apihandlers.PartnerDetails)
		e.GET("/swagger/*", echoSwagger.WrapHandler)

		// serves api on specified port
		listenAddress := fmt.Sprintf(":%d", port)
		e.Logger.Fatal(e.Start(listenAddress))
	},
}

func init() {
	rootCommand.AddCommand(serverApi)
}
