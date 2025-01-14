package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// Function to open the browser automatically
func openBrowser(url string) {
	var err error

	// Use system-specific commands to open the browser
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Printf("Failed to open browser: %v\n", err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dndMenu2.html")
}

func renameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	sourceDir := r.FormValue("source")
	targetDir := r.FormValue("target")
	newName := r.FormValue("name")

	if sourceDir == "" || targetDir == "" || newName == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	var logs []string
	logs = append(logs, fmt.Sprintf("Renaming files in '%s' to '%s' in '%s'.", sourceDir, newName, targetDir))

	// Perform renaming
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			logs = append(logs, fmt.Sprintf("Error accessing file '%s': %v", path, err))
			return nil
		}

		if !info.IsDir() {
			newFileName := fmt.Sprintf("%s_%s", newName, info.Name())
			newPath := filepath.Join(targetDir, newFileName)

			if err := os.Rename(path, newPath); err != nil {
				logs = append(logs, fmt.Sprintf("Failed to rename '%s': %v", path, err))
			} else {
				logs = append(logs, fmt.Sprintf("Renamed '%s' to '%s'", path, newPath))
			}
		}

		return nil
	})

	if err != nil {
		logs = append(logs, fmt.Sprintf("Error walking through directory: %v", err))
	}

	w.Write([]byte(strings.Join(logs, "\n"))) // Send logs back to the browser
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/rename", renameHandler)

	url := "http://localhost:3000"
	fmt.Printf("Server listening at %s\n", url)

	// Open the browser automatically
	go openBrowser(url)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
