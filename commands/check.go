package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

type Check struct {
	*Cmd
}

func (root *Cmd) initCheck() {
	c := Check{ Cmd: root }

	checkCmd := &cobra.Command{
		Use: "check",
		Short: "Consul Check endpoint interface",
		Long: "Consul Check endpoint interface",
		Run: func (cmd *cobra.Command, args []string) {
			root.Help()
		},
	}

	c.AddDeregisterSub(checkCmd)
	c.AddFailSub(checkCmd)
	c.AddPassSub(checkCmd)
	c.AddRegisterSub(checkCmd)
	c.AddWarnSub(checkCmd)

	c.AddCommand(checkCmd)
}

func (c *Check) CheckIdArg(args []string) error {
	switch {
	case len(args) == 0:
		return fmt.Errorf("No check id specified")
	case len(args) > 1:
		return fmt.Errorf("Only one check id allowed")
	}

	return nil
}