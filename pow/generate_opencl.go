// +build ignore


package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"text/template"
)

const (
	sourceDir = "opencl_sources"
	resultFile = "kernelSource.go"
	headerFile = "header.c"
	kernelFile = "kernel.cl"
)

// This program compiles all the opencl code into a single string
// which is then baked into the final go executable


func compileSources() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("couldn't find opencl source directory")
	}

	dir := filepath.Dir(filename) + "/" + sourceDir
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", fmt.Errorf("couldn't read directory %s: %s", dir, err.Error())
	}

	contents := make(map[string][]byte)
	for _, file := range files {
		filepath := dir + "/" + file.Name()
		data, err := ioutil.ReadFile(filepath)
		if err != nil {
			return "", fmt.Errorf("couldn't read file %s: %s", filepath, err.Error())
		}
		contents[file.Name()] = data
	}
	var out strings.Builder
	if _, ok := contents[headerFile]; !ok {
		return "", fmt.Errorf("missing %s file", headerFile)
	}
	out.Write(contents[headerFile])
	delete(contents, headerFile)

	mainSrc, ok := contents[kernelFile]
	if !ok {
		return "", fmt.Errorf("missing %s file", kernelFile)
	}
	delete(contents, kernelFile)

	for _,data := range contents {
		out.Write(data)
	}
	out.Write(mainSrc)

	return out.String(), nil
}

func main() {
	kernelStr, err := compileSources()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to compile opencl sources: %s\n", err.Error())
		return
	}

	outFile, err := os.Create(resultFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create output file: %s: %s\n", resultFile, err.Error())
	}
	defer outFile.Close()

	packageTemplate.Execute(outFile, struct {
		Timestamp time.Time
		Directory string
		Source    string
	}{
		Timestamp: time.Now(),
		Directory: sourceDir,
		Source:    kernelStr,
	})

}

var packageTemplate = template.Must(template.New("").Parse(`
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .Timestamp }}
// using data from {{ .Directory }}
package pow

const KernelSource = `+ "`" + `
{{ .Source }}
` + "`"))
