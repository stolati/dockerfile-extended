package docker_template

import (
	"../values_context"
	"io/ioutil"
	"text/template"
	"bytes"
	"path"
	"github.com/Masterminds/sprig"
	"fmt"
)

func ApplyTemplate(dockerFile string, ctx values_context.MainCtx, debug bool) (output string, err error) {

	content, readFileErr := ioutil.ReadFile(dockerFile)
	if readFileErr != nil {
		return "", readFileErr
	}

	content_str := string(content[:])
	tmpl, tmplErr := template.New(path.Base(dockerFile)).Funcs(sprig.TxtFuncMap()).Parse(content_str)
	if tmplErr != nil {
		return "", tmplErr
	}

	buf := new(bytes.Buffer)
	tmplExecErr := tmpl.Execute(buf, ctx)
	if tmplExecErr != nil {
		return "", tmplExecErr
	}

	if debug {
		fmt.Println("#####################")
		fmt.Println("Template output :")
		fmt.Println("#####################")
		fmt.Println(buf.String())
		fmt.Println()
	}

	return buf.String(), nil
}
