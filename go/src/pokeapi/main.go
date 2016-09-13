package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	Scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter a pokemon")
		Scanner.Scan()
		if Scanner.Text() == "done" {
			break
		}
		cp, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + Scanner.Text())
		checkerr(err)
		pnew, _ := ioutil.ReadAll(cp.Body)
		mdata := Pokedata{}
		err = json.Unmarshal(pnew, &mdata)
		checkerr(err)
		printit(mdata)
	}
	fmt.Println("bye")
}

func printit(data Pokedata) {
	fmt.Println(data.Name, ":")
	var whtspc = `	`
	rang := len(data.Abilities)
	for i := 0; i < rang; i++ {
		fmt.Println(whtspc, "Ability ", i, ": ", data.Abilities[i].Ability.Name)
	}
	fmt.Println(whtspc, "Weight"+": ", data.Weight)
	for q := 0; q < len(data.Types); q++ {
		fmt.Println(whtspc, "Type ", q, ": ", data.Types[q].Type.Name)
	}
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
