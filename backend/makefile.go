package backend

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type MakefileEntry struct {
	Prefix      string `json:"prefix"`
	Value       string `json:"value"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (a *App) GenerateMakefile(entries []MakefileEntry) string {
	var content strings.Builder

	// Loop over the entries and add each one to the Makefile
	for _, entry := range entries {
		// Write the .PHONY target with the title and description
		content.WriteString(fmt.Sprintf(".PHONY: %s\n%s: ## %s.\n", entry.Title, entry.Title, entry.Description))

		// Write the command to execute with the prefix and value
		content.WriteString(fmt.Sprintf("\t@%s %s\n\n", entry.Prefix, entry.Value))
	}

	// Write the Makefile to disk
	err := ioutil.WriteFile(os.Getenv("HOME")+"/Desktop/Makefile", []byte(content.String()), 0644)
	if err != nil {
		return fmt.Sprintf("Unable to write file: %v", err)
	}

	return fmt.Sprintf("Makefile created : %v", os.Getenv("HOME")+"/Desktop/Makefile")
}
