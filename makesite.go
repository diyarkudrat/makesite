package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
)

type post struct {
	User    string
	Content string
}

func main() {

}

func renderTemplate(content string) string {
	paths := []string{
		"templates/template.tmpl",
	}

	temp := new(bytes.Buffer)
	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))

	err := t.Execute(temp, post{User: "Diyar", Content: content})
	if err != nil {
		panic(err)
	}

	return temp.String()
}

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(fileContents))
	return string(fileContents)
}

func WriteFile(tmpl []byte, file string) {
	if err := ioutil.WriteFile("exports/"+file, tmpl, 0666); err != nil {
		panic(err)
	}
}
