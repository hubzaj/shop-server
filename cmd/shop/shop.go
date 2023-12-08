package shop

import (
	"github.com/hubzaj/golang-component-test/pkg/runner"
	"github.com/spf13/cobra"
)

func NewShopCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "shop",
		Short: "start shop service",
		Run: func(cmd *cobra.Command, args []string) {
			runner.StartShop()
		},
	}
}
