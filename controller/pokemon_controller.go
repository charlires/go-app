package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/charlires/go-app/service"
)

type PokeService interface {
	GetPokemon(pokeName string) (*service.Pokemon, error)
	GetPokemonType(pokeType string) (*service.PokemonType, error)
}

// Pokemon controller struct
type Pokemon struct {
	render      *render.Render
	pokeService PokeService
}

// NewPokemon returns a Pokemon controller
func NewPokemon(r *render.Render, ps PokeService) *Pokemon {
	return &Pokemon{r, ps}
}

// defeat - ..
func (p *Pokemon) Defeat(w http.ResponseWriter, r *http.Request) {
	var pokemon1DoubleDamagePokemon2 bool
	var pokemon1HalfDamagePokemon2 bool
	var pokemon1NoDamagePokemon2 bool

	var pokemon2DoubleDamagePokemon1 bool
	var pokemon2HalfDamagePokemon1 bool
	var pokemon2NoDamagePokemon1 bool
	var msg string

	pathParams := mux.Vars(r)
	pokemon1, err := p.pokeService.GetPokemon(pathParams["pokemon1"])
	if err != nil {
		p.render.JSON(w, http.StatusOK, err)
	}
	pokeType1, err := p.pokeService.GetPokemonType(pokemon1.Types[0].Type.Name)
	if err != nil {
		p.render.JSON(w, http.StatusOK, err)
	}
	pokemon2, err := p.pokeService.GetPokemon(pathParams["pokemon2"])
	if err != nil {
		p.render.JSON(w, http.StatusOK, err)
	}
	pokeType2, err := p.pokeService.GetPokemonType(pokemon2.Types[0].Type.Name)
	if err != nil {
		p.render.JSON(w, http.StatusOK, err)
	}

	damageToRelations := map[string][]service.Type{
		"double": pokeType1.DamageRelations.DoubleDamageTo,
		"half":   pokeType1.DamageRelations.HalfDamageTo,
		"none":   pokeType1.DamageRelations.NoDamageTo,
	}

	damageFromRelations := map[string][]service.Type{
		"double": pokeType1.DamageRelations.DoubleDamageFrom,
		"half":   pokeType1.DamageRelations.HalfDamageFrom,
		"none":   pokeType1.DamageRelations.NoDamageFrom,
	}

	for key, damageToRelation := range damageToRelations {
		for _, pokeType := range damageToRelation {
			switch key {
			case "double":
				pokemon1DoubleDamagePokemon2 = pokeType.Name == pokeType2.Name
			case "half":
				pokemon1HalfDamagePokemon2 = pokeType.Name == pokeType2.Name
			case "none":
				pokemon1NoDamagePokemon2 = pokeType.Name == pokeType2.Name
			}
		}
	}

	for key, damageFromRelation := range damageFromRelations {
		for _, pokeType := range damageFromRelation {
			switch key {
			case "double":
				pokemon2DoubleDamagePokemon1 = pokeType.Name == pokeType2.Name
			case "half":
				pokemon2HalfDamagePokemon1 = pokeType.Name == pokeType2.Name
			case "none":
				pokemon2NoDamagePokemon1 = pokeType.Name == pokeType2.Name
			}
		}
	}

	// Both pokemons can cause same damage
	if (pokemon1DoubleDamagePokemon2 && pokemon2DoubleDamagePokemon1) ||
		(pokemon1HalfDamagePokemon2 && pokemon2HalfDamagePokemon1) {
		msg = "Both pokemons can win"
	}

	// First pokemon is weak that the second
	if pokemon1NoDamagePokemon2 && (pokemon2DoubleDamagePokemon1 || pokemon2HalfDamagePokemon1) {
		msg = "Second pokemon will defeat the first"
	}

	// First pokemon is weak that the second
	if pokemon1DoubleDamagePokemon2 || pokemon2HalfDamagePokemon1 {
		msg = "First pokemon will defeat the second"
	}

	// Second pokemon is weak that the first
	if pokemon2NoDamagePokemon1 && (pokemon1DoubleDamagePokemon2 || pokemon1HalfDamagePokemon2) {
		msg = "First pokemon will defeat the second"
	}

	// Second pokemon is weak that the first
	if pokemon2DoubleDamagePokemon1 || pokemon1HalfDamagePokemon2 {
		msg = "Second pokemon will defeat the first"
	}

	if msg == "" {
		msg = "Both pokemons can win"
	}

	p.render.JSON(w, http.StatusOK, map[string]interface{}{
		"pokemon1": pokeType1,
		"pokemon2": pokeType2,
		"message":  msg,
	})
}

// CommonMoves - ..
func (p *Pokemon) CommonMoves(w http.ResponseWriter, r *http.Request) {
	p.render.Text(w, http.StatusOK, "This is fine! 🔥\n")
}
