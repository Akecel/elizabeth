package backend

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Command struct {
	Command string `json:"command"`
}

type MakefileEntry struct {
	Commands    []Command `json:"commands"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

func (a *App) GenerateMakefile(entries []MakefileEntry) string {
	var content strings.Builder

	// Write the header of the Makefile
	header, _ := ioutil.ReadFile("./assets/makefile_header.txt")
	content.WriteString(fmt.Sprintf("%s\n\n", string(header)))

	entries = append(entries, SetHelpCommand())

	// Loop over the entries and add each one to the Makefile
	for _, entry := range entries {
		// Write the .PHONY target with the title and description
		content.WriteString(fmt.Sprintf(".PHONY: %s\n%s: ## %s.\n", entry.Title, entry.Title, entry.Description))

		// Write the different commands
		for _, cmd := range entry.Commands {
			content.WriteString(fmt.Sprintf("\t@%s\n", cmd.Command))
		}

		content.WriteString("\n")
	}

	// Write the Makefile to disk
	err := ioutil.WriteFile(os.Getenv("HOME")+"/Desktop/Makefile", []byte(content.String()), 0644)
	if err != nil {
		return fmt.Sprintf("Unable to write file: %v", err)
	}

	return fmt.Sprintf("Makefile created : %v", os.Getenv("HOME")+"/Desktop/Makefile")
}

func SetHelpCommand() MakefileEntry {
	return MakefileEntry{
		Commands: []Command{
			{
				Command: `printf "Usage: make <command>\n\n"`,
			}, {
				Command: `printf "Commands:\n"`,
			}, {
				Command: `awk -F ':(.*)## ' '/^[a-zA-Z0-9%\\\/_.-]+:(.*)##/ { \
					printf "  \033[36m%-30s\033[0m %s\n", $$1, $$NF \
				}' $(MAKEFILE_LIST)`,
			},
		},
		Title:       "help",
		Description: "Provides help information on available commands.",
	}
}
