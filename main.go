// RapidRig-Go
// Written by Dahlton

// RapidRig-Go aims to assist in automating setting up
// a local development environment by making or using
// a bare Git repository and Git submodules to manage
// dotfiles in an organized and effective manner

// Check out https://github.com/RapidRig/rapidrig-go for more
package main

import (
	"github.com/RapidRig/rapidrig-go/cmd"
	"github.com/RapidRig/rapidrig-go/internal/app"
	"github.com/RapidRig/rapidrig-go/internal/modules"
	_ "github.com/RapidRig/rapidrig-go/internal/modules"
)

// main starts everything off, now handled by Cobra.
func main() {
	newApp := app.New()
	for _, module := range modules.ToBeRegistered {
		module.SetApp(newApp)
		newApp.ModuleRegistry.Register(module)
	}
	_ = cmd.ExecuteApplication(newApp)
}
