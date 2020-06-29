package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// DemoController - Interface for demo requests
type DemoController interface {
	Demo(w http.ResponseWriter, r *http.Request)
}

// PokemonController - Interface for Pokemon requests
type PokemonController interface {
	Compare(w http.ResponseWriter, r *http.Request)
	CommonMoves(w http.ResponseWriter, r *http.Request)
}

// Setup returns router instance which is used in main package to register handlers.
func Setup(
	demoController DemoController,
	pokemonController PokemonController,
) http.Handler {
	r := mux.NewRouter()

	// Demos endpoints
	r.HandleFunc(
		"/",
		demoController.Demo).Methods("GET").Name("demo")

	// Poke Endpoints
	pokeRouter := r.PathPrefix("/pokeapi/v1/pokemons").Subrouter()
	pokeRouter.HandleFunc(
		"/{pokemon1}/defeat/{pokemon2}",
		pokemonController.Compare).Methods("GET").Name("defeat")
	pokeRouter.HandleFunc(
		"/{pokemon1}/common_moves/{pokemon2}",
		pokemonController.CommonMoves).Methods("GET").Name("common_moves")

	return r
}
