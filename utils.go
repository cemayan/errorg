package errorg

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	RepoUrl = "github.com/cemayan/errorg"
)

// getModuleFolder returns library location that is downloaded
func getModuleFolder() (string, error) {
	cmdOut, err := exec.Command("go", "list", "-m", "-f", "'{{.Dir}}'", RepoUrl).Output()
	return strings.Trim(strings.TrimSpace(string(cmdOut)), "''"), err
}

// readYaml reads the given yaml
func readYaml(name string, configFolder string) []byte {

	currentPath, err := getModuleFolder()
	if err != nil {
		fmt.Errorf("fatal error config file: %w", err)
	}

	currentPath = currentPath + "/" + configFolder
	currentFile := currentPath + "/" + name + ".yaml"

	file, err := os.ReadFile(currentFile)
	if err != nil {
		return nil
	}

	return file
}
