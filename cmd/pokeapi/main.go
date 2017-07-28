package main

import "github.com/liamnaddell/pokeapi"
import "os"
import "github.com/mattn/go-gtk/glib"
import "github.com/mattn/go-gtk/gtk"

//import "fmt"
import "strconv"

//import "github.com/mattn/go-gtk/gdk"

//label.SetSizeRequest(20, 20)

func main() {
	pokemon := pokeapi.StartGetPokemon(os.Args[1])
	//fmt.Println(pokemon.Types[0].Type.Name)
	window := basicGtk()

	//vbox/label
	vbox := gtk.NewVBox(false, 1)
	Tsep := gtk.NewHSeparator()
	label := newHeader(pokemon.Name)
	vbox.PackStart(label, false, true, 0)
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

	//vbox.Add(sep)
	//show all
	hbox.ShowAll()
	vbox.ShowAll()
	window.Add(vbox)
	window.ShowAll()
	gtk.Main()

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
