package main

import "github.com/liamnaddell/pokeapi"
import "os"
import "github.com/mattn/go-gtk/glib"
import "github.com/mattn/go-gtk/gtk"

//import "github.com/mattn/go-gtk/gdk"

//label.SetSizeRequest(20, 20)

func main() {
	pokemon := pokeapi.StartGetPokemon(os.Args[1])
	window := basicGtk()
	vbox := gtk.NewVBox(false, 1)
	label := newHeader(pokemon.Name)
	vbox.PackStart(label, false, false, 3)

	sep := gtk.NewVSeparator()
	vbox.Add(sep)
	//show all
	vbox.ShowAll()
	window.Add(vbox)
	window.ShowAll()
	gtk.Main()

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
