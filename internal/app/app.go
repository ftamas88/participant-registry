package app

import (
	"context"
	"fmt"
	"grail-participant-registry/internal/domain"
	"grail-participant-registry/internal/service"
	"grail-participant-registry/internal/storage/memory"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"

	"grail-participant-registry/internal/controller"
	"grail-participant-registry/internal/routing"

	"github.com/sirupsen/logrus"
)

type App struct {
	Context    context.Context
	Server     *http.Server
	Cancel     context.CancelFunc
	Logger     *logrus.Entry
	version    string
	commitHash string
}

type Config struct {
	Context  context.Context
	HTTPPort int
}

// Component is part of the app which runs in a goroutine
type Component interface {
	// Run should cleanup the component if the context is done, and should wait for any current
	// tasks to complete before returning an error if there was one.
	Run(ctx context.Context) error
}

var version string
var commitHash string

// New returns a new instance of the App
func New(conf Config) *App {
	ctx, cancelFunc := context.WithCancel(conf.Context)
	go shutdownHandler(cancelFunc)

	// Pass this logger into Structs which write logs
	logger := logrus.WithFields(logrus.Fields{
		"app_version":        Version(),
		"app_commit_version": CommitHash(),
	})

	pRepo := memory.NewParticipantRepository()

	router := routing.NewRouter(
		&routing.RouterConfig{
			WellKnown: &controller.WellKnownController{
				AppVersion:    Version(),
				AppCommitHash: CommitHash(),
			},
			Participant: &controller.ParticipantController{
				Service: service.NewParticipantService(
					pRepo,
				),
			},
		},
	)

	go func() {
		// Random seed for testing
		<-time.After(3 * time.Second)
		p1 := domain.Participant{
			ID:          domain.ParticipantID(uuid.New().String()),
			Reference:   domain.ParticipantReference("CD-12-EF"),
			Name:        "Test",
			DateOfBirth: time.Now(),
			Phone:       "123-123-123",
			Address:     domain.ParticipantAddress{},
		}

		_ = pRepo.Add(p1)

		<-time.After(3 * time.Second)
		p1.Name = "Test2"
		_ = pRepo.Update(p1)
	}()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.HTTPPort),
		Handler: router,
	}

	return &App{
		Context:    ctx,
		Server:     server,
		Cancel:     cancelFunc,
		Logger:     logger,
		version:    version,
		commitHash: commitHash,
	}
}

// Version returns the git tag used to build the application
func Version() string {
	return version
}

// CommitHash returns the git commit hash used to build the application
func CommitHash() string {
	return commitHash
}

// Run starts the main web-server
//
// Returns the first error encountered or nil once all components shutdown gracefully.
func (app App) Run() error {
	app.Logger.Infof("HTTP server running on port %s", app.Server.Addr)

	return app.Server.ListenAndServe()
}

// shutdownHandler listens for a SIGTERM signal
// and gracefully cancels the main application context
// once this is completed exits the app
func shutdownHandler(cancelFunction context.CancelFunc) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	// Invoke the cancel function
	cancelFunction()

	os.Exit(1)
}
