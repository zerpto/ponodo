package handlers

import (
	"context"
	"errors"
	"fmt"
	clicontracts "github.com/zerpto/ponodo/cli/contracts"
	"github.com/zerpto/ponodo/contracts"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/go-playground/validator/v10"
)

type HttpHandler struct {
	App           contracts.AppContract
	RouterSetupFn func(contracts.AppContract)
}

func (h *HttpHandler) Short() string {
	return "Run http server to expose api endpoints."
}

func (h *HttpHandler) Long() string {
	return "Run http server to expose api endpoints."
}

func (h *HttpHandler) Example() string {
	return `zerpto http --port 8080`
}

func (h *HttpHandler) Run(cmd *cobra.Command, args []string) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Set Gin
	cfg := h.App.GetConfigLoader().Config
	if cfg.GetDebug() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	h.App.SetGin(r)

	if h.RouterSetupFn != nil {
		h.RouterSetupFn(h.App)
	}

	// Set validator
	v := validator.New(validator.WithRequiredStructEnabled())
	h.App.SetValidator(v)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Msg(fmt.Sprintf("listen: %s\n", err))
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Info().Msg("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Msg(fmt.Sprintf("Server forced to shutdown: %v", err))
	}

	log.Info().Msg("Server exiting")
}

func (h *HttpHandler) Use() string {
	return "http"
}

func NewHttpHandler(app contracts.AppContract, routerSetupFn func(contracts.AppContract)) clicontracts.CommandContract {
	return &HttpHandler{
		App:           app,
		RouterSetupFn: routerSetupFn,
	}
}
