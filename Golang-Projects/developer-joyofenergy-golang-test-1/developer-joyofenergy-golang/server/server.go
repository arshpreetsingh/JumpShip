package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"

	"joi-energy-golang/endpoint/priceplans"
	"joi-energy-golang/endpoint/readings"
	"joi-energy-golang/repository"
)

const (
	serverPort = "localhost:8080"
)

// Run starts the HTTP server
func Run() {
	log.SetFormatter(&log.JSONFormatter{})

	handler := setUpServer()
	srv := &http.Server{Addr: serverPort, Handler: handler}
	go func() {
		log.Info("Starting server")

		err := srv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Info("Server shut down. Waiting for connections to drain.")
			} else {
				log.WithError(err).
					WithField("server_port", srv.Addr).
					Fatal("failed to start server")
			}
		}
	}()

	// Wait for an interrupt
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)    // interrupt signal sent from terminal
	signal.Notify(sigint, syscall.SIGTERM) // sigterm signal sent from system
	<-sigint

	log.Info("Shutting down server")

	attemptGracefulShutdown(srv)
}

func setUpServer() http.Handler {
	accounts := repository.NewAccounts(defaultSmartMeterToPricePlanAccounts())
	meterReadings := repository.NewMeterReadings(
		defaultMeterElectricityReadings(),
	)
	pricePlans := repository.NewPricePlans(
		defaultPricePlans(),
		&meterReadings,
	)

	// We can use here mux package as well. But Why?
	/*
		Newer developers find the distinction between http.Handle and
		http.HandleFunc subtle and confusing so I think it's worth understanding
		the http.Handler concept up front. A mux is just another kind of
		http.Handler (one that routes requests to other http.Handlers) and
		that reality is hidden away when you rely on the DefaultServeMux.
	*/
	//
	mux := http.NewServeMux()

	readingsLogger := log.WithField("endpoint", "readings")
	readingsService := readings.NewService(readingsLogger, &meterReadings)
	mux.Handle("/readings/store", readings.MakeStoreReadingsHandler(readingsService, readingsLogger))
	mux.Handle("/readings/read/", readings.MakeGetReadingsHandler(readingsService, readingsLogger))

	pricePlansLogger := log.WithField("endpoint", "pricePlans")
	pricePlansService := priceplans.NewService(pricePlansLogger, &pricePlans, &accounts)
	mux.Handle("/price-plans/compare-all/", priceplans.MakeCompareAllPricePlansHandler(pricePlansService, pricePlansLogger))
	mux.Handle("/price-plans/recommend/", priceplans.MakeRecommendPricePlansHandler(pricePlansService, pricePlansLogger))
	mux.Handle("/price-plans/recommend-weekly/", priceplans.MakeRecommendPricePlansWeeklyHandler(pricePlansService, pricePlansLogger))

	return mux
}

func attemptGracefulShutdown(srv *http.Server) {
	if err := shutdownServer(srv, 25*time.Second); err != nil {
		log.WithError(err).Error("failed to shutdown server")
	}
}

func shutdownServer(srv *http.Server, maximumTime time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), maximumTime)
	defer cancel()
	return srv.Shutdown(ctx)
}
