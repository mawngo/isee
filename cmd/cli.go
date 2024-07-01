package cmd

import (
	"fmt"
	"github.com/mawngo/isee/cmd/grayscale"
	"github.com/phsym/console-slog"
	"github.com/spf13/cobra"
	"log/slog"
	"os"
	"time"
)

func Init() *slog.LevelVar {
	level := &slog.LevelVar{}
	logger := slog.New(
		console.NewHandler(os.Stderr, &console.HandlerOptions{
			Level:      level,
			TimeFormat: time.Kitchen,
		}))
	slog.SetDefault(logger)
	cobra.EnableCommandSorting = false
	return level
}

type CLI struct {
	command *cobra.Command
}

// NewCLI create new CLI instance and setup application config.
func NewCLI() *CLI {
	level := Init()
	command := cobra.Command{
		Use:   "isee files...",
		Short: "Generate ascii art from image",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			debug, err := cmd.Flags().GetBool("debug")
			if err != nil {
				return err
			}
			if debug {
				level.Set(slog.LevelDebug)
			}
			return nil
		},
	}
	command.PersistentFlags().Bool("debug", false, "Enable debug mode")
	command.AddCommand(grayscale.NewCmd())
	command.Flags().SortFlags = false
	return &CLI{&command}
}

func (cli *CLI) Execute() {
	if err := cli.command.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
