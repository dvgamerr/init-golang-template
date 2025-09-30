package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/yourusername/init-golang-template/internal/config"
)

var (
	Version   = "dev"
	BuildTime = "unknown"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Application error: %v", err)
	}
}

func run() error {
	rootCmd := &cobra.Command{
		Use:     "app",
		Short:   "Go Application Template",
		Long:    "A modern Go application template with best practices",
		Version: fmt.Sprintf("%s (built: %s)", Version, BuildTime),
		RunE:    runApp,
	}

	rootCmd.Flags().StringP("config", "c", "", "config file path")
	rootCmd.Flags().StringP("env", "e", "development", "environment (development, staging, production)")

	return rootCmd.Execute()
}

func runApp(cmd *cobra.Command, _ []string) error {
	configPath, err := cmd.Flags().GetString("config")
	if err != nil {
		return fmt.Errorf("failed to get config flag: %w", err)
	}

	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return fmt.Errorf("failed to get env flag: %w", err)
	}

	cfg, err := config.Load(configPath, env)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	fmt.Fprintf(os.Stdout, "Starting application...\n")
	fmt.Fprintf(os.Stdout, "Environment: %s\n", cfg.Environment)
	fmt.Fprintf(os.Stdout, "Server: %s:%d\n", cfg.Server.Host, cfg.Server.Port)

	return nil
}
