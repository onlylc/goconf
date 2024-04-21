package cmd

import (
	"errors"
	"fmt"
	"goconf/cmd/config"
	"goconf/cmd/crontab"
	"os"
	"goconf/cmd/api"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "goconf",
	Short:        "goconf",
	SilenceUsage: true,
	Long:         `goconf`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	fmt.Printf("%s\n", "欢迎使用查看命令")
}

func init() {
	rootCmd.AddCommand(config.StartCmd)
	rootCmd.AddCommand(crontab.StartCmd)
	rootCmd.AddCommand(api.StartCmd)
}

// Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
