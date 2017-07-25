package main

import (
	"flag"
	"log"
	"path/filepath"
	"../values_context"
	"../docker_run"
	"../docker_template"
	"../docker_parser"
	"fmt"
	"strings"
)

func main() {

	dockerfilePtr := flag.String("dockerfile", "", "the dockerfile to process")
	dockersDirectoryPtr := flag.String("dir", ".", "the directory to where search for dockerfiles")
	debugPtr := flag.Bool("debug", false, "Show the dockerfile generated")
	dryRunPtr := flag.Bool("dry-run", false, "Don't really build the dockerfile")

	flag.Parse()
	bypassArgs := flag.Args()

	if *dockerfilePtr != "" {

		err := LaunchDocker(*dockerfilePtr, "", bypassArgs, *debugPtr, *dryRunPtr)
		if err != nil {
			log.Fatal(err)
		}

		println("Success")
		return
	}

	// Instead of one file, we have a directory with a bunch of dockerfiles in it
	globDirFiles, globDirErr := filepath.Glob(*dockersDirectoryPtr + "/**/Dockerfile.*.ext")
	if globDirErr != nil {
		log.Fatal(globDirErr)
	}
	globRootFiles, globRootErr := filepath.Glob(*dockersDirectoryPtr + "/Dockerfile.*.ext")
	if globRootErr != nil {
		log.Fatal(globRootErr)
	}

	filenames := append(globDirFiles, globRootFiles...)
	if len(filenames) == 0 {
		absPath, _ := filepath.Abs(*dockersDirectoryPtr)
		log.Fatal("No dockerfile Dockerfile.*.ext found in " + absPath)
		return
	}

	for _, dockerfile := range filenames {
		launchErr := LaunchDocker(dockerfile, "", bypassArgs, *debugPtr, *dryRunPtr)
		if launchErr != nil {
			log.Fatal(launchErr)
		}

	}

	print(len(filenames))

}

func LaunchDocker(dockerfile string, needToBeTag string, bypassArgs []string, debug bool, dryRun bool) (err error) {

	if debug {
		text := "** Running dockerfile : " + dockerfile + " **"
		header := strings.Repeat("*", len(text))

		fmt.Println(header)
		fmt.Println(header)
		fmt.Println(text)
		fmt.Println(header)
		fmt.Println(header)
		fmt.Println()
	}

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
