package domain

import "github.com/google/uuid"

type Pokemon struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	PokedexNumber  int    `json:"pokedex_number"`
	ImgName        string `json:"img_name"`
	Generation     int    `json:"generation"`
	EvolutionStage string `json:"evolution_stage"`
	Evolved        int    `json:"evolved"`
	FamilyID       int    `json:"family_id"`
	CrossGen       int    `json:"cross_gen"`
	TypeOne        string `json:"type_one"`
	TypeTwo        string `json:"type_two"`
	WeatherOne     string `json:"weather_one"`
	WeatherTwo     string `json:"weather_two"`
	StatTotal      int    `json:"stat_total"`
	ATK            int    `json:"atk"`
	DEF            int    `json:"def"`
	STA            int    `json:"sta"`
	Legendary      int    `json:"legendary"`
	Aqueireable    int    `json:"aquireable"`
	Spawns         int    `json:"spawns"`
	Regional       int    `json:"regional"`
	Raidable       int    `json:"raidable"`
	Hatchable      int    `json:"hatchable"`
	Shiny          int    `json:"shiny"`
	Nest           int    `json:"nest"`
	New            int    `json:"new"`
	NotGettable    int    `json:"not_gettable"`
	FutureEvolve   int    `json:"future_evolve"`
}

func NewPokemon(name string) *Pokemon {
	pokeUUID, _ := uuid.NewUUID()
	return &Pokemon{
		ID:   pokeUUID.String(),
		Name: name,
	}
}
