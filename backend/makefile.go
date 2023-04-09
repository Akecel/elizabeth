package backend

import (
    "fmt"
    "os"
)
// CreateMakefile create a empty makefile and name
func (a *App) CreateMakefile() string {
	err := os.WriteFile(os.Getenv("HOME") + "/Makefile", []byte("Hello"), 0755)
    if err != nil {
        return fmt.Sprintf("Unable to write file: %v", err)
    }

	return fmt.Sprintf("File created!")
}
