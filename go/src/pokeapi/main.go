package main

import "net/http"

//import "fmt"
import "io/ioutil"

func main() {
	poke, _ := http.Get("http://pokeapi.co/api/v2/pokemon/")
	text, _ := ioutil.ReadAll(poke.Body)
	ioutil.WriteFile("output.txt", text, 0666)
}
