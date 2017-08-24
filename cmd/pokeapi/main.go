package main

import "github.com/liamnaddell/pokeapi"
import "github.com/mattn/go-gtk/glib"
import "github.com/mattn/go-gtk/gtk"
import "strconv"
import "os"
import "github.com/urfave/cli"

var Version string

//label.SetSizeRequest(20, 20)
type Types struct {
	Type struct {
		Name string `json:"name"`
	} `json:"type"`
}

func main() {
	app := cli.NewApp()
	app.Name = "pokeapi"
	app.Usage = "graphical tool for querying the PokeApi database"
	app.Action = start
	app.Version = Version
	app.Run(os.Args)
}

func start(c *cli.Context) error {
	pokemon, err := pokeapi.StartGetPokemon("pikachu")
	if err != nil {
		cli.NewExitError(err, 1)
	}
	//fmt.Println(pokemon.Types[0].Type.Name)
	window := basicGtk()

	//vbox/label
	vbox := gtk.NewVBox(false, 1)
	Tsep := gtk.NewHSeparator()
	headerLabel := newHeader(pokemon.Name)
	vbox.PackStart(headerLabel, false, true, 0)
	vbox.PackStart(Tsep, false, true, 0)

	//boxes
	hbox := gtk.NewHBox(false, 1)

	vbox.PackStart(hbox, false, true, 0)
	//label 2
	label2 := labelWMkup("<span size=\"10000\"> type:<b>  " + pokemon.Types[0].Type.Name + "  </b></span>")
	hbox.PackStart(label2, false, true, 0)
	sep := gtk.NewVSeparator()
	hbox.PackStart(sep, false, true, 0)

	//bottom separator
	sep2 := gtk.NewHSeparator()
	//vbox.PackStart(sep2, false, true, 2)
	vbox.PackStart(sep2, false, true, 0)

	//label 3
	str := strconv.Itoa(pokemon.Id)
	label3 := labelWMkup("<span size=\"15000\"> id:<b> " + str + " </b></span>")
	sep3 := gtk.NewVSeparator()
	hbox.PackEnd(label3, false, true, 0)
	hbox.PackEnd(sep3, false, true, 0)

	//hbox2
	hbox2 := gtk.NewHBox(true, 3)
	vbox.PackStart(hbox2, true, true, 0)

	//vboxleft
	vboxleft := gtk.NewVBox(false, 0)
	hbox2.PackStart(vboxleft, true, true, 0)
	button := gtk.NewButtonWithLabel("Who's That Pokemon!")
	vboxleft.PackEnd(button, false, false, 0)
	entry := gtk.NewEntry()
	vboxleft.PackEnd(entry, false, false, 0)

	//button
	button.Clicked(func() {
		//reset all in box
		NewPokemon, err := pokeapi.StartGetPokemon(entry.GetText())
		if err != nil {
			button.SetLabel("Could Not find that pokemon")
			return
		}

		//reset header
		headerLabel.SetMarkup("<span foreground=\"blue\" size=\"20000\"> <b>" + NewPokemon.Name + "</b></span>")

		//reset type
		label2.SetMarkup("<span size=\"10000\"> type:<b>  " + NewPokemon.Types[0].Type.Name + "  </b></span>")

		//id stuff
		str2 := strconv.Itoa(NewPokemon.Id)
		label3.SetMarkup("<span size=\"15000\"> id:<b> " + str2 + " </b></span>")
		button.SetLabel("Who's That Pokemon!")
	})
	//sep

	//vboxright

	//mid sep

	//vbox.Add(sep)
	//show all
	vboxleft.ShowAll()
	hbox2.ShowAll()
	hbox.ShowAll()
	vbox.ShowAll()
	window.Add(vbox)
	window.ShowAll()
	gtk.Main()

	return nil
}

func labelWMkup(markup string) *gtk.Label {
	label := gtk.NewLabel("")
	label.SetMarkup(markup)
	return label
}

func basicGtk() *gtk.Window {
	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetSizeRequest(300, 400)
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		gtk.MainQuit()
	}, "foo")
	return window
}

func newHeader(text string) *gtk.Label {
	label := gtk.NewLabel("")
	label.SetMarkup("<span foreground=\"blue\" size=\"20000\"> <b>" + text + "</b></span>")
	return label
}
