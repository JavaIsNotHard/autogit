package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func getUnstagedFiles() []string {
	cmd := exec.Command("git", "diff", "--name-only")

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(output))

	var unstagedFiles []string

	var current = 0
	for i, filesChar := range output {
		if filesChar == '\n' {
			unstagedFiles = append(unstagedFiles, string(output[current:i]))
			current += i + 1
		}
	}

	return unstagedFiles
}

func getUntrackedFiles() []string {
	cmd := exec.Command("git", "ls-files", "--others", "--exclude-standard")

	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}

	var untrackedfiles []string

	var current = 0
	for i, filesChar := range output {
		if filesChar == '\n' {
			untrackedfiles = append(untrackedfiles, string(output[current:i]))
			current += i + 1
		}
	}

	return untrackedfiles
}

func addCommitAndPusd() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Do you want to proceed further ?")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "yes" || input == "y" {
		fmt.Println("Enter the untracked files you want to add: ")
		for i, value := range getUntrackedFiles() {
			fmt.Printf("%d %v\n", i+1, value)
		}
	} else {
		fmt.Println("Hello World NO")
	}
}

func main() {
	addCommitAndPusd()
}
