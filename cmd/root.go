package cmd

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// CommandRouter is the main commands router.
type CommandRouter struct {
	logger  ErrorLogger
	rootCmd *cobra.Command
	changer Changer
}

// server runs server.
func (r *CommandRouter) change(cmd *cobra.Command, args []string) {
	r.changer.Change(args[0])
}

// Run the router.
func (r *CommandRouter) Run() {
	r.rootCmd.AddCommand(
		&cobra.Command{
			Use:   "change",
			Short: "Run the DNS changer",
			Run:   r.change,
			Args:  cobra.ExactArgs(1),
		},
	)
	err := r.rootCmd.Execute()
	if err != nil {
		r.logger.Error(errors.Wrap(err, "root command"))
	}
}

// NewCommandRouter creates a new CommandRouter.
func NewCommandRouter(log ErrorLogger, c Changer) CommandRouter {
	return CommandRouter{
		logger:  log,
		rootCmd: &cobra.Command{Use: "ndns.app"},
		changer: c,
	}
}
