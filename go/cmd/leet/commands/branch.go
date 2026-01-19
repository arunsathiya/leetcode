package commands

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func CreateBranch(branchName string, useWorktree bool) error {
	if useWorktree {
		return createWorktree(branchName)
	}
	
	// Create regular git branch
	cmd := exec.Command("git", "branch", branchName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create branch: %v", err)
	}
	
	cmd = exec.Command("git", "checkout", branchName)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to checkout branch: %v", err)
	}
	
	fmt.Printf("✅ Created and checked out branch: %s\n", branchName)
	return nil
}

func createWorktree(branchName string) error {
	// Get repo root
	repoRoot, err := getRepoRoot()
	if err != nil {
		return err
	}
	
	// Create worktree path
	worktreePath := filepath.Join(repoRoot, ".worktrees", branchName)
	
	// Create worktree with new branch
	cmd := exec.Command("git", "worktree", "add", "-b", branchName, worktreePath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create worktree: %v", err)
	}
	
	fmt.Printf("✅ Created worktree: %s\n", worktreePath)
	return nil
}

func GetWorktreePath(branchName string) string {
	repoRoot, err := getRepoRoot()
	if err != nil {
		return ""
	}
	
	worktreePath := filepath.Join(repoRoot, ".worktrees", branchName)
	
	// Check if worktree exists
	if _, err := os.Stat(worktreePath); err == nil {
		return worktreePath
	}
	
	return ""
}

func SuggestBranchName(number int, title string) string {
	kebabTitle := toKebabCase(title)
	return fmt.Sprintf("solve/%d-%s", number, kebabTitle)
}

func getRepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get repo root: %v", err)
	}
	
	return strings.TrimSpace(string(output)), nil
}
