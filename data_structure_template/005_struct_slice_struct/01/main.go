package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// type items struct {
// 	Wisdom    []sage
// 	Transport []car
// }

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}
func main() {
	b := sage{
		Name:  "Buddah",
		Motto: "buddah bio",
	}
	g := sage{
		Name:  "gandi",
		Motto: "gandi bio",
	}
	m := sage{
		Name:  "mlk",
		Motto: "mlk bio",
	}

	f := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        2,
	}
	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages := []sage{b, g, m}
	cars := []car{f, c}

	// data := items{
	// 	Wisdom:    sages,
	// 	Transport: cars,
	// }

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
