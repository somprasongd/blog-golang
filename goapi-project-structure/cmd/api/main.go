package main

import (
	"goapi-project-structure/pkg/app"
	"goapi-project-structure/pkg/config"
	"goapi-project-structure/pkg/module"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	app := app.New(cfg)
	// Cleanup when server stopped
	defer app.Close()

	// For Liveness Probe
	app.CreateLivenessFile()

	// Initialize data sources
	app.InitDS()

	// Create router (mux/gin/fiber)
	app.InitRouter()

	// Initialize module with dependency injection
	module.Init(app.Context)
	// Start server
	app.ServeHTTP()
}