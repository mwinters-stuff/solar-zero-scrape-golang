// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/mwinters-stuff/solar-zero-scrape-golang/api/restapi/operations"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/api/restapi/operations/http_api"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/api/restapi/operations/kubernetes"
	"github.com/mwinters-stuff/solar-zero-scrape-golang/solarzero"
)

//go:generate swagger generate server --target ../../api --name SolarZeroScrape --spec ../swagger.yaml --principal interface{}

func configureFlags(api *operations.SolarZeroScrapeAPIAPI) {
	opts := &solarzero.SolarZeroOptions{}
	api.CommandLineOptionsGroups = append(api.CommandLineOptionsGroups, swag.CommandLineOptionsGroup{
		ShortDescription: "config",
		LongDescription:  "Solar Zero Config",
		Options:          opts,
	})
}

func configureAPI(api *operations.SolarZeroScrapeAPIAPI) http.Handler {
	cfg := api.CommandLineOptionsGroups[0].Options.(*solarzero.SolarZeroOptions)

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	api.ServeError = errors.ServeError

	api.Logger = log.Info().Msgf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	solarZeroScrape := solarzero.NewSolarZeroScrape(cfg)
	go solarZeroScrape.Start()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()
	// api.TxtProducer = runtime.TextProducer()

	api.HTTPAPIGetHandler = http_api.GetHandlerFunc(func(params http_api.GetParams) middleware.Responder {
		return http_api.NewGetOK().WithPayload(
			map[string]any{
				"currentData": solarZeroScrape.CurrentData(),
				"dayData":     solarZeroScrape.DayData(),
				"monthData":   solarZeroScrape.MonthData(),
				"yearData":    solarZeroScrape.YearData(),
			})
	})

	api.KubernetesGetHealthzHandler = kubernetes.GetHealthzHandlerFunc(func(params kubernetes.GetHealthzParams) middleware.Responder {
		if solarZeroScrape.Healthy() {
			return kubernetes.NewGetHealthzOK().WithPayload(map[string]string{"status": "OK"})
		}
		return middleware.Error(http.StatusNotFound, map[string]string{"status": "UNHEALTHY"})
	})

	api.HTTPAPIGetPanicHandler = http_api.GetPanicHandlerFunc(func(params http_api.GetPanicParams) middleware.Responder {
		os.Exit(255)
		return http_api.NewGetPanicOK()
	})

	api.KubernetesGetReadyzHandler = kubernetes.GetReadyzHandlerFunc(func(params kubernetes.GetReadyzParams) middleware.Responder {
		if solarZeroScrape.Ready() {
			return kubernetes.NewGetHealthzOK().WithPayload(map[string]string{"status": "OK"})
		}
		return middleware.Error(http.StatusNotFound, map[string]string{"status": "NOTREADY"})
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
