package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Pokemon struct {
	Weight int `json:"weight"`
	Name string `json:"name"`
	Id int `json:"id"`
}

var pokeclient = http.Client{
	Timeout: time.Second * 10, // Maximum of 2 secs
}

func getLink(link string) *http.Response {
	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "pokeapi-getter-go")

	res, getErr := pokeclient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	return res
}
func getPokemon(name string) Pokemon {
	url := "http://pokeapi.co/api/v2/pokemon/" + name
	res := getLink(url)
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	pokemon := Pokemon{}
	jsonErr := json.Unmarshal(body, &pokemon)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	if pokemon.Weight == 0 && pokemon.Name == "" {
		log.Fatal("not a pookeman")
	}
	return pokemon
}
func main() {
	pokemon := getPokemon(os.Args[1])
	fmt.Println(pokemon.Weight)
	fmt.Println(pokemon.Name)
	fmt.Println(pokemon.Id)
}
