package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/ism0080/profile/data"
	"github.com/ism0080/profile/templates"
	cp "github.com/otiai10/copy"
)

const (
	BaseUrl  string = "https://isaacmackle.com"
	OutputDir string = "public"
)

func main() {
	createDirIfNotExists()

	// Create an index page.
	homePage := createFile("index.html")
	homePageData := data.HomePageData(BaseUrl)

	// Write it out.
	err := templates.Index(homePageData).Render(context.Background(), homePage)
	if err != nil {
		log.Fatalf("failed to write index page: %v", err)
	}

	runTailwind()
	err = cp.Copy("./assets", fmt.Sprintf("%s/assets", OutputDir))
	if err != nil {
		log.Fatalf("Failed to copy assets directory: %v", err)
	}
}

func runTailwind() {
	cmd := exec.Command("tailwindcss", "-i", "./index.css", "-o", fmt.Sprintf("%s/styles.min.css", OutputDir), "--minify")

	_, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func createDirIfNotExists() {
	// Check if the directory exists
	if _, err := os.Stat(OutputDir); os.IsNotExist(err) {
		// Directory does not exist, create it
		if err := os.Mkdir(OutputDir, 0755); err != nil {
			log.Fatalf("failed to create output directory: %v", err)
		}
	} else if err != nil {
		// Some other error occurred while checking directory existence
		log.Fatalf("failed to check directory existence: %v", err)
	}
}

func createFile(fileName string) *os.File {
	name := path.Join(OutputDir, fileName)
	f, err := os.Create(name)
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	return f
}
