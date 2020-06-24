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
	dirFlag := flag.String("directory", "None", "generates all .txt files in directory")
	outputDirFlag := flag.String("output", "templates/", "Generator output directory")
	flag.Parse()

	if *dirFlag == "None" {
		runForFile(*fileFlag, "txt_dir/")
	}

}

func runFile(fileFlag, directory string) {

	var fileName string = fileFlag

	if fileName[strings.Index(fileFlag, "."):len(fileFlag)] != ".txt" {
		return
	}

	fileName = fileName[0:strings.Index(fileFlag, ".")] + ".html"

	var data string = readFile(directory + fileFlag)
	renderTemplate("template.tmpl", data, fileName)
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
