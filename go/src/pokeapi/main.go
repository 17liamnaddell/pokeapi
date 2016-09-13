package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	//C means that we arnt writing go-colors.R everytime we want to write in the color red, instead we write C.R. In python this is writen like
	//import blah as b
	//MAKE SURE TO WRITE C IN CAPS
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
	Scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter a pokemon")
		Scanner.Scan()
		if Scanner.Text() == "done" {
			break
		}
		cp, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + Scanner.Text())

		pnew, _ := ioutil.ReadAll(cp.Body)
		mdata := Pokedata{}
		err = json.Unmarshal(pnew, &mdata)
		if mdata.Weight == 0 {
			fmt.Println(C.R + "no such pokemon")
			continue
		}
		checkerr(err)
		printit(mdata)
	}
	fmt.Println(C.R + "bye")
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
