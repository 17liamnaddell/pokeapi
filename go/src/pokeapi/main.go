package main

import "net/http"
import "log"
import "fmt"
import "io/ioutil"

//import "encoding/json"
import "os"

type mytype []interface{}

func main() {
	poke, _ := http.Get("http://pokeapi.co/api/v2/pokemon/")
	text, _ := ioutil.ReadAll(poke.Body)
	bulbas, err := http.Get("http://pokeapi.co/api/v2/pokemon/1")
	drive, err1 := ioutil.ReadAll(bulbas.Body)
	checkerr(err)
	checkerr(err1)

	data := map[string]interface{}{}
	err4 := json.Unmarshal(poke, &data)
	AssignPokemon(poke)
}
func AssignPokemon() string {

}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
