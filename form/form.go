package form

import (
	"github.com/charmbracelet/huh"
)

func CreateForm(formData *FormData) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name").
				Placeholder("My Awesome Project").
				Value(&formData.ProjectName),
		),
		huh.NewGroup(
			huh.NewText().
				Title("Project description").
				Value(&formData.Description),
		),
		huh.NewGroup(
			huh.NewText().
				Title("Install instructions").
				Value(&formData.InstallInstructions),
		),
		huh.NewGroup(
			huh.NewText().
				Title("Usage instructions").
				Value(&formData.UsageInstructions),
		),
	)
}
