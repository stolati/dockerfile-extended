package docker_run

import (
	"os/exec"
	"os"
	"io/ioutil"
	"../docker_parser"
	"strings"
	"path"
	"fmt"
)

// TODO --iidfile can be handy

func BuildDocker(
	dockerInfo docker_parser.Parser,
	contextPath string,
	otherArguments []string,
	needToBeTag string,
	dryRun bool,
	debug bool) (err error) {
	binary, lookErr := exec.LookPath("docker")
	if lookErr != nil {
		return lookErr
	}

	if strings.ToUpper(dockerInfo.ContextPath) == "NONE" {
		if dryRun {
			contextPath = "<tempdir>"
		} else {
			tmpDir, tmpDirErr := ioutil.TempDir("", "docker_extended.")
			if tmpDirErr != nil {
				return tmpDirErr
			}
			defer os.RemoveAll(tmpDir)
			contextPath = tmpDir

		}
	} else if dockerInfo.ContextPath != "" {
		contextPath = path.Clean(path.Join(contextPath, dockerInfo.ContextPath))
	}

	dockerTmpName := "<tempfile>"

	if ! dryRun {
		dockerTmp, tmpErr := ioutil.TempFile(contextPath, "docker_extended.")
		if tmpErr != nil {
			return tmpErr
		}
		dockerTmpName = dockerTmp.Name()
		defer os.Remove(dockerTmpName)
		defer dockerTmp.Close()
	}

	// Building the args
	args := []string{
		"build",
	}

	args = append(args, otherArguments...)

	for _, tag := range dockerInfo.Tags {
		args = append(args, "--tag", tag)
	}
	if needToBeTag != "" {
		args = append(args, "--tag", needToBeTag)
	}

	args = append(args, "--file", dockerTmpName, contextPath)

	if debug {
		fmt.Println("#####################")
		fmt.Println("Docker command :")
		fmt.Println("#####################")
		fmt.Println("cd ", contextPath)
		fmt.Println("docker", strings.Join(args, " "))
		fmt.Println("rm", dockerTmpName)
		fmt.Println()
	}

	if dryRun {
		return nil
	}

	ioutil.WriteFile(dockerTmpName, []byte(dockerInfo.DockerfileContent), 0777)

	cmd := exec.Command(binary, args...)
	cmd.Stdout, cmd.Stderr, cmd.Env = os.Stdout, os.Stderr, os.Environ()
	startErr := cmd.Start()
	if startErr != nil {
		return startErr
	}

	waitErr := cmd.Wait()
	if waitErr != nil {
		return waitErr
	}

	return nil
}

func CleanTag(tag string, dryRun bool, debug bool)(err error){

	args := []string{
		"images",
		"rm",
		tag,
	}

	if debug {
		fmt.Println("#####################")
		fmt.Println("Docker removing tag cmd :")
		fmt.Println("#####################")
		fmt.Println("docker", strings.Join(args, " "))
		fmt.Println()
	}

	if dryRun {
		return nil
	}

	cmd := exec.Command("docker", args...)
	cmd.Stdout, cmd.Stderr, cmd.Env = os.Stdout, os.Stderr, os.Environ()
	startErr := cmd.Start()
	if startErr != nil {
		return startErr
	}

	waitErr := cmd.Wait()
	if waitErr != nil {
		return waitErr
	}

	return nil
}