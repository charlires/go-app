package service

import (
	"encoding/json"
	"fmt"

	"gopkg.in/resty.v1"
)

const (
	pokemonEndpoint     = "/api/v2/pokemon/{pokeName}"
	pokemonTypeEndpoint = "/api/v2/type/{pokeType}"
)

// Pokemon
type Pokemon struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	Moves []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
	} `json:"moves"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

type Type struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// PokemonMove
type PokemonType struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	MoveDamageClass struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"move_damage_class"`
	DamageRelations struct {
		DoubleDamageFrom []Type `json:"double_damage_from"`
		DoubleDamageTo   []Type `json:"double_damage_to"`
		HalfDamageFrom   []Type `json:"half_damage_from"`
		HalfDamageTo     []Type `json:"half_damage_to"`
		NoDamageFrom     []Type `json:"no_damage_from"`
		NoDamageTo       []Type `json:"no_damage_to"`
	} `json:"damage_relations"`
}

// Service for PokeAPI requests
type Service struct {
	client *resty.Client
}

// New creates a new PokeAPI client
func NewPokemon(host string) *Service {
	client := resty.New().
		SetHostURL(host).
		OnAfterResponse(func(c *resty.Client, r *resty.Response) error {
			return nil
		})

	return &Service{client}
}

func (s *Service) GetPokemon(
	pokeName string,
) (*Pokemon, error) {
	out := &Pokemon{}
	resp, err := s.client.R().
		SetPathParams(map[string]string{"pokeName": pokeName}).
		SetHeader("Accept", "application/json").
		Get(pokemonEndpoint)

	if err != nil {
		return nil, fmt.Errorf("getting pokemon from pokeapi")
	}

	body := resp.Body()
	if err := json.Unmarshal(body, out); err != nil {
		return nil, fmt.Errorf("marshalling pokemon")
	}

	return out, nil
}

func (s *Service) GetPokemonType(
	pokeType string,
) (*PokemonType, error) {
	out := &PokemonType{}
	resp, err := s.client.R().
		SetPathParams(map[string]string{"pokeType": pokeType}).
		SetHeader("Accept", "application/json").
		Get(pokemonTypeEndpoint)

	if err != nil {
		return nil, fmt.Errorf("getting pokemon type from pokeapi")
	}

	body := resp.Body()
	if err := json.Unmarshal(body, out); err != nil {
		return nil, fmt.Errorf("marshalling pokemon type")
	}

	return out, nil
}
