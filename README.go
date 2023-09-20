//go:build ignore
// +build ignore

// This program generates README.md. It can be invoked by running
// go generate
package main

import (
	"log"
	"os"
	"os/exec"
	"text/template"
	// "time"
	"fmt"
	"runtime"
	"strings"
)

func die(err error) {
	r := "panic exit."
	if err == nil {
		return
	}
	pc, file, no, ok := runtime.Caller(10)
	details := runtime.FuncForPC(pc)
	d := ""
	if details != nil {
		d = " in " + details.Name() + "()"
	}
	if ok {
		r = fmt.Sprintf("// died %s with err=%v\n// %s#%d", d, err, file, no)
	} else {
		r = fmt.Sprintf("// died %s with err=%v", d, err)
	}

	dumpstack()

	log.Fatal(r)
}

func dumpstack() {
	pc := make([]uintptr, 30)
	n := runtime.Callers(0, pc)
	if n == 0 {
		// No pcs available. Stop now.
		// This can happen if the first argument to runtime.Callers is large.
		return
	}

	pc = pc[:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)

	// Loop to get frames.
	// A fixed number of pcs can expand to an indefinite number of Frames.
	n = 1
	for {
		frame, more := frames.Next()
		// To keep this example's output stable
		// even if there are changes in the testing package,
		// stop unwinding when we leave package runtime.
		fmt.Printf("[%2d] : %s in %s line %d\n", n, frame.Function, frame.File, frame.Line)
		if strings.HasPrefix(frame.Function, "main.main") {
			break
		}
		n += 1
		if !more {
			break
		}
	}
}

func run(cmd string) string {
	fmt.Printf("---> run(" + cmd + ")\n")
	cwd, err := os.Getwd()
	die(err)
	defer os.Chdir(cwd)
	die(os.Chdir(os.TempDir()))
	tmpdir := "movieinfo.tmp"
	_ = os.Mkdir(tmpdir, 0755)
	die(os.Chdir(tmpdir))
	c := exec.Command("bash", "-c", cmd)
	out, err := c.CombinedOutput()
	die(err)
	return string(out)
}

func runclean(cmd string) string {
	fmt.Printf("---> runclean(" + cmd + ")\n")
	cwd, err := os.Getwd()
	die(err)
	defer os.Chdir(cwd)
	die(os.Chdir(os.TempDir()))
	tmpdir := "movieinfo.tmp"
	die(os.RemoveAll(tmpdir))
	_ = os.Mkdir(tmpdir, 0755)
	die(os.Chdir(tmpdir))
	c := exec.Command("bash", "-c", cmd)
	out, err := c.CombinedOutput()
	die(err)
	return string(out)
}

func main() {

	cwd, err := os.Getwd()
	err = os.Setenv("PATH", cwd+":"+os.Getenv("PATH"))

	funcMap := template.FuncMap{
		"run":      run,
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

	packageTemplate.ExecuteTemplate(f, "README.md.template", nil)
}
