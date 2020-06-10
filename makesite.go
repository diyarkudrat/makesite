package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"html/template"
)

type post struct {
	User string
	Content string
}

func main() {


}

func renderTemplate(content string) string {
	paths := []string{
		"template.tmpl",
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
