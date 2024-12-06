package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

func main() {
    url, err := getRepoUrl()
    if err != nil {
        fmt.Printf("Error %v\n", err)
        return
    }

    fmt.Printf("url: %s\n", url)

    err = openBrowser(url)
    if err != nil {
        fmt.Printf("Error %v\n", err)
    }
}

func getRepoUrl() (string, error) {
    cwd, err := os.Getwd()
    if err != nil {
        return "", fmt.Errorf("Failed to get current working directory: %v", err)
    }

    fmt.Printf("cwd: %s\n", cwd)

    args := []string{"config", "--get", "remote.origin.url"}

    cmd := exec.Command("git", args...)

    cmd.Dir = cwd

    output, err := cmd.CombinedOutput()
    if err != nil {
        return "", fmt.Errorf("Failed to run git command: %v", err)
    }

    httpUrl := getHttpUrl(string(output))

    return httpUrl, nil
}

func getHttpUrl(url string) string {
    re := regexp.MustCompile(`git@github\..*?:`)

    result := re.ReplaceAllString(url, "https://github.com/")

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
