package parser

import (
	"context"
	"fmt"

	"github.com/stl0755/icp-tool/model"
)

// App struct
type App struct {
	Ctx      context.Context
	Features []*model.Feature
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
