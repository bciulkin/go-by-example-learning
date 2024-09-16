package cmd

import (
  "fmt"
  "os/exec"
  "github.com/spf13/cobra"
  "errors"
)

func init() {
  rootCmd.AddCommand(setupCmd)

}

var setupCmd = &cobra.Command{
  Use:   "setup [db user] [db password]",
  Short: "Setups DB credentials with flags",
  Long:  `Setups env variables to run *go-by-example* server locally.
  More info in project's README`,
  Args: func(cmd *cobra.Command, args []string) error {
    if len(args) != 2 {
      return errors.New("requires two args exactly")
    }
    return nil
  },
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println(args)
    command := exec.Command("export","DBUSER="+ args[0],";","export","DBPASS="+args[0])
    stdout, err := command.Output()

    if err != nil {
      fmt.Println(err.Error())
      return
    }

    fmt.Println(string(stdout))
  },
}
