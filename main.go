package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
)

type Component struct {
	Configuration Configuration `xml:"configuration"`
}

type Configuration struct {
	Type      string `xml:"type,attr"`
	Name      string `xml:"name,attr"`
	Envs      Envs   `xml:"envs"`
	FilePath  Value  `xml:"filePath"`
	Directory Value  `xml:"directory"`
	Kind      Value  `xml:"kind"`
}

type Envs struct {
	XMLName xml.Name `xml:"envs"`
	Envs    []Env    `xml:"env"`
}

type Env struct {
	XMLName xml.Name `xml:"env"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:"value,attr"`
}

type Value struct {
	Value string `xml:"value,attr"`
}

func main() {
	runConfigurationsDir := "/.idea/runConfigurations"
	var files []string

	filepath.Walk(runConfigurationsDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, "xml") {
			files = append(files, path)
		}

		return nil
	})

	var components []Component
	for _, file := range files {
		component := parse(file)
		components = append(components, component)
	}

	component := getConfiguration(components)

	var str strings.Builder
	str.WriteString("export PROJECT_DIR=$(pwd)\n")

	for _, env := range component.Configuration.Envs.Envs {
		str.WriteString("export " + env.Name + "=\"" + env.Value + "\"\n")
	}

	str.WriteString("\n")

	var path string
	if component.Configuration.Kind.Value == "FILE" {
		path = component.Configuration.FilePath.Value
	} else {
		path = component.Configuration.Directory.Value
	}

	if component.Configuration.Type == "GoApplicationRunConfiguration" {
		str.WriteString("go run " + path)
		runString(str, "./run.sh")
	}
	if component.Configuration.Type == "GoTestRunConfiguration" {
		str.WriteString("go test " + path + " -v -count=1")
		runString(str, "./test.sh")
	}
}

func runString(str strings.Builder, filename string) {
	contents := str.String()
	contents = strings.Replace(contents, "$PROJECT_DIR$", "$PROJECT_DIR", -1)

	d1 := []byte(contents)
	ioutil.WriteFile(filename, d1, 0777)

	cmd := exec.Command("/bin/sh", filename)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func getConfiguration(components []Component) Component {
	var files []string
	for _, component := range components {
		files = append(files, component.Configuration.Name)
	}
	prompt := promptui.Select{
		Label: "Select Configuration",
		Items: files,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return components[i]
}

func parse(file string) Component {
	xmlFile, err := os.Open(file)
	defer xmlFile.Close()

	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var component Component
	xml.Unmarshal(byteValue, &component)

	return component
}
