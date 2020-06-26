package controller

import (
	"net/http"

	"github.com/unrolled/render"
)

// Demo controller struct
type Demo struct {
	render *render.Render
}

// NewDemo returns a Demo controller
func NewDemo(r *render.Render) *Demo {
	return &Demo{r}
}

// Demo - ..
func (p *Demo) Demo(w http.ResponseWriter, r *http.Request) {
	p.render.Text(w, http.StatusOK, "This is fine! 🔥\n")
}
