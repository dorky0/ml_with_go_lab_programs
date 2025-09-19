GitHub Experiment Runner
A simple Go utility to download and run a single experiment folder from a GitHub repository without cloning the entire project.

This script is perfect for large repositories where you only need to work with a small part of the codebase at a time. It uses Git's sparse-checkout feature to efficiently fetch only the files you need.

## Features
⚡️ Efficient: Downloads only the folder you need, saving significant time and disk space.

⚙️ Simple to Use: Run any experiment just by changing a single line of code.

🚀 Fully Automated: Fetches dependencies, compiles, and runs your Go experiment with a single command.

## Prerequisites
Before you begin, ensure you have the following installed on your system:

Go: Version 1.16 or later

Git: Must be installed and accessible in your system's PATH

## 🚀 How to Use
### 1. Choose Your Experiment
Open the main.go file in a text editor and find the main function. To run a different experiment, simply change the experimentName value.

Go

func main() {
	// --- Choose Which Experiment to Run ---
	// Change this value to the name of the folder you want to run.
	// For example, to run the "exp5" folder, change it to "exp5".
	experimentName := "exp1"

    // ... rest of the code ...
}
### 2. Run the Program
Open your terminal, navigate to the directory where main.go is saved, and run the following command:

Bash

go run main.go
The script will automatically download and run the experiment you selected.

## 🤔 How It Works
This script uses a Git feature called sparse checkout. Instead of cloning the whole repository, it tells Git, "I am only interested in the contents of the specific folder you chose." It then pulls only that folder, making the process extremely fast and lightweight