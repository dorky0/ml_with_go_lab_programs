package main

import (
	
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

const (
	repoURL       = "https://github.com/dorky0/ml_with_go_lab_programs.git"
	defaultBranch = "main"
)

func main() {
	experimentName := "exp1a" // --> change experiment here

	runCmd := func(dir string, name string, args ...string) {
		cmd := exec.Command(name, args...)
		cmd.Dir, cmd.Stdout, cmd.Stderr = dir, os.Stdout, os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("❌ Command '%s' failed: %v", name, err)
		}
	}

	os.RemoveAll(experimentName)
	if err := os.Mkdir(experimentName, 0755); err != nil {
		log.Fatalf("❌ Failed to create directory: %v", err)
	}

	runCmd(experimentName, "git", "init")
	runCmd(experimentName, "git", "remote", "add", "origin", repoURL)
	runCmd(experimentName, "git", "config", "core.sparseCheckout", "true")

	sparseFile := filepath.Join(experimentName, ".git", "info", "sparse-checkout")
	if err := os.WriteFile(sparseFile, []byte(experimentName+"/"), 0666); err != nil {
		log.Fatalf("❌ Failed to write sparse-checkout file: %v", err)
	}
	runCmd(experimentName, "git", "pull", "--depth=1", "origin", defaultBranch)
	
	log.Println("\n\n---  (this below only actual output)   ---\n")
	runCmd(filepath.Join(experimentName, experimentName), "go", "run", experimentName+".go")
}
