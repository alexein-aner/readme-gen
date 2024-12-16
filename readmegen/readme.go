package readmegen

import (
	"fmt"
	"os"
	"text/template"
	"readme-gen/form"
)

type ReadmeData struct {
	ProjectName        string
	Description        string
	InstallInstructions string
	UsageInstructions  string
}

func GenerateReadme(formData *form.FormData) error {
	readmeData := ReadmeData{
		ProjectName:        formData.ProjectName,
		Description:        formData.Description,
		InstallInstructions: formData.InstallInstructions,
		UsageInstructions:  formData.UsageInstructions,
	}

	readmeTemplate := `# {{.ProjectName}}

## Description
{{.Description}}

## Installation Instructions
{{.InstallInstructions}}

## Usage Instructions
{{.UsageInstructions}}
`

	tmpl, err := template.New("readme").Parse(readmeTemplate)
	if err != nil {
		return fmt.Errorf("Error parsing template: %w", err)
	}

	file, err := os.Create("README.md")
	if err != nil {
		return fmt.Errorf("Error creating README.md: %w", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, readmeData)
	if err != nil {
		return fmt.Errorf("Error writing to README.md: %w", err)
	}

	return nil
}
