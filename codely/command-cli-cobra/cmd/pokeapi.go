/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this csv except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"example.com/codelytv/cobra-03/cmd/ports"
	"fmt"

	"github.com/spf13/cobra"
)

const pathFileFlag = "pathfile"
const filenameFlag = "filename"
const offsetFlag = "offset"
const limitFlag = "limit"

func Init(getPokemonsRepo ports.GetPokemonsRepo, savePokemons ports.SavePokemonsInFileRepo) *cobra.Command {
	var pokeApiCmd = &cobra.Command{
		Use:   "pokeapi",
		Short: "Genera un archivo en base a la respuesta api https://pokeapi.co/api/v2/pokemon?offset=100&limit=100",
		Long: `Genera un archivo con los datos de cada pokemon entre un offset y un limite de tamaño de respuesta. Por ejemplo si se 
			consulta con un offset=0 y un limit 100, obtiene los 100 primeros pokemons de la pokedex.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("pokeapi called")
			processCmd(cmd, getPokemonsRepo, savePokemons)
		},
	}
	pokeApiCmd.Flags().StringP(pathFileFlag, "p", "/home", "Ruta absoluta donde guardar el archivo")
	pokeApiCmd.Flags().StringP(filenameFlag, "f", "pokemons", "Nombre del archivo donde se guardaran los resultados")
	pokeApiCmd.Flags().IntP(offsetFlag, "o", 0, "Define a partir desde que número de pokemon realiza la búsqueda")
	pokeApiCmd.Flags().IntP(limitFlag, "l", 100, "Define cuantos pokemones queres almacenar")
	return pokeApiCmd
}

func processCmd(cmd *cobra.Command, getPokemonsRepo ports.GetPokemonsRepo, savePokemons ports.SavePokemonsInFileRepo) {
	pathFile, _ := cmd.Flags().GetString(pathFileFlag)
	fmt.Printf("pathfile -> %s \n", pathFile)
	fileName, _ := cmd.Flags().GetString(filenameFlag)
	fmt.Printf("filename -> %s \n", fileName)
	offSet, _ := cmd.Flags().GetInt(offsetFlag)
	fmt.Printf("offset -> %d \n", offSet)
	limit, _ := cmd.Flags().GetInt(limitFlag)
	fmt.Printf("limit -> %d \n", limit)
	var pokemons, error = getPokemonsRepo.Execute(limit, offSet)
	if error != nil {
		return
	}
	savePokemons.Execute(pokemons, pathFile, fileName)

}
