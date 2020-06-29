package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/charlires/go-app/usecase"
)

type PokeUseCase interface {
	Compare(pokeName1, pokeName2 string) (*usecase.PokeComparative, error)
}

// Pokemon controller struct
type Pokemon struct {
	render      *render.Render
	pokeUseCase PokeUseCase
}

// NewPokemon returns a Pokemon controller
func NewPokemon(r *render.Render, pu PokeUseCase) *Pokemon {
	return &Pokemon{r, pu}
}

// defeat - ..
func (p *Pokemon) Compare(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	pokeComparative, err := p.pokeUseCase.Compare(pathParams["pokemon1"], pathParams["pokemon2"])
	if err != nil {
		p.render.Text(w, http.StatusInternalServerError, "something failed")
	}
	p.render.JSON(w, http.StatusOK, pokeComparative)
}

// CommonMoves - ..
func (p *Pokemon) CommonMoves(w http.ResponseWriter, r *http.Request) {
	p.render.Text(w, http.StatusOK, "This is fine! 🔥\n")
}
