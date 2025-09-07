package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/arunsathiya/leetcode/cmd/leet/commands"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		if len(os.Args) < 5 {
			fmt.Println("Usage: leet new <number> <title> <function_signature> [--branch] [--worktree]")
			fmt.Println("Example: leet new 123 \"Two Sum\" \"func twoSum(nums []int, target int) []int\"")
			fmt.Println("         leet new 123 \"Two Sum\" \"func twoSum(nums []int, target int) []int\" --branch")
			os.Exit(1)
		}
		
		number, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error: Invalid problem number '%s'\n", os.Args[2])
			os.Exit(1)
		}
		
		title := os.Args[3]
		signature := os.Args[4]
		
		// Check for flags
		createBranch := false
		useWorktree := false
		for i := 5; i < len(os.Args); i++ {
			switch os.Args[i] {
			case "--branch":
				createBranch = true
			case "--worktree":
				createBranch = true
				useWorktree = true
			}
		}
		
		if err := commands.NewProblem(number, title, signature, createBranch, useWorktree); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		
	case "test":
		var problem string
		if len(os.Args) > 2 {
			problem = os.Args[2]
		}
		if err := commands.RunTests(problem); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		
	case "bench":
		var problem string
		if len(os.Args) > 2 {
			problem = os.Args[2]
		}
		if err := commands.RunBenchmarks(problem); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		
	case "branch":
		if len(os.Args) < 4 {
			fmt.Println("Usage: leet branch <number> <title> [--worktree]")
			fmt.Println("Example: leet branch 123 \"Two Sum\"")
			fmt.Println("         leet branch 123 \"Two Sum\" --worktree")
			os.Exit(1)
		}
		
		number, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Error: Invalid problem number '%s'\n", os.Args[2])
			os.Exit(1)
		}
		
		title := os.Args[3]
		useWorktree := len(os.Args) > 4 && os.Args[4] == "--worktree"
		
		branchName := commands.SuggestBranchName(number, title)
		if err := commands.CreateBranch(branchName, useWorktree); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		
	default:
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("LeetCode CLI Tool")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  leet new <number> <title> <function_signature>  Create new problem")
	fmt.Println("  leet test [problem_number]                      Run tests")
	fmt.Println("  leet bench [problem_number]                     Run benchmarks")
	fmt.Println("  leet branch <number> <title> [--worktree]       Create feature branch")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  leet new 123 \"Two Sum\" \"func twoSum(nums []int, target int) []int\"")
	fmt.Println("  leet new 123 \"Two Sum\" \"func twoSum(nums []int, target int) []int\" --branch")
	fmt.Println("  leet test 123")
	fmt.Println("  leet test")
	fmt.Println("  leet bench 123")
	fmt.Println("  leet branch 123 \"Two Sum\"")
	fmt.Println("  leet branch 123 \"Two Sum\" --worktree")
}