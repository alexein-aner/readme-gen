package main

import (
	"fmt"
	"os"
	"readme-gen/form"
	"readme-gen/readmegen"
)

func main() {
	var formData form.FormData

	f := form.CreateForm(&formData)

	err := f.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	err = readmegen.GenerateReadme(&formData)
	if err != nil {
		fmt.Println("Error generating README:", err)
		os.Exit(1)
	}

	fmt.Println("README.md has been generated successfully!")
}
