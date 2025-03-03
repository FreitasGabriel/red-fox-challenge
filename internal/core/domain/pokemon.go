package domain

type Pokemon struct {
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
	Aquireable     int    `json:"aquireable"`
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

func NewPokemon(name, img_name, evolution_stage, type_one, type_two, weather_one, weather_two string,
	pokedex_number, generation, evolved, family_id, cross_gen, stat_total, atk, def, sta, legendary, aquireable,
	spawns, regional, raidable, hatchable, shiny, nest, new, not_gettable, future_evolve int) *Pokemon {

	return &Pokemon{
		Name:           name,
		PokedexNumber:  pokedex_number,
		ImgName:        img_name,
		Generation:     generation,
		EvolutionStage: evolution_stage,
		Evolved:        evolved,
		FamilyID:       family_id,
		CrossGen:       cross_gen,
		TypeOne:        type_one,
		TypeTwo:        type_two,
		WeatherOne:     weather_one,
		WeatherTwo:     weather_two,
		StatTotal:      stat_total,
		ATK:            atk,
		DEF:            def,
		STA:            sta,
		Legendary:      legendary,
		Aquireable:     aquireable,
		Spawns:         spawns,
		Regional:       regional,
		Raidable:       raidable,
		Hatchable:      hatchable,
		Shiny:          shiny,
		Nest:           nest,
		New:            new,
		NotGettable:    not_gettable,
		FutureEvolve:   future_evolve,
	}
}
