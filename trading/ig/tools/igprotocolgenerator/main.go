package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func main() {
	fmt.Println("Start...")
	for _, webRequestData := range data {
		
		fmt.Printf("Busy with %s %s, %d\n", webRequestData.Command, webRequestData.Method, webRequestData.Version)
		
		tmpl := template.New("").Delims("<<", ">>").Funcs(template.FuncMap{
			"CreateInputParam":      webRequestData.CreateInputParam,
			"CreateSessionParam":    webRequestData.CreateSessionParam,
			"AssignSessionParam":    webRequestData.AssignSessionParam,
			"CreateInputParamCode":  webRequestData.CreateInputParamCode,
			"AssignRequestBody":     webRequestData.AssignRequestBody,
			"AssignRequestLength":   webRequestData.AssignRequestLength,
			"UseInputParam":         webRequestData.UseInputParam,
			"CloseResponseBody":     webRequestData.CloseResponseBody,
			"KeepResponseBodyAlive": webRequestData.KeepResponseBodyAlive,
			"UseInputParamBytes":    webRequestData.UseInputParamBytes,
			"CreateReaderParam":     webRequestData.CreateReaderParam,
			"URLPostfix":            webRequestData.URLPostfix,
		})
		tmpl = template.Must(tmpl.ParseFiles("Template.temp"))

		filename := fmt.Sprintf(
			"../../igprotocol/%s_%s_%d.go",
			webRequestData.Command,
			webRequestData.Method,
			webRequestData.Version)

		oFile, err := os.Create(filename)
		if err != nil {
			panic(err)
		}

		if err = tmpl.ExecuteTemplate(oFile, "Template.temp", webRequestData); err != nil {
			panic(err)
		}
		oFile.Close()

		rungofmt(filename, true)
	}
	fmt.Println("Finish...")
}

func rungofmt(path string, fiximports bool) error {
	args := []string{"-w", path}
	output, err := exec.Command("gofmt", args...).CombinedOutput()

	if fiximports && err == nil {
		args = []string{"-w", path}
		output, err = exec.Command("goimports", args...).CombinedOutput()
	}

	if err != nil {
		fmt.Println("Error executing gofmt", strings.Join(args, " "))
		os.Stdout.Write(output)
	}

	return err
}
