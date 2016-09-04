package main

import "net/http"
import "log"
import "fmt"
import "io/ioutil"
import "encoding/json"

type mytype []interface{}

func main() {
	poke, _ := http.Get("http://pokeapi.co/api/v2/pokemon/")
	text, _ := ioutil.ReadAll(poke.Body)
	ioutil.WriteFile("output.txt", text, 0666)
	bulbas, err := http.Get("http://pokeapi.co/api/v2/pokemon/1")
	drive, err1 := ioutil.ReadAll(bulbas.Body)
	ioutil.WriteFile("inauth.txt", drive, 0666)
	checkerr(err)
	checkerr(err1)

	data := map[string]interface{}{}
	err4 := json.Unmarshal(drive, &data)
	fmt.Println("errror: ", err4)
	inn := data["abilities"].([]interface{})
	fmt.Println(inn)
	fmt.Println(inn[0])
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
