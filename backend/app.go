package backend

import (
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called at application startup
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// DOMReady is called after the front-end dom has been loaded
func (a *App) DOMReady(ctx context.Context) {
	a.ctx = ctx
}

// Shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
}
