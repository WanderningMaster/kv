package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/WanderningMaster/kv/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kv",
	Short: "A CLI for storing key-values",
	Long:  "A CLI for storing key-values",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		cfg := config.LoadConfig()

		ctx = context.WithValue(ctx, "cfg", cfg)

		cmd.SetContext(ctx)

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
