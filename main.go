package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"github.com/inconshreveable/go-update"
	"github.com/spf13/cobra"
)

const Version = "v2.0.0"
const UpdateURL = "https://github.com/PunamOffice/mycli/releases/latest/download/punamcli.exe"

func main() {
	rootCmd := &cobra.Command{
		Use:   "punamcli",
		Short: "My personal CLI tool",
	}

	helloCmd := &cobra.Command{
		Use:   "hello",
		Short: "Prints a greeting",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello! I am Punam. Welcome to punamcli!")
		},
	}

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Shows the current version of mycli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("punamcli version", Version)
		},
	}

	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "Updates mycli to the latest version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Checking for updates...")
			err := selfUpdate()
			if err != nil {
				fmt.Println("Update failed:", err)
				os.Exit(1)
			}
			fmt.Println("Update successful! Restart mycli to use the new version.")
		},
	}
	syncCmd := &cobra.Command{
    		Use:   "sync",
    		Short: "Saves and pushes your code to GitHub",
    		Run: func(cmd *cobra.Command, args []string) {
        		fmt.Println("Syncing code to GitHub...")

        		// Step 1: git add .
        		run("git", "add", ".")

        		// Step 2: git commit
        		run("git", "commit", "-m", "auto sync from punamcli")

        		// Step 3: git push
        		run("git", "push")

        		fmt.Println("Done! Code is on GitHub.")
    		},
	}

	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(syncCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func selfUpdate() error {
	url := UpdateURL

	fmt.Println("Downloading from:", url)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("bad response from server: %s", resp.Status)
	}

	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return fmt.Errorf("could not apply update: %w", err)
	}

	return nil
}

func run(name string, args ...string) {
    cmd := exec.Command(name, args...)
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    cmd.Run()
}