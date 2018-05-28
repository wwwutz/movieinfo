// +build ignore

// This program generates README.md. It can be invoked by running
// go generate
package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var packageTemplate = template.Must(template.New("").ParseFiles("README.md.template"))
//var packageTemplate = template.Must(template.New("").Parse(`BOO {{ .Timestamp }}`))


func main() {

	carls := []string{}

	f, err := os.Create("README.md")
	die(err)
	defer f.Close()

	packageTemplate.ExecuteTemplate(f, "README.md.template",struct {
		Timestamp time.Time
		URL       string
		Carls     []string
	}{
		Timestamp: time.Now(),
		Carls:     carls,
	})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
