package controllers

import (
	"net/http"
	"picapp/views"
)

func NewGallery() *Galleries{
return &Galleries{
	NewGallery: views.NewView("bootstrap", "galleries/new"),
}
}

type Galleries struct {
	NewGallery *views.View
}

func(g *Galleries) New(w http.ResponseWriter, r *http.Request) {
	if err := g.NewGallery.Render(w, nil); err != nil {
		panic(err)
	}
}