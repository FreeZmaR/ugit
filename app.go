package main

import (
	"context"
	"os/exec"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetGitStatus() (string, error) {
	cmd := exec.Command("git", "status")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func (a *App) GetBranches() ([]string, error) {
	cmd := exec.Command("git", "branch", "--list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	branches := strings.Split(string(output), "\n")
	for i := range branches {
		branches[i] = strings.TrimSpace(branches[i])
	}

	return branches, nil
}
