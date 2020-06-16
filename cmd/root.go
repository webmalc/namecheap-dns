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
	server  Server
	client  Client
}

// change runs changer
func (r *CommandRouter) change(cmd *cobra.Command, args []string) {
	r.changer.Change(args[0])
}

// server runs server
func (r *CommandRouter) runServer(cmd *cobra.Command, args []string) {
	r.server.Run()
}

// request makes a request to the server
func (r *CommandRouter) request(cmd *cobra.Command, args []string) {
	r.client.Request()
}

// Run the router.
func (r *CommandRouter) Run() {
	r.rootCmd.AddCommand(
		&cobra.Command{
			Use:   "change [ip to change]",
			Short: "Run the DNS changer",
			Run:   r.change,
			Args:  cobra.ExactArgs(1),
		},
		&cobra.Command{
			Use:   "request",
			Short: "Make request to the server",
			Run:   r.request,
			Args:  cobra.ExactArgs(0),
		},
		&cobra.Command{
			Use:   "server",
			Short: "Run the server",
			Run:   r.runServer,
			Args:  cobra.ExactArgs(0),
		},
	)
	err := r.rootCmd.Execute()
	if err != nil {
		r.logger.Error(errors.Wrap(err, "root command"))
	}
}

// NewCommandRouter creates a new CommandRouter.
func NewCommandRouter(
	log ErrorLogger, ch Changer, s Server, cl Client,
) CommandRouter {
	return CommandRouter{
		logger:  log,
		rootCmd: &cobra.Command{Use: "ndns.app"},
		changer: ch,
		server:  s,
		client:  cl,
	}
}
