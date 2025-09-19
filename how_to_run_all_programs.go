// Filename: main.go
// Description: This program downloads and runs a specific folder from any
// GitHub repository by using its direct URL.

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// --- STEP 1: Configure Your Project Details ---
const (
	// Paste the full HTTPS URL of your GitHub repository here.
	// It must end with .git
	repoURL = "https://github.com/YOUR_USERNAME/YOUR_REPOSITORY.git"

	// The repository's default branch (usually "main" or "master").
	defaultBranch = "main"
)

// main is the entry point of the program. It orchestrates the process of
// downloading and running the specified experiment.
func main() {
	// --- STEP 2: Choose Which Experiment to Run ---
	// Change this value to the name of the folder you want to run, e.g., "exp2".

	experimentName := "exp1"        // -----> change experiment here like exp1a, exp2a, exp2b, exp12

	// --- Program Execution Starts Here ---
	log.Printf("🚀 Starting runner for experiment: %s", experimentName)
	log.Printf(">> Target repository: %s", repoURL)

	// Action 1: Download just the specific experiment folder.
	log.Println(">> Phase 1: Downloading remote folder...")
	if err := downloadExperiment(experimentName); err != nil {
		log.Fatalf("❌ Failed to download experiment: %v", err)
	}
	log.Println("✅ Download complete.")

	// Action 2: Run the Go program inside the downloaded folder.
	log.Println(">> Phase 2: Running the experiment's code...")
	if err := runExperiment(experimentName); err != nil {
		log.Fatalf("❌ Failed to run experiment: %v", err)
	}
	log.Println("🎉 Experiment finished successfully!")
}

// downloadExperiment uses "sparse checkout" to download a specific subdirectory.
func downloadExperiment(expName string) error {
	// For a clean run, remove any old directory with the same name.
	log.Printf("   - Cleaning up workspace for '%s'...", expName)
	os.RemoveAll(expName)

	// Create a new, empty directory to hold the experiment files.
	if err := os.Mkdir(expName, 0755); err != nil {
		return fmt.Errorf("could not create directory '%s': %w", expName, err)
	}

	// --- Sparse Checkout Process ---
	log.Println("   - Initializing a temporary Git repository...")

	// 1. `git init`: Create an empty Git repository.
	if err := runCommand(expName, "git", "init"); err != nil {
		return err
	}
	// 2. `git remote add`: Link it to the remote repository URL.
	if err := runCommand(expName, "git", "remote", "add", "origin", repoURL); err != nil {
		return err
	}
	// 3. `git config`: Enable the sparse checkout feature.
	if err := runCommand(expName, "git", "config", "core.sparseCheckout", "true"); err != nil {
		return err
	}

	// 4. Tell Git which folder we want by writing its path into a special file.
	sparseCheckoutFile := filepath.Join(expName, ".git", "info", "sparse-checkout")
	log.Printf("   - Specifying folder to download: %s", expName)
	if err := os.WriteFile(sparseCheckoutFile, []byte(expName+"/"), 0666); err != nil {
		return fmt.Errorf("could not configure sparse checkout: %w", err)
	}

	// 5. `git pull`: Pull from the repository, downloading only our specified folder.
	log.Printf("   - Fetching from remote...")
	return runCommand(expName, "git", "pull", "--depth=1", "origin", defaultBranch)
}

// runExperiment executes the Go program located inside the downloaded folder.
func runExperiment(expName string) error {
	// The path to the code will be nested, e.g., `./exp1/exp1/exp1.go`.
	runDirectory := filepath.Join(expName, expName)
	fileToRun := expName + ".go"

	log.Printf("   - Executing `go run %s`...", fileToRun)
	log.Println("--- [ Experiment Output Begins ] ---")

	err := runCommand(runDirectory, "go", "run", fileToRun)

	log.Println("--- [ Experiment Output Ends ] ---")
	return err
}

// runCommand is a helper function that executes a system command and shows its output.
func runCommand(workingDir, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = workingDir   // Set the command's working directory.
	cmd.Stdout = os.Stdout // Pipe output to our terminal.
	cmd.Stderr = os.Stderr // Pipe errors to our terminal.
	return cmd.Run()
}