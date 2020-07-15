package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"

	"github.com/bregydoc/gtranslate"
	"golang.org/x/text/language"
	"gopkg.in/russross/blackfriday.v2"
)

type pageData struct {
	Content string
	Title   string
}

func main() {

	fileFlag := flag.String("file", "first-post.txt", "define input text")
	dirFlag := flag.String("directory", "none", "generates all .txt files in directory")
	outputDirFlag := flag.String("output", "templates/", "Generator output directory")
	flag.Parse()

	if *dirFlag == "none" {
		runFile(*fileFlag, "txt_dir/")
	} else {
		runDir(*dirFlag, *outputDirFlag)
	}

}

func runFile(fileFlag, directory string) {

	var fileName string = fileFlag

	if fileName[strings.Index(fileFlag, "."):len(fileFlag)] != ".txt" {
		return
	}

	if strings.Contains(strings.ToLower(fileFlag), ".md") {

		var data string = readFile(directory + fileFlag)
		tmpl := renderTemplate("template.tmpl", data, fileName)
		output := blackfriday.Run(tmpl)
		ioutil.WriteFile(output, fileFlag)

		return
	}

	fileName = fileName[0:strings.Index(fileFlag, ".")] + ".html"

	var data string = readFile(directory + fileFlag)
	renderTemplate("template.tmpl", data, fileName)
}

func runDir(directory, output string) {

	if directory[len(directory)-1] != "/"[0] {
		directory += "/"
	}

	files, err := ioutil.ReadDir(directory)

	if err != nil {
		panic(err)
	}

	for _, file := range files {

		if file.IsDir() == false {
			runFile(file.Name(), directory)
		} else {
			runDir(directory+"/"+file.Name(), output)
		}
	}
}

func renderTemplate(tPath, textData, fileName string) {
	paths := []string{
		tPath,
	}

	f, err := os.Create("templates/" + fileName)
	if err != nil {
		panic(err)
	}

	t, err := template.New(tPath).ParseFiles(paths...)
	if err != nil {
		panic(err)
	}

	originName := fileName[0:strings.Index(fileName, ".")]

	txtTranslated := translateText(textData)

	err = t.Execute(f, pageData{txtTranslated, originName})
	if err != nil {
		panic(err)
	}

	f.Close()

}

func translateText(txtData string) string {

	translated, err := gtranslate.Translate(txtData, language.English, language.Spanish)
	if err != nil {
		panic(err)
	}

	// fmt.Printf("en: %s | ja: %s \n", txtData, translated)
	// en: Hello World | ja: こんにちは世界
	return string(translated)
}

func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	// fmt.Println(string(fileContents))
	return string(fileContents)
}
