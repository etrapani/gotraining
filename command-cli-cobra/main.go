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
package main

import (
	"example.com/gotraining/command-cli-cobra/cmd"
	"example.com/gotraining/command-cli-cobra/cmd/ports"
	"example.com/gotraining/command-cli-cobra/infraestructure/repository/csv"
	"example.com/gotraining/command-cli-cobra/infraestructure/repository/pokeapi"
	"github.com/spf13/cobra"
)

func main() {
	//f, _ := os.Create("pokeapi.cpu.prof")
	//defer f.Close()
	//pprof.StartCPUProfile(f)
	//defer pprof.StopCPUProfile()

	var getPokemonsRepo ports.GetPokemonsRepo
	getPokemonsRepo = pokeapi.NewPokeApiRepository()

	var savePokemonsRepo ports.SavePokemonsInFileRepo
	savePokemonsRepo = csv.NewPokeApiRepository()

	rootCmd := &cobra.Command{Use: "pokeapi-cli"}
	rootCmd.AddCommand(cmd.Init(getPokemonsRepo, savePokemonsRepo))
	rootCmd.Execute()

	//f, _ := os.Create("pokeapi.mem.prof")
	//defer f.Close()
	//pprof.WriteHeapProfile(f)
}
