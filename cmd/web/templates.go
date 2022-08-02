package main

import "github.com/stemitom/go-snippet/pkg/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
