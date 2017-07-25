package main

import (
	"flag"
	"log"
	"path/filepath"
	"../values_context"
	"../docker_run"
	"../docker_template"
	"../docker_parser"
)

func main() {

	dockerfilePtr := flag.String("dockerfile", "./Dockerfile.child.ext", "the dockerfile to process")
	debugPtr := flag.Bool("debug", false, "Show the dockerfile generated")
	dryRunPtr := flag.Bool("dry-run", false, "Don't really build the dockerfile")

	flag.Parse()
	bypassArgs := flag.Args()

	err := LaunchDocker(*dockerfilePtr, "", bypassArgs, *debugPtr, *dryRunPtr)
	if err != nil {
		log.Fatal(err)
	}

	println("Success")

}

func LaunchDocker(dockerfile string, needToBeTag string, bypassArgs []string, debug bool, dryRun bool)(err error){

	dirContext := filepath.Dir(dockerfile)
	templateCtx := values_context.GetContext(dirContext, debug)

	templateOutput, tmplErr := docker_template.ApplyTemplate(dockerfile, templateCtx, debug)
	if tmplErr != nil {
		return tmplErr
	}

	parserRes, parserErr := docker_parser.Parse(templateOutput, debug)
	if parserErr != nil {
		log.Fatal(parserErr)
	}

	if parserRes.FromFile != "" {
		// Here we need another file to be FROM
		parentDockerfile := filepath.Clean(filepath.Join(dirContext, parserRes.FromFile))

		launchDockerErr := LaunchDocker(parentDockerfile, parserRes.TmpTag, bypassArgs, debug, dryRun)
		if launchDockerErr != nil {
			return launchDockerErr
		}
		defer docker_run.CleanTag(parserRes.TmpTag, dryRun, debug)
	}

	buildErr := docker_run.BuildDocker(parserRes, dirContext, bypassArgs, needToBeTag, dryRun, debug)
	if buildErr != nil {
		log.Fatal(buildErr)
	}

	return nil
}

