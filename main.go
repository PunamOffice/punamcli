package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"

	"github.com/inconshreveable/go-update"
	"github.com/spf13/cobra"
)

// Change this version number manually each release
const Version = "v1.0.0"

// This is where your new binary will be downloaded from.
// We will set this up properly in the next step.
const UpdateURL = "https://github.com/YOUR_USERNAME/mycli/releases/latest/download/mycli_%s.exe"

func main() {
	rootCmd := &cobra.Command{
		Use:   "mycli",
		Short: "My personal CLI tool",
	}

	// hello command
	helloCmd := &cobra.Command{
		Use:   "hello",
		Short: "Prints a greeting",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello! I am Punam.")
		},
	}

	// version command
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Shows the current version of mycli",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("mycli version", Version)
		},
	}

	// update command — the real challenge
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

	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(updateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func selfUpdate() error {
	// Build the download URL based on the OS
	os_name := runtime.GOOS   // "windows", "linux", "darwin"
	url := fmt.Sprintf(UpdateURL, os_name)

	fmt.Println("Downloading from:", url)

	// Download the new binary
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("bad response from server: %s", resp.Status)
	}

	// Apply the update — this does the rename trick safely
	err = update.Apply(resp.Body, update.Options{})
	if err != nil {
		return fmt.Errorf("could not apply update: %w", err)
	}

	return nil
}