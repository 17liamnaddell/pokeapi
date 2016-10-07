package main

import (
	"encoding/json"
	"fmt"
	C "github.com/skilstak/go-colors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
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

var home = os.Getenv("HOME")

func checkargs(i int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("You forgot to put in a pokemon/generation")
			os.Exit(0)
		}
	}()
	_ = os.Args[i+1]
}

func main() {
	if _, err := os.Stat(home + "/.pokeapi"); os.IsNotExist(err) == true {
		fmt.Println(err)
		os.Chdir(os.Getenv("HOME"))
		err := os.Mkdir(".pokeapi", 0777)
		fmt.Println(err, "Creating new .pokeapi directory")
	}
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-f" || os.Args[i] == "--find" {
			os.Chdir(home)
			checkargs(i)
			pokedir, err := os.Open(".pokeapi")
			checkerr(err)
			var myint int
			pokedirs, err := pokedir.Readdirnames(myint)
			checkerr(err)
			stored := false
			for l := 0; l < len(pokedirs); l++ {
				if pokedirs[l] == os.Args[i+1] {
					pokemon, _ := ioutil.ReadFile(".pokeapi/" + os.Args[i+1])
					pokedat := Pokedata{}
					json.Unmarshal(pokemon, &pokedat)
					printit(pokedat)
					stored = true
					break
				}

			}
			if stored == false {
				pokedat := GetPokemon(os.Args[i+1])
				if pokedat.Weight == 0 {
					fmt.Println(os.Args[i+1], " is not a pokemon")
					os.Exit(0)
				}
				printit(pokedat)
			}
		} else if os.Args[i] == "-rf" || os.Args[i] == "--removeData" {
			os.Chdir(home + "/.pokeapi")
			os.RemoveAll(".")
		} else if os.Args[i] == "-fa" || os.Args[i] == "--findAll" {
			checkargs(i)
			ListPokemon(os.Args[i+1])
			//get all data from each pokemon and write that to the file

		}
	}

}

var Allpokemon []string

func GetPokemon(name string) Pokedata {
	os.Chdir(home + "/.pokeapi")
	pokelink := "https://pokeapi.co/api/v2/pokemon/" + name
	fmt.Println(pokelink)
	poke, err := http.Get(pokelink)
	fmt.Println(poke.Body)
	checkerr(err)
	pokedat := Pokedata{}
	idontcare, _ := ioutil.ReadAll(poke.Body)
	err = json.Unmarshal(idontcare, &pokedat)
	writeme, _ := json.Marshal(pokedat)
	fmt.Println("writing")
	err1234 := ioutil.WriteFile(name, []byte(writeme), 0777)
	fmt.Println(err1234)
	checkerr(err)
	os.Chdir(home)
	return pokedat
}

func ListLink(URL string) {
	raw, err := http.Get(URL)

	checkerr(err)
	rawJson, _ := ioutil.ReadAll(raw.Body)
	poka := Idkwtth{}
	err = json.Unmarshal(rawJson, &poka)
	for i := 0; i < len(poka.Results); i++ {
		Allpokemon = append(Allpokemon, poka.Results[i].Name)
	}
	log.Print(Allpokemon)
}

type Gen struct {
	Gen   int
	Start int
	End   int
}

var Agen Gen

func Sortgen(gen int) {
	switch gen {
	case 1:
		Agen = Gen{1, 1, 151}
	case 2:
		Agen = Gen{2, 152, 251}
	case 3:
		Agen = Gen{3, 252, 386}
	case 4:
		Agen = Gen{4, 387, 493}
	case 5:
		Agen = Gen{5, 494, 649}
	case 6:
		Agen = Gen{6, 650, 721}
	default:
		fmt.Println("there is no generation ", gen)
		os.Exit(0)
	}
}

func ListPokemon(gen string) {
	os.Chdir(home + "/.pokeapi")
	nen, _ := strconv.Atoi(gen)
	Sortgen(nen)
	fmt.Println(Agen)
	Link := "https://pokeapi.co/api/v2/pokemon"
	for {
		fmt.Println("listn" + Link)
		raw, _ := http.Get(Link)
		rawJson, _ := ioutil.ReadAll(raw.Body)
		fmt.Println(string(rawJson))
		poka := Idkwtth{}
		_ = json.Unmarshal(rawJson, &poka)
		if poka.Next != "" {
			fmt.Println("linkn")
			ListLink(Link)
			Link = poka.Next
		} else {
			broke := false
			for i := Agen.Start; i < len(Allpokemon); i++ {
				if i > Agen.End {
					broke = true
					fmt.Println("broke")
					break
				}
				fmt.Println("count: ", i)
				GetPokemon(Allpokemon[i])
			}
			if broke == true {
				break
			}
		}
	}
}

type Prints struct {
	Type  string
	AT    int
	Thing string
}

var whtspc = `	`
var MyTempl, _ = template.New("Our template demo").Parse(C.G + whtspc + "{{.Type}} {{.AT}}: {{.Thing}}" + C.Y + "\n")

func printit(data Pokedata) {
	fmt.Println(C.R+data.Name, ":")
	rang := len(data.Abilities)

	for i := 0; i < rang; i++ {
		realdata := Prints{"Ability", i, data.Abilities[i].Ability.Name}
		_ = MyTempl.Execute(os.Stdout, realdata)
	}
	fmt.Println(C.M+whtspc, "Weight"+": ", data.Weight)
	for q := 0; q < len(data.Types); q++ {
		realdata := Prints{Type: "Type", AT: q, Thing: data.Types[q].Type.Name}
		_ = MyTempl.Execute(os.Stdout, realdata)

	}
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
