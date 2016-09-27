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
			Name string `json:"name"`
		} `json:"ability"`
	} `json:"abilities"`
	Name  string `json:"name"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func main() {
	go captureCC()
	fmt.Println(os.Args)
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-f" || os.Args[i] == "--find" {
			pokelink := "https://pokeapi.co/api/v2/pokemon/" + os.Args[i+1]
			fmt.Println(pokelink)
			poke, err := http.Get(pokelink)
			fmt.Println(poke.Body)
			checkerr(err)
			pokedat := Pokedata{}
			idontcare, _ := ioutil.ReadAll(poke.Body)
			err = json.Unmarshal(idontcare, &pokedat)
			fmt.Println(pokedat)
			checkerr(err)
			printit(pokedat)
		}
	}

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
