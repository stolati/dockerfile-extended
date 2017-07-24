package docker_run

import (
	"io/ioutil"
	"os/exec"
	"os"
)

func BuildDocker(dockerFileContent string, contextPath string, dryRun bool) (err error) {
	binary, lookErr := exec.LookPath("docker")
	if lookErr != nil {
		return lookErr
	}

	if dryRun {
		return nil
	}

	docker_tmp, tmpErr := ioutil.TempFile(contextPath, "docker_extended.")
	if tmpErr != nil {
		return tmpErr
	}
	defer os.Remove(docker_tmp.Name())
	defer docker_tmp.Close()

	ioutil.WriteFile(docker_tmp.Name(), []byte(dockerFileContent), 0777)

	cmd := exec.Command(binary, "build", "--no-cache", "-f", docker_tmp.Name(), ".")
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
