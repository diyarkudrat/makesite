package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type post struct {
	User    string
	Content string
}

func main() {

	fileFlag := flag.String("file", "first-post.txt", "define input text")
	flag.Parse()
	var fileName string = *fileFlag
	fileName = fileName[0:strings.Index(*fileFlag, ".")] + ".html"

	var fileData string = readFile(*fileFlag)
	renderTemplate("template.tmpl", fileData, fileName)

}

func renderTemplate(tPath, textData, fileName string) {
	paths := []string{
		tPath,
	}

	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	t, err := template.New(tPath).ParseFiles(paths...)
	if err != nil {
		panic(err)
	}

	originName := fileName[0:strings.Index(fileName, ".")]

	err = t.Execute(f, post{textData, originName})
	if err != nil {
		panic(err)
	}

	f.Close()

}

func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fileContents))
	return string(fileContents)
}
