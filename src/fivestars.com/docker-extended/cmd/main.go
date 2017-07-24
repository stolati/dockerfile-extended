package main

import (
	"flag"
	"log"
	"path/filepath"
	"../values_context"
	"../docker_run"
	"../docker_template"
)

func main() {

	dockerfilePtr := flag.String("dockerfile", "./Dockerfile.1.ext", "the dockerfile to process")
	debugPtr := flag.Bool("debug", false, "Show the dockerfile generated")
	dryRunPtr := flag.Bool("dry-run", false, "Don't really build the dockerfile")

	flag.Parse()
	dir_context := filepath.Dir(*dockerfilePtr)

	templateCtx := values_context.GetContext(dir_context, *debugPtr)

	templateOutput, tmplErr := docker_template.ApplyTemplate(*dockerfilePtr, templateCtx, *debugPtr)
	if tmplErr != nil {
		log.Fatal(tmplErr)
	}

	buildErr := docker_run.BuildDocker(templateOutput, dir_context, *dryRunPtr)
	if buildErr != nil {
		log.Fatal(buildErr)
	}
	println("Success")

}
