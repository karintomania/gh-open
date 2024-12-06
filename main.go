package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

var version string = "0.1.0"

var isPR bool
var isVersion bool

func main() {
	getFlags()

	if isVersion {
        fmt.Printf("gh-open v%s\n", version);
		return
	}

	url, err := getRepoUrl()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = openBrowser(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func getFlags() {
	flag.BoolVar(&isPR, "PR", false, "Open PR page")
	flag.BoolVar(&isPR, "p", false, "Alias for -PR")
    flag.BoolVar(&isVersion, "version", false, "Show version of gh-open")

	flag.Parse()
}

func getRepoUrl() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("Failed to get current working directory: %v", err)
	}

	cmd := exec.Command("git", "config", "--get", "remote.origin.url")

	cmd.Dir = cwd

	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("Failed to run git command: %v", err)
	}

	errOut, err := io.ReadAll(stderr)
	if len(errOut) != 0 || err != nil {
		return "", fmt.Errorf("Failed to run git command: %v", errOut)
	}

	output, err := io.ReadAll(stdout)
	if err != nil {
		return "", fmt.Errorf("Failed to run git command: %v", err)
	}

	if len(output) == 0 {
		return "", fmt.Errorf("No the remote repository found")
	}

	httpUrl := getHttpUrl(string(output))

	if isPR {
		httpUrl = httpUrl + "/pulls"
	}

	return httpUrl, nil
}

// replace git@github.com: from ssh repo url with http url
func getHttpUrl(url string) string {
	re := regexp.MustCompile(`git@github\..*?:`)

	result := re.ReplaceAllString(url, "https://github.com/")

	re = regexp.MustCompile(`\.git\s?$`)

	result = re.ReplaceAllString(result, "")

	return result
}

func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "rundll32"
		args = []string{"url.dll,FileProtocolHandler", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	case "linux":
		cmd = "xdg-open"
		args = []string{url}
	default:
		return fmt.Errorf("unsupported platform")
	}

	return exec.Command(cmd, args...).Start()
}
