package server

import (
	"context"
	"fmt"

	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/slok/go-http-metrics/middleware"

	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	stdmiddleware "github.com/slok/go-http-metrics/middleware/std"
)

type RegisterController interface {
	HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}

type RegisterControllerHandler func(srv RegisterController)

type Server interface {
	Run(ctx context.Context) error
	RegisterController(RegisterControllerHandler)
	RegisterHandler(handler http.Handler)
	Close() error
}

var _ Server = (*httpServer)(nil)

type Config struct {
	Logger *slog.Logger
	Env    string `env:"APP_ENV" envDefault:"develop"`
	Port   int    `env:"APP_PORT" envDefault:"8080"`
}

type httpServer struct {
	cfg    Config
	log    *slog.Logger
	srv    *http.Server
	mux    *http.ServeMux
	doneCh chan struct{}
}

func New(cfg Config) Server {
	return &httpServer{
		cfg:    cfg,
		log:    cfg.Logger,
		mux:    http.NewServeMux(),
		doneCh: make(chan struct{}),
	}
}

func (s *httpServer) RegisterController(r RegisterControllerHandler) {
	r(s.mux)
}

func (s *httpServer) RegisterHandler(h http.Handler) {
	s.mux.Handle("GET /metrics", h)
}

func (s *httpServer) Run(ctx context.Context) error {
	metricsMiddleware := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	h := stdmiddleware.Handler("", metricsMiddleware, logRequest(s.mux, s.log))

	s.srv = &http.Server{
		Addr:         "0.0.0.0:" + fmt.Sprintf("%d", s.cfg.Port),
		Handler:      h,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     log.New(io.Discard, "", 0),
	}

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	go func() {
		if err := s.runHttpServer(); err != nil {
			s.log.Error("(s.runHttpServer)", slog.Any("error", err.Error()))
			cancel()
		}
	}()
	s.log.Info("server is listening", slog.Int("port", s.cfg.Port))

	<-ctx.Done()

	s.waitShootDown(3 * time.Second)

	<-s.doneCh

	s.log.Info("server exited properly")
	return nil
}

func (s httpServer) Close() error {
	go func() {
		s.doneCh <- struct{}{}
	}()

	return s.srv.Close()
}

func (s httpServer) waitShootDown(duration time.Duration) {
	s.srv.Close()

	go func() {
		time.Sleep(duration)
		s.doneCh <- struct{}{}
	}()
}

func (s httpServer) runHttpServer() error {
	return s.srv.ListenAndServe()
}
