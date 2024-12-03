package main

import "github.com/nimilgp/paste-bin/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
