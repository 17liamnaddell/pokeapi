package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	//C means that we arnt writing go-colors.R everytime we want to write in the color red, instead we write C.R. In python this is writen like
	//import blah as b
	//MAKE SURE TO WRITE C IN CAPS
	//To get this package, write to the terminal:go get github.com/skilstak/go-colors
	C "github.com/skilstak/go-colors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
)

//when we get data from Pokeapi, This will ignore all of the other data that is given, just go to pokeapi.co/api/v2/pokemon/tepig to see why we shorten the data
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
	//the go command starts a goroutine. A gorutine allows a programmer to run a block of code while other code is running, functions like this end when the non gorutine code stops running. Otherwise, they operate like normal functions. There is a lot more detail about goroutines later in go programming though. They are what make go so powerfull
	//captureCC blocks the user from typing Control C and forces them to type done, that way they get the bye message
	go captureCC()

	//make a scanner that scans the Stdin (The command line)
	Scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter a pokemon")

		//actually uses the scanner
		Scanner.Scan()

		if Scanner.Text() == "done" {
			break
		}
		//gets data from the webpage, if you visit the link refrenced above on line 18, you will see what data we are getting.
		//KEEP IN MIND: when you get the data, it is in encoded HTML form.
		cp, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + Scanner.Text())

		//Decodes the HTML and gets the raw json data. If you go to the link on line 18, and hit inspect element, all the data is located inside the <body></body> tag. We are just
		//getting the data from that tag
		//underscore is the trash can of golang, if you don't care about a variable, you can put an underscore there, and it goes to nowhere. You don't have to do anything with t		  //he error or item in question
		pnew, _ := ioutil.ReadAll(cp.Body)

		//Creating a new instance of Pokedata, with all empty values
		mdata := Pokedata{}

		//decodes the data from pnew(line 57) and puts it in json format, or otherwise makes it from just plain text that has json format, to computer code. Like turing json code		  //in a .txt file to json in a .json file
		//MAKE ABSOLUTELY SURE TO PUT A = AND NOT AN :=, ERR IS ALREADY DECLARED ON LINE 53
		err = json.Unmarshal(pnew, &mdata)

		//?!?!?!?!?!?!?!!? who cares about the Weight, and why are we checking for a 0?
		//Well, there is no error for the pokemon's name being spelled wrong, or that name not being a pokemon, like xbox1. If there is no value provided for a int in golang, it		  //is automatically 0 If we get no Weight or data, then there must not be a pokemon under that name
		if mdata.Weight == 0 {
			fmt.Println(C.R + "no such pokemon")
			continue
		}
		//checking the error for the error on line 63
		checkerr(err)
		printit(mdata)
	}
	fmt.Println(C.R + "bye")
}

func captureCC() {
	//as long as the non gorutine functions are running,
	for {
		//channels or chans are golang's way of transmitting data, os.Signal is what kind of data the channel is transfering, and 1 is how much data can be stored
		c := make(chan os.Signal, 1)
		//if there is a Ctrl + c, then tell the channel
		signal.Notify(c, os.Interrupt)
		//emojs in golang, not really. They are what we recive from the channel c, in this case, if there was a person pressing ctrl c, then b would be = to ctrl + c (Dont print it, it's in binary form or unicode form)
		b := <-c
		//if there is an interrupt, then print line 91 instead of exiting
		if b == os.Interrupt {
			fmt.Println(C.B + "type done to leave, or Enter a pokemon")
		}
	}
}

//data is of type pokedata
func printit(data Pokedata) {
	//print in red
	fmt.Println(C.R+data.Name, ":")

	//stands for white space
	var whtspc = `	`

	//look at the struct on line 19, it is an array of structs, or []struct. We are finding how many structs are inside of data.Abilites. Same goes for the type
	rang := len(data.Abilities)

	//for every ability
	for i := 0; i < rang; i++ {
		//print in blue, a tab, the ability #, a : and the name of the ability. in golang  if you wanted to access the 1st entry of a array, you would do it like this array[0]
		//in this case we are using i as an indexer, to access the ablity in question that we are iterating over in the for loop
		fmt.Println(C.B+whtspc, "Ability ", i, ": ", data.Abilities[i].Ability.Name)
	}
	//You get the picture
	fmt.Println(C.M+whtspc, "Weight"+": ", data.Weight)
	//same as the for loop above it
	for q := 0; q < len(data.Types); q++ {
		fmt.Println(C.G+whtspc, "Type ", q, ": ", data.Types[q].Type.Name+C.Y)
	}
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
