package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func main() {
    url := getRepoUrl()
    err := openBrowser(url)
    if err != nil {
        fmt.Printf("Error %v\n", err)
    }
}

func getRepoUrl() string {
    cwd, err := os.Getwd()
    if err != nil {
        fmt.Errorf("Failed to get current working directory: %v", err)
    }

    fmt.Printf("cwd: %s\n", cwd)

    cmd := exec.Command("git", "status")

    cmd.Dir = cwd

    output, err := cmd.CombinedOutput()
    if err != nil {
        fmt.Errorf("Failed to run git command: %v", err)
    }

    // Print the output
    fmt.Printf("Git command output:\n%s\n", output)
    return string(output)
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
