package usecase

import (
	"github.com/charlires/go-app/service"
)

type PokeService interface {
	GetPokemon(pokeName string) (*service.Pokemon, error)
	GetPokemonType(pokeType string) (*service.PokemonType, error)
}

type PokeComparative struct {
	PokemonName1   string `json:"pokemon_name_1"`
	PokemonType1   string `json:"pokemon_type_1"`
	PokemonName2   string `json:"pokemon_name_2"`
	PokemonType2   string `json:"pokemon_type_2"`
	DoubleDamageTo bool   `json:"double_damage_to"`
	HalfDamageFrom bool   `json:"half_damage_from"`
	NoDamageFrom   bool   `json:"no_damage_from"`
}

// Pokemon controller struct
type Pokemon struct {
	pokeService PokeService
}

// NewPokemon returns a Pokemon controller
func NewPokemon(ps PokeService) *Pokemon {
	return &Pokemon{ps}
}

// Compare - ..
func (p *Pokemon) Compare(pokeName1, pokeName2 string) (*PokeComparative, error) {
	var pokemon1DoubleDamagePokemon2 bool
	var pokemon2HalfDamagePokemon1 bool
	var pokemon2NoDamagePokemon1 bool

	pokemon1, err := p.pokeService.GetPokemon(pokeName1)
	if err != nil {
		return nil, err
	}
	pokeType1, err := p.pokeService.GetPokemonType(pokemon1.Types[0].Type.Name)
	if err != nil {
		return nil, err
	}
	pokemon2, err := p.pokeService.GetPokemon(pokeName2)
	if err != nil {
		return nil, err
	}
	pokeType2, err := p.pokeService.GetPokemonType(pokemon2.Types[0].Type.Name)
	if err != nil {
		return nil, err
	}

	damageToRelations := map[string][]service.Type{
		"double": pokeType1.DamageRelations.DoubleDamageTo,
	}

	damageFromRelations := map[string][]service.Type{
		"half": pokeType1.DamageRelations.HalfDamageFrom,
		"none": pokeType1.DamageRelations.NoDamageFrom,
	}

	for key, damageToRelation := range damageToRelations {
		for _, pokeType := range damageToRelation {
			switch key {
			case "double":
				pokemon1DoubleDamagePokemon2 = pokeType.Name == pokeType2.Name
			}
		}
	}

	for key, damageFromRelation := range damageFromRelations {
		for _, pokeType := range damageFromRelation {
			switch key {
			case "half":
				pokemon2HalfDamagePokemon1 = pokeType.Name == pokeType2.Name
			case "none":
				pokemon2NoDamagePokemon1 = pokeType.Name == pokeType2.Name
			}
		}
	}

	return &PokeComparative{
		PokemonName1:   pokemon1.Name,
		PokemonType1:   pokeType1.Name,
		PokemonName2:   pokemon2.Name,
		PokemonType2:   pokeType2.Name,
		DoubleDamageTo: pokemon1DoubleDamagePokemon2,
		HalfDamageFrom: pokemon2HalfDamagePokemon1,
		NoDamageFrom:   pokemon2NoDamagePokemon1,
	}, nil
}
