package main

import "net/http"
import "fmt"

func main() {
	poke, err := http.Get("http://pokeapi.co/api/v2/pokemon/1")

	fmt.Println(err, poke)
}
