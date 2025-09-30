package cli

import (
    "fmt"
    "github.com/spf13/cobra"
)

func CmdGovTemplates() *cobra.Command {
    cmd := &cobra.Command{
        Use:   "gov-templates",
        Short: "Show example governance proposal templates",
        RunE: func(cmd *cobra.Command, args []string) error {
            fmt.Println("Templates are in scripts/proposals/:")
            fmt.Println("- text-proposal.json")
            fmt.Println("- param-change.json")
            return nil
        },
    }
    return cmd
}
