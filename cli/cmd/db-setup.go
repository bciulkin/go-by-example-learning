package cmd

import (
  "log"
  "os/exec"
  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(dbSetupCmd)
}

var dbSetupCmd = &cobra.Command{
  Use:   "db-setup",
  Short: "Run db-setup, to initilize DB for running server locally",
  Long:  `Run db-setup in order to initlize DB required for running server locally.
  By default runs with MySql local default setup. More info in project's README`,
  Run: func(cmd *cobra.Command, args []string) {
    command := exec.Command("mysql", "-u", "root", "-pnew-password", "animals", "<", "./db_scripts/init.sql")
    stdout, err := command.Output()

    if err != nil {
      log.Println(err.Error())
      return
    }

    log.Println(string(stdout))

  },
}
