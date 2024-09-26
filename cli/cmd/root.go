package cmd

import (
  "fmt"
  "os"
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
  Use:   "server-runner",
  Short: "server-runner is a simple CLI to run go-by-example project",
  Long: `server-runner is a simple CLI to run go-by-example project.
         It provides options to run it with DB, env vars etc.
         ** WIP **`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("This is CLI for go-by-example project\n" +
      "Try running one of commands:\n" +
      "serv-runner db-setup\n" +
      "serv-runner test (TBD)\n" +
      "serv-runner local\n")
  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Fprintln(os.Stderr, err)
    os.Exit(1)
  }
}
