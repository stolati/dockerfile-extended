package docker_template

import (
	"../values_context"
	"io/ioutil"
	"text/template"
	"log"
	"bytes"
	"path"
)

func ApplyTemplate(dockerFile string, ctx values_context.MainCtx, debug bool) (output string, err error) {

	content, readFileErr := ioutil.ReadFile(dockerFile)
	if readFileErr != nil {
		return "", readFileErr
	}

	content_str := string(content[:])
	tmpl, tmplErr := template.New(path.Base(dockerFile)).Parse(content_str)
	if tmplErr != nil {
		return "", tmplErr
	}

	buf := new(bytes.Buffer)
	tmplExecErr := tmpl.Execute(buf, ctx)
	if tmplExecErr != nil {
		return "", tmplExecErr
	}

	if debug {
		log.Println("Dockerfile to execute :")
		println(buf.String())
	}

	return buf.String(), nil
}
