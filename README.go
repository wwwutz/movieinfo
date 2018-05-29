// +build ignore

// This program generates README.md. It can be invoked by running
// go generate
package main

import (
	"log"
	"os"
	"os/exec"
	"text/template"
	"time"
)

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func run(cmd string) string {
	c := exec.Command("sh","-c",cmd)
	out, err := c.CombinedOutput()
	die(err)
	return string(out)
}

func runclean(cmd string) string {
	c := exec.Command("sh","-c",cmd)
	out, err := c.CombinedOutput()
	die(err)
	return string(out)
}

func main() {

	funcMap := template.FuncMap{
        "run": run,
        "runclean": runclean,
	}
	var packageTemplate = template.Must(template.New("").Funcs(funcMap).ParseFiles("README.md.template"))

	// Create an *exec.Cmd
	cmd := exec.Command("echo", "Called from Go!")
	output, err := cmd.CombinedOutput()
	log.Print(string(output))

	f, err := os.Create("README.md")
	die(err)
	defer f.Close()

	packageTemplate.ExecuteTemplate(f, "README.md.template", struct {
		Timestamp time.Time
	}{
		Timestamp: time.Now(),
	})
}
