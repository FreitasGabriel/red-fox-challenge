package service

import (
	"github.com/FreitasGabriel/red-fox-challenge/config/logger"
	"github.com/FreitasGabriel/red-fox-challenge/internal/core/domain"
	"github.com/FreitasGabriel/red-fox-challenge/internal/utils"
	"github.com/xuri/excelize/v2"
)

func (ps *pokemonService) ReadFile(filepath string) ([]*domain.Pokemon, error) {

	pokemons := []*domain.Pokemon{}

	file, err := excelize.OpenFile(filepath)
	if err != nil {
		logger.Error("error to open file", err)
		return nil, err
	}

	defer file.Close()

	rows, err := file.GetRows("Sheet1")
	if err != nil {
		logger.Error("could not possible read the rows", err)
		return nil, err
	}

	if len(rows) > 0 {
		rows = rows[1:]
	}

	for _, cell := range rows {

		if len(cell) < 28 {
			logger.Info("Skipping row with insufficient columns. Row data: ")
			continue
		}

		name := cell[1]
		pokedex_number := utils.ConvertStringToInt(cell[2])
		img_name := cell[3]
		generation := utils.ConvertStringToInt(cell[4])
		evolution_stage := cell[5]
		evolved := utils.ConvertStringToInt(cell[6])
		family_id := utils.ConvertStringToInt(cell[7])
		cross_gen := utils.ConvertStringToInt(cell[8])
		type_one := cell[9]
		type_two := cell[10]
		weather_one := cell[11]
		weather_two := cell[12]
		stat_total := utils.ConvertStringToInt(cell[13])
		atk := utils.ConvertStringToInt(cell[14])
		def := utils.ConvertStringToInt(cell[15])
		sta := utils.ConvertStringToInt(cell[16])
		legendary := utils.ConvertStringToInt(cell[17])
		aquireable := utils.ConvertStringToInt(cell[18])
		spawns := utils.ConvertStringToInt(cell[19])
		regional := utils.ConvertStringToInt(cell[20])
		raidable := utils.ConvertStringToInt(cell[21])
		hatchable := utils.ConvertStringToInt(cell[22])
		shiny := utils.ConvertStringToInt(cell[23])
		nest := utils.ConvertStringToInt(cell[24])
		new := utils.ConvertStringToInt(cell[25])
		not_gettable := utils.ConvertStringToInt(cell[26])
		future_evolve := utils.ConvertStringToInt(cell[27])

		poke := domain.NewPokemon(name, img_name, evolution_stage, type_one, type_two, weather_one, weather_two, pokedex_number,
			generation, evolved, family_id, cross_gen, stat_total, atk, def, sta, legendary, aquireable, spawns, regional, raidable,
			hatchable, shiny, nest, new, not_gettable, future_evolve)

		pokemons = append(pokemons, poke)

	}
	return pokemons, nil
}
