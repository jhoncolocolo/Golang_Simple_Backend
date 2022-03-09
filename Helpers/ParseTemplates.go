package Helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// parse.go
func ParseHtmlTemplates() *template.Template {
	var directories []string
	var filenames []string

	// Root directory of template files
	root := "views"

	// Get all directories on /templates and check if there's repeated files
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// Is file
			filename := info.Name()

			hasRepeatedFiles := contains(filenames, filename)
			if hasRepeatedFiles {
				return fmt.Errorf("You can't have repeated template files: %s", filename)
			}

			filenames = append(filenames, filename)
		} else {
			// Is directory
			directories = append(directories, path)
		}

		return nil
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Create a template for parsing all directories
	tmpl := template.Must(template.ParseGlob(root + "/*.tmpl"))

	// Parse all directories (minus the root directory)
	for _, path := range directories[1:] {
		pattern := path + "/*.tmpl"
		template.Must(tmpl.ParseGlob(pattern))
	}

	return tmpl
}

// contains private method
func contains(filenames []string, filename string) bool {
	for _, f := range filenames {
		if f == filename {
			return true
		}
	}
	return false
}
