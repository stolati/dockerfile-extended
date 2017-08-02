package values_context

import (
	"os"
	"strings"
	"os/exec"
	"strconv"
	"path/filepath"
	"fmt"
	"os/user"
	"runtime"
)

type Ctx map[string]string
type MainCtx map[string]map[string]string

func envContext() (res Ctx) {
	res = make(Ctx)
	for _, e := range os.Environ() {
		z := strings.SplitN(e, "=", 2)
		if z[0] != "" {
			res[z[0]] = z[1]
		}
	}
	return
}

func gitCall(dir string, arg ... string) (res string, err error) {

	gitBinary, lookErr := exec.LookPath("git")
	if lookErr != nil {
		print("git lookup error")
		return "", lookErr
	}

	cmd := exec.Command(gitBinary, arg...)
	cmd.Dir = dir
	cmdOut, cmdErr := cmd.Output()
	if cmdErr != nil {
		return "", cmdErr
	}

	return strings.TrimSpace(string(cmdOut[:])), nil
}

func gitContext(dir string) (res Ctx) {
	res = make(Ctx)

	hash, hashErr := gitCall(dir, "rev-parse", "HEAD")
	branch, branchErr := gitCall(dir, "rev-parse", "--abbrev-ref", "HEAD")
	porcelain, porcelainErr := gitCall(dir, "status", "--porcelain")
	projectPath, projectPathErr := gitCall(dir, "rev-parse", "--show-toplevel")
	remoteUrl, _ := gitCall(dir, "config", "--get", "remote.origin.url")
	if hashErr != nil || branchErr != nil || porcelainErr != nil || projectPathErr != nil {
		return // If not a git repository, return nothing
	}

	urlSplitted := strings.Split(remoteUrl, "/")
	projectNameGit := urlSplitted[len(urlSplitted)-1]
	projectName := strings.Replace(projectNameGit, ".git", "", -1)

	res["HASH_FULL"] = hash
	res["HASH_10"] = hash[:10]
	res["BRANCH"] = branch
	res["IS_MASTER"] = strconv.FormatBool(branch == "master")
	res["IS_STAGING"] = strconv.FormatBool(branch == "staging")
	res["IS_PORCELAIN"] = strconv.FormatBool(porcelain == "")
	res["PROJECT_PATH"] = projectPath
	res["PROJECT_NAME"] = projectName
	return
}

func localContext(dir string) (res Ctx) {
	res = make(Ctx)
	hostname, hostErr := os.Hostname()
	if hostErr == nil {
		res["HOSTNAME"] = hostname
	}
	cwdDir, cwdErr := os.Getwd()
	if cwdErr == nil {
		res["RUN_CWD"] = cwdDir
	}
	dockerDir, dockerDirErr := filepath.Abs(dir)
	if dockerDirErr == nil {
		res["DOCKER_CWD"] = dockerDir
	}
	user, userErr := user.Current()
	if userErr == nil {
		res["USERNAME"] = user.Username
	}
	res["OS_NAME"] = runtime.GOOS
	return
}

func printContextDebug(ctx MainCtx) {

	fmt.Println("#####################")
	fmt.Println("Context of template :")
	fmt.Println("#####################")

	for name, subCtx := range ctx {
		fmt.Println(name + ":")
		for k, v := range subCtx {
			fmt.Printf("    %s: \"%s\"\n", k, v)
		}
	}
	fmt.Println()
}

func GetContext(dir string, debug bool) (mainCtx MainCtx) {
	mainCtx = MainCtx{
		"Env":   envContext(),
		"Local": localContext(dir),
		"Git":   gitContext(dir),
	}

	if debug {
		printContextDebug(mainCtx)
	}

	return
}
