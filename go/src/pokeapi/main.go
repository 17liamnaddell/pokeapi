package main

import (
	//"bufio"
	"encoding/json"
	"fmt"
	C "github.com/skilstak/go-colors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Pokedata struct {
	Abilities []struct {
		Ability struct {
			Name string
		}
	} `json:"abilities"`
	Name  string `json:"name"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

type Idkwtth struct {
	Results []struct {
		Name string `json:"name"`
	} `json:"results"`
	Next string `json:"next"`
}

func main() {
	var home = os.Getenv("HOME")
	if _, err := os.Stat(home + "/.pokeapi"); os.IsNotExist(err) == true {
		fmt.Println(err)
		os.Chdir(os.Getenv("HOME"))
		err := os.Mkdir(".pokeapi", 0777)
		fmt.Println(err)
	}
	go captureCC()
	fmt.Println(os.Args)
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-f" || os.Args[i] == "--find" {
			os.Chdir(home + "/.pokeapi")
			pokedataDirs, _ := os.Open("/.pokeapi")
			myint := 0
			dirs, _ := pokedataDirs.Readdirnames(myint)
			for q := 0; q < len(dirs); q++ {
				if dirs[i] == os.Args[i+1] {
					fmt.Println("kekappa")
				}
			}

			pokedat := GetPokemon(os.Args[i+1])

			printit(pokedat)
		} else if os.Args[i] == "-rf" || os.Args[i] == "--removeData" {
			os.Chdir(home + "/.pokeapi")
			os.RemoveAll(".")
		} else if os.Args[i] == "-fa" || os.Args[i] == "--findAll" {
			//list all pokemon
			ListPokemon()
			//get all data from each pokemon and write that to the file

		}
	}

}

var Allpokemon []string

func GetPokemon(name string) Pokedata {
	pokelink := "https://pokeapi.co/api/v2/pokemon/" + name
	fmt.Println(pokelink)
	poke, err := http.Get(pokelink)
	fmt.Println(poke.Body)
	checkerr(err)
	pokedat := Pokedata{}
	idontcare, _ := ioutil.ReadAll(poke.Body)
	_ = ioutil.WriteFile(name, idontcare, 0777)
	err = json.Unmarshal(idontcare, &pokedat)
	checkerr(err)
	return pokedat
}

func ListPokemon() {
	raw, err := http.Get("https://pokeapi.co/api/v2/pokemon")
	checkerr(err)
	rawJson, _ := ioutil.ReadAll(raw.Body)
	poka := Idkwtth{}
	err = json.Unmarshal(rawJson, &poka)
	for i := 0; i < len(poka.Results); i++ {
		Allpokemon = append(Allpokemon, poka.Results[i].Name)
	}
	log.Print(Allpokemon)
}

func captureCC() {
	for {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		b := <-c
		if b == os.Interrupt {
			fmt.Println(C.B + "type done to leave, or Enter a pokemon")
		}
	}
}

func printit(data Pokedata) {
	fmt.Println(C.R+data.Name, ":")

	var whtspc = `	`

	rang := len(data.Abilities)

	for i := 0; i < rang; i++ {
		fmt.Println(C.B+whtspc, "Ability ", i, ": ", data.Abilities[i].Ability.Name)
	}
	fmt.Println(C.M+whtspc, "Weight"+": ", data.Weight)
	for q := 0; q < len(data.Types); q++ {
		fmt.Println(C.G+whtspc, "Type ", q, ": ", data.Types[q].Type.Name+C.Y)
	}
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
