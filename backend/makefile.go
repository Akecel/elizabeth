package backend

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type MakefileEntry struct {
	Command     string `json:"command"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (a *App) GenerateMakefile(entries []MakefileEntry) string {
	var content strings.Builder

	header, _ := ioutil.ReadFile("./assets/makefile_header.txt")
	content.WriteString(fmt.Sprintf("%s\n\n", string(header)))

	// Loop over the entries and add each one to the Makefile
	for _, entry := range entries {
		// Write the .PHONY target with the title and description
		content.WriteString(fmt.Sprintf(".PHONY: %s\n%s: ## %s.\n", entry.Title, entry.Title, entry.Description))

		// Write the command to execute with the prefix and value
		content.WriteString(fmt.Sprintf("\t@%s\n\n", entry.Command))
	}

	// Write the Makefile to disk
	err := ioutil.WriteFile(os.Getenv("HOME")+"/Desktop/Makefile", []byte(content.String()), 0644)
	if err != nil {
		return fmt.Sprintf("Unable to write file: %v", err)
	}

	return fmt.Sprintf("Makefile created : %v", os.Getenv("HOME")+"/Desktop/Makefile")
}
