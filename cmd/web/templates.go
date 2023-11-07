package main

import "snippetbox.inthava.me/internal/models"

type TemplateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
