package backend

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/options"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) DOMReady(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
}

func SetBackgroundColor() *options.RGBA {
	return &options.RGBA{
		R: 12,
		G: 13,
		B: 29,
		A: 255,
	}
}
