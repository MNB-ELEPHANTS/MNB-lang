package main

import (
	"io/ioutil"
	"log"
	"mnb/mnb-lang/lexer"
	"mnb/mnb-lang/parser"
	"mnb/mnb-lang/translator"
	"os"
	"os/exec"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	p := parser.New()
	l := lexer.New(p)
	t := translator.New(l)
	result := t.Translate(string(content))

	goFileName := strings.Split(os.Args[1], ".")[0] + ".go"

	writeFile(result, goFileName)
	GoCompile(goFileName)
}

func writeFile(content, filename string) {
	err := ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func GoCompile(filename string) {
	cmd := exec.Command("go", "build", filename)
	cmd.Run()
}
